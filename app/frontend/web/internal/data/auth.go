package data

import (
	"context"
	"time"

	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/user")),
	}
}

func (r *authRepo) GetIdByCid(ctx context.Context, cid string) (string, error) {
	res, err := r.data.cache.Get(ctx, "auth:cid:"+cid).Result()
	if err != nil {
		r.log.Error(err)
	}
	return res, nil
}

func (r *authRepo) SetIdByCid(ctx context.Context, cid, id string) error {
	err := r.data.cache.Set(ctx, "auth:cid:"+cid, id, time.Hour).Err()
	if err != nil {
		r.log.Error(err)
	}
	return nil
}

func (r *authRepo) GetToken(ctx context.Context, id string) (string, error) {
	res, err := r.data.cache.Get(ctx, "auth:token:"+id).Result()
	if err != nil {
		r.log.Error(err)
	}
	return res, nil
}
func (r *authRepo) SetToken(ctx context.Context, id, token string) error {
	err := r.data.cache.Set(ctx, "auth:token:"+id, token, time.Hour*24).Err()
	if err != nil {
		r.log.Error(err)
	}
	return nil
}
