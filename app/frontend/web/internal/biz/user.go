package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	// // Create 创建用户
	// Create(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error)
	// // Delete 删除用户
	// Delete(ctx context.Context, req *apicommon.IdRequest) (*apicommon.EmptyReply, error)
	// Info 获取个人信息
	// Info(ctx context.Context, req *apicommon.IdRequest) (*apiweb.UserReply, error)
	// List 获取用户列表
	// List(ctx context.Context, req *apiweb.ListUserRequest) (*apiweb.ListUserReply, error)
	// // Update 更新用户
	// Update(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error)
	// // UpdateInfo 更新个人信息
	// UpdateInfo(ctx context.Context, req *apiweb.UpdateInfoRequest) (*apicommon.EmptyReply, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "web", "biz/user")),
	}
}

// // Create 创建用户
// func (uc *UserUsecase) Create(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	return uc.repo.Create(ctx, req)
// }

// // Delete 删除用户
// func (uc *UserUsecase) Delete(ctx context.Context, req *apicommon.IdRequest) (*apicommon.EmptyReply, error) {
// 	return uc.repo.Delete(ctx, req)
// }

// Info 获取个人信息
// func (uc *UserUsecase) Info(ctx context.Context, req *apicommon.IdRequest) (*apiweb.UserReply, error) {
// 	return uc.repo.Info(ctx, req)
// }

// // List 获取用户列表
// func (uc *UserUsecase) List(ctx context.Context, req *apiweb.ListUserRequest) (*apiweb.ListUserReply, error) {
// 	return uc.repo.List(ctx, req)
// }

// // Update 更新用户
// func (uc *UserUsecase) Update(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error) {
// 	return uc.repo.Update(ctx, req)
// }

// // UpdateInfo 更新个人信息
// func (uc *UserUsecase) UpdateInfo(ctx context.Context, req *apiweb.UpdateInfoRequest) (*apicommon.EmptyReply, error) {
// 	return uc.repo.UpdateInfo(ctx, req)
// }
