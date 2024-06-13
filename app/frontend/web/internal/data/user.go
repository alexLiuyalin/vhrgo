package data

import (
	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/user")),
	}
}

// Create 创建用户
// func (r *userRepo) Create(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	r.data.user.Create(ctx, &apicore.UserRequest{})
// 	return nil, nil
// }

// // Delete 删除用户
// func (r *userRepo) Delete(ctx context.Context, req *apicommon.IdRequest) (*apicommon.EmptyReply, error) {
// 	r.data.user.Delete(ctx, nil)
// 	return nil, nil
// }

// Info 获取个人信息
// func (r *userRepo) Info(ctx context.Context, req *apicommon.IdRequest) (*apiweb.UserReply, error) {
// 	// r.data.user.Info(ctx, nil)
// 	return nil, nil
// }

// List 获取用户列表
// func (r *userRepo) List(ctx context.Context, req *apiweb.ListUserRequest) (*apiweb.ListUserReply, error) {
// 	r.data.user.List(ctx, nil)
// 	return nil, nil
// }

// // Update 更新用户
// func (r *userRepo) Update(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	r.data.user.Update(ctx, nil)
// 	return nil, nil
// }

// // UpdateInfo 更新个人信息
// func (r *userRepo) UpdateInfo(ctx context.Context, req *apiweb.UpdateInfoRequest) (*apicommon.EmptyReply, error) {
// 	r.data.db.Query("")
// 	return nil, nil
// }
