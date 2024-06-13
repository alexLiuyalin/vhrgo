package service

import (
	"context"

	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MenuService struct {
	apiweb.MenuHTTPServer

	uc  *biz.MenuUsecase
	log *log.Helper
}

func NewMenuService(
	uc *biz.MenuUsecase,
	logger log.Logger) *MenuService {
	return &MenuService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "core", "service/menu")),
	}
}

func (s *MenuService) List(ctx context.Context,
	req *apiweb.ListMenuRequest) (*apiweb.ListMenuReply, error) {
	res, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.ListMenuReply{
		Data: res,
	}, nil
}
