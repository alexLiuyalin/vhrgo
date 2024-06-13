package biz

import (
	"context"

	apiweb "vhrgo/api/frontend/web/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type MenuRepo interface {
	List(ctx context.Context, data string) ([]*apiweb.ListMenuReply_Menu, error)
}

type MenuUsecase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUseCase(repo MenuRepo, logger log.Logger) *MenuUsecase {
	return &MenuUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "web", "biz/user")),
	}
}

func (uc *MenuUsecase) List(ctx context.Context) ([]*apiweb.ListMenuReply_Menu, error) {
	return uc.repo.List(ctx, " where parentId IN (SELECT id FROM menu WHERE parentId IS NULL)")
}
