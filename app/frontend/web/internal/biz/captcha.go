package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	apiweb "vhrgo/api/frontend/web/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/wenlng/go-captcha/captcha"
)

type CaptchaUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewCaptchaUseCase(repo UserRepo, logger log.Logger) *CaptchaUseCase {
	return &CaptchaUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "web", "biz/captcha")),
	}
}

func (uc *CaptchaUseCase) GetCaptcha(ctx context.Context) (*apiweb.GetCaptchaReply, error) {
	capt := captcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return &apiweb.GetCaptchaReply{}, err
	}
	uc.writeCache(dots, key) // 缓存在本地

	return &apiweb.GetCaptchaReply{
		B64:  b64,
		Tb64: tb64,
		Cid:  key,
	}, nil
}

func (uc *CaptchaUseCase) VerifyCaptcha(ctx context.Context,
	req *apiweb.VerifyCaptcha) error {
	cacheData := uc.readCache(req.Cid)
	if cacheData == "" {
		return errors.New("验证超时")
	}
	src := strings.Split(req.VerifyValue, ",")
	var dct map[int]captcha.CharDot
	if err := json.Unmarshal([]byte(cacheData), &dct); err != nil {
		return nil
	}
	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)

			// 检测点位置
			// chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
			if !chkRet {
				break
			}
		}
	}
	if !chkRet {
		return errors.New("验证不通过")
	}
	return nil
}

// 缓存在本地bin下
func (uc *CaptchaUseCase) writeCache(v interface{}, file string) {
	bt, _ := json.Marshal(v)
	// month := time.Now().Month().String()

	path, _ := os.Getwd() // 获取根路径名

	// cacheDir := path + "/app/core/service/bin/captcha/" + month + "/"
	_ = os.MkdirAll(path, 0660)
	file = path + file + ".json"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	// 检查过期文件
	// checkCacheOvertimeFile()
	_, _ = io.WriteString(logFile, string(bt))
	return
}

// 读本地文件
func (uc *CaptchaUseCase) readCache(file string) string {
	month := time.Now().Month().String()
	path, _ := os.Getwd() // 获取根路径名

	cacheDir := path + "/app/core/service/bin/captcha/" + month + "/"
	file = cacheDir + file + ".json"
	if _, err := os.Stat(file); os.IsNotExist(err) { // 检查文件是否存在
		return ""
	}

	bt, err := os.ReadFile(file) // 读内容
	err = os.Remove(file)        // 删除这个文件
	if err == nil {
		return string(bt)
	}
	return ""

}
