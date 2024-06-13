package service

import (
	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService is a user service.
type UserService struct {
	// apiweb.UnimplementedUserServer
	// apiweb.UserHTTPServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc,
		log: log.NewHelper(log.With(logger, "web", "service/user"))}
}

// Create 创建用户
// func (s *UserService) Create(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	return s.uc.Create(ctx, req)
// }

// // Delete 删除用户
// func (s *UserService) Delete(ctx context.Context, req *apicommon.IdRequest) (*apicommon.EmptyReply, error) {
// 	return s.uc.Delete(ctx, req)
// }

// Info 获取个人信息
// func (s *UserService) Info(ctx context.Context, req *apicommon.IdRequest) (*apiweb.UserReply, error) {
// 	return s.uc.Info(ctx, req)
// }

// // List 获取用户列表
// func (s *UserService) List(ctx context.Context, req *apiweb.ListUserRequest) (*apiweb.ListUserReply, error) {
// 	return s.uc.List(ctx, req)
// }

// // Update 更新用户
// func (s *UserService) Update(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	return s.uc.Update(ctx, req)
// }

// // UpdateInfo 更新个人信息
// func (s *UserService) UpdateInfo(ctx context.Context, req *apiweb.UpdateInfoRequest) (*apicommon.EmptyReply, error) {
// 	return s.uc.UpdateInfo(ctx, req)
// }
