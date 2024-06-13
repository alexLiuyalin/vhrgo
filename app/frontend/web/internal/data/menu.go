package data

import (
	"context"
	"fmt"

	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"
	"vhrgo/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.MenuRepo = (*menuRepo)(nil)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/menu")),
	}
}

func (r menuRepo) List(ctx context.Context, data string) (res []*apiweb.ListMenuReply_Menu, err error) {
	var menus []*model.Menu
	result := &apiweb.ListMenuReply_Menu{}
	sqlStr := "select * from menu " + data
	if err = r.data.db.Select(&menus, sqlStr); err != nil {
		return nil, err
	}
	for _, v := range menus {
		if result, err = r.composes(ctx, v); err != nil {
			continue
		}
		res = append(res, result)
	}
	return
}

// composes 组装返回值
func (r *menuRepo) composes(ctx context.Context, model *model.Menu) (*apiweb.ListMenuReply_Menu, error) {

	// 获取子菜单
	children, err := r.List(ctx, fmt.Sprintf(" where parentId = %d", model.ID))
	if err != nil {
		return nil, err
	}

	data := &apiweb.ListMenuReply_Menu{
		Path:      model.Path.String,
		Component: model.Component.String,
		Name:      model.Name.String,
		IconCls:   model.IconCls.String,
		Meta:      model.RequireAuth.Valid,
		Children:  children,
	}

	return data, nil
}
