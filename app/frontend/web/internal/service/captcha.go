package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"
)

type CaptchaService struct {
	apiweb.CaptChaHTTPServer

	uc  *biz.CaptchaUseCase
	log *log.Helper
}

func NewCaptchaService(
	uc *biz.CaptchaUseCase,
	logger log.Logger) *CaptchaService {
	return &CaptchaService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "core", "service/captcha")),
	}
}

func (s *CaptchaService) Get(ctx context.Context,
	req *apiweb.GetCaptcha) (*apiweb.GetCaptchaReply, error) {
	res, err := s.uc.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CaptchaService) Verify(ctx context.Context,
	req *apiweb.VerifyCaptcha) (*apiweb.VerifyCaptchaReply, error) {
	if err := s.uc.VerifyCaptcha(ctx, &apiweb.VerifyCaptcha{
		Cid:         req.Cid,
		VerifyValue: req.VerifyValue,
	}); err != nil {
		return nil, err
	}
	return &apiweb.VerifyCaptchaReply{
		Message: "success",
		Code:    "200",
	}, nil
}
