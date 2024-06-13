package server

import (
	"context"
	"fmt"
	nethttp "net/http"
	"reflect"
	"strings"

	. "vhrgo/pkg/common"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"google.golang.org/genproto/googleapis/api/httpbody"
	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/conf"
	"vhrgo/app/frontend/web/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	captcha *service.CaptchaService,
	user *service.UserService,
	auth *service.AuthService,
	menu *service.MenuService,
	employee *service.EmployeeService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "X-Client-ID", "X-Otp", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.ExposedHeaders([]string{"token"}),
		)),
		// http.ResponseEncoder(customResponseEncoder), // 自定义返回编码器
		http.ErrorEncoder(customErrorEncoder), // 自定义错误编码器
		http.Middleware(commonMiddleware()),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	apiweb.RegisterUserHTTPServer(srv, user)
	apiweb.RegisterCaptChaHTTPServer(srv, captcha)
	apiweb.RegisterAuthHTTPServer(srv, auth)
	apiweb.RegisterMenuHTTPServer(srv, menu)
	apiweb.RegisterEmployeeHTTPServer(srv, employee)
	return srv
}

// customErrorEncoder 自定义错误编码器
func customErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request,
	err error) {
	se := kerrors.FromError(err) // 获取错误信息

	// 自定义返回格式
	reply := map[string]interface{}{
		"code":    fmt.Sprintf("%d", se.Code),
		"reason":  se.Reason,
		"message": se.Message,
		"data":    se.Metadata,
	}

	// 根据请求头中的 Accept 获取对应的编码器
	codec, _ := http.CodecForRequest(r, "Accept")
	// 序列化
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// customResponseEncoder 自定义返回编码器
// https://github.com/go-kratos/kratos/issues/1952
func customResponseEncoder(w nethttp.ResponseWriter,
	r *nethttp.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	// 实现重定向
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		nethttp.Redirect(w, r, url, code)
		return nil
	}

	var data []byte
	var err error
	// 处理 HttpBody 类型
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		data = httpBody.Data
		// 设置响应头
		w.Header().Set("Content-Type", httpBody.GetContentType())
	} else {
		// 自定义返回格式
		reply := map[string]interface{}{
			"code":    "200",
			"message": "success",
		}

		// 根据请求头中的 Accept 获取对应的编码器
		codec, _ := http.CodecForRequest(r, "Accept")

		// 用反射获取 proto.Message 的 Data 字段
		msgValue := reflect.ValueOf(v)
		dataField := msgValue.Elem().FieldByName("Data")
		if dataField.IsValid() {
			dataValue := dataField.Interface()
			// 序列化
			// data, err = codec.Marshal(dataValue)
			// if err != nil {
			// 	return err
			// }
			reply["data"] = dataValue
		}

		// 获取 Total 字段的值
		totalField := msgValue.Elem().FieldByName("Total")
		if totalField.IsValid() {
			reply["total"] = totalField.Int()
		}

		// 对自定义返回格式进行序列化
		data, err = codec.Marshal(reply)
		if err != nil {
			return err
		}
		// 根据 Accept 设置响应头
		// w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	}

	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// commonMiddleware 通用中间件
func commonMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 取 TOKEN
				auths := strings.SplitN(tr.RequestHeader().Get("Authorization"),
					" ", 2) // 获取头部信息,根据空格分隔头部信息
				if len(auths) == 2 && auths[0] == "Bearer" {
					ctx = context.WithValue(ctx, ContextTokenKey{},
						auths[1])
				}

				// 获取 Cid 信息
				cid := tr.RequestHeader().Get("X-Client-ID")
				if cid != "" {
					// 将 cid 存到 ctx
					ctx = context.WithValue(ctx, ContextCidKey{}, cid)
				}

				// 如果客户端使用代理
				ip := tr.RequestHeader().Get("X-Forwarded-For")
				if ip == "" {
					// 直接取默认
					ip = tr.(*http.Transport).Request().RemoteAddr
				}

				ctx = context.WithValue(ctx, ContextIpKey{}, ip)

			}

			return handler(ctx, req)
		}
	}
}

// // jwtAuthMiddleware 验证token中间件
// func jwtAuthMiddleware(c *conf.Server, auth *biz.AuthUseCase) middleware.Middleware {
// 	return func(handler middleware.Handler) middleware.Handler {
// 		return func(ctx context.Context, req interface{}) (reply interface{},
// 			err error) {
// 			if err := auth.Check(ctx); err != nil {
// 				return ctx, err
// 			}
// 			return handler(ctx, req)
// 		}
// 	}
// }
