package data

import (
	"context"
	"vhrgo/app/frontend/web/internal/biz"
	"vhrgo/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.HrRepo = (*hrRepo)(nil)

type hrRepo struct {
	data *Data
	log  *log.Helper
}

func NewHrRepo(data *Data, logger log.Logger) biz.HrRepo {
	return &hrRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/hr")),
	}
}

func (r hrRepo) Get(ctx context.Context, username string) (hr model.HR, err error) {
	sqlStr := "select * from hr where username = ?"
	err = r.data.db.Get(&hr, sqlStr, username)
	if err != nil {
		return model.HR{}, err
	}
	return
}
