package service

import (
	"context"

	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

type AuthService struct {
	apiweb.AuthHTTPServer

	uc  *biz.AuthUseCase
	log *log.Helper
}

func NewAuthService(
	uc *biz.AuthUseCase,
	logger log.Logger) *AuthService {
	return &AuthService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "core", "service/auth")),
	}
}
func (s *AuthService) Login(ctx context.Context,
	req *apiweb.LoginRequest) (*apiweb.LoginReply, error) {
	token, res, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	tr, _ := transport.FromServerContext(ctx)
	tr.ReplyHeader().Set("Token", token) // 把 token 放在头部
	return &apiweb.LoginReply{
		Id:        res.ID,
		Name:      res.Name.String,
		Phone:     res.Phone.String,
		Telephone: res.Telephone.String,
		Address:   res.Address.String,
		Enabled:   res.Enabled.Bool,
		Username:  res.Username.String,
		Userface:  res.Userface.String,
		Remark:    res.Remark.String,
	}, nil

}

// Logout 登出
func (s *AuthService) Logout(ctx context.Context,
	req *apiweb.LogoutRequest) (*apiweb.Reply, error) {

	if err := s.uc.Logout(ctx); err != nil {
		return nil, err
	}

	return &apiweb.Reply{
		Code:    "200",
		Message: "success",
	}, nil
}
