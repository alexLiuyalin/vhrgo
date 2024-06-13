package data

import (
	"context"
	"time"

	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CaptchaRepo = (*captchaRepo)(nil)

type captchaRepo struct {
	data *Data
	log  *log.Helper
}

func NewCaptchaRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/user")),
	}
}

func (r *captchaRepo) CacheCaptcha(ctx context.Context, cid string, captcha []byte) error {
	err := r.data.cache.Set(ctx, ("captcha:" + cid), captcha, time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *captchaRepo) GetCacheCaptcha(ctx context.Context, cid string) ([]byte, error) {
	res, err := r.data.cache.Get(ctx, ("captcha:" + cid)).Bytes()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *captchaRepo) CacheSuccessCaptcha(ctx context.Context, cid string) error {
	err := r.data.cache.Set(ctx, ("success:captcha:" + cid), "success", time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *captchaRepo) GetSuccessCaptcha(ctx context.Context, cid string) error {
	err := r.data.cache.Get(ctx, ("success:captcha:" + cid)).Err()
	if err != nil {
		return err
	}
	return nil
}
