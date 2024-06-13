package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/data/model"
	. "vhrgo/pkg/common"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	// // Create 创建用户
	// Create(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error)
	// // Delete 删除用户
	// Delete(ctx context.Context, req *apicommon.IdRequest) (*apicommon.EmptyReply, error)
	// Info 获取个人信息
	Info(ctx context.Context, id string) (model.Employee, error)
	// List 获取用户列表
	// List(ctx context.Context, req *apiweb.ListUserRequest) (*apiweb.ListUserReply, error)
	// // Update 更新用户
	// Update(ctx context.Context, req *apiweb.UserRequest) (*apicommon.EmptyReply, error)
	// // UpdateInfo 更新个人信息
	// UpdateInfo(ctx context.Context, req *apiweb.UpdateInfoRequest) (*apicommon.EmptyReply, error)
	GetPassword(ctx context.Context, username string) (model.User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	authRepo AuthRepo
	repo     UserRepo
	log      *log.Helper
}

func NewUserUseCase(repo UserRepo, authRepo AuthRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		authRepo: authRepo,
		repo:     repo,
		log:      log.NewHelper(log.With(logger, "web", "biz/user")),
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
func (uc *UserUsecase) Info(ctx context.Context, req *apiweb.UserRequest) (*apiweb.UserReply, error) {

	cid := ctx.Value(ContextCidKey{}).(string)
	id, err := uc.authRepo.GetIdByCid(ctx, cid)
	if err != nil {
		uc.log.Error(err)
	}
	// return uc.repo.Info(ctx, req)
	result, err := uc.repo.Info(ctx, id)
	if err != nil {
		uc.log.Error(err.Error())
	}

	return &apiweb.UserReply{
		Data: &apiweb.UserReply_User{
			Id:             result.ID,
			WorkId:         result.WorkId.String,
			Name:           result.Name.String,
			Gender:         result.Gender.String,
			Picture:        result.Picture.String,
			IdCard:         result.IdCard.String,
			IdCardPicture:  result.IdCardPicture.String,
			Wedlock:        result.Wedlock.String,
			NationId:       result.NationId.Int64,
			NativePlace:    result.NativePlace.String,
			PoliticId:      result.PoliticId.Int64,
			Email:          result.Email.String,
			Phone:          result.Phone.String,
			Address:        result.Address.String,
			DepartmentId:   result.DepartmentId.Int64,
			JobLevelId:     result.JobLevelId.Int64,
			PosId:          result.PosId.Int64,
			EngageForm:     result.EngageForm.String,
			TiptopDegree:   result.TiptopDegree.String,
			Specialty:      result.Specialty.String,
			School:         result.School.String,
			BeginDate:      result.BeginDate.Time.Format(time.DateTime),
			WorkStatus:     result.WorkStatus.Int32,
			ContractTerm:   result.ContractTerm.Int32,
			ConversionTime: result.ConversionTime.Time.Format(time.DateTime),
			NotWorkDate:    result.NotWorkDate.Time.Format(time.DateTime),
			BeginContract:  result.BeginContract.Time.Format(time.DateTime),
			EndContract:    result.EndContract.Time.Format(time.DateTime),
			WorkAge:        result.WorkAge.Int32,
		},
	}, nil
}

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
