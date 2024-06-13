package biz

import (
	"context"
	"vhrgo/data/model"
)

type HrRepo interface {
	Get(ctx context.Context, username string) (model.HR, error)
}
