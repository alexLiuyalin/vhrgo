package biz

import (
	"context"
	"time"

	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type EmployeeRepo interface {
	List(ctx context.Context, req *apiweb.GetBasicRequest) ([]*model.Employee, error)
	Nations(ctx context.Context) ([]string, error)
	Joblevels(ctx context.Context) ([]model.JobLevel, error)
	Politicsstatus(ctx context.Context) ([]string, error)
	Deps(ctx context.Context) ([]model.Department, error)
	Positions(ctx context.Context) ([]model.Position, error)
	MaxWorkID(ctx context.Context) (string, error)
	Export(ctx context.Context) (string, error)
}

type EmployeeUsecase struct {
	repo EmployeeRepo
	log  *log.Helper
}

func NewEmployeeUseCase(repo EmployeeRepo, logger log.Logger) *EmployeeUsecase {
	return &EmployeeUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "web", "biz/user")),
	}
}

func (uc *EmployeeUsecase) List(ctx context.Context, req *apiweb.GetBasicRequest) ([]*apiweb.EmployeeData, error) {
	res, err := uc.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	var employees []*apiweb.EmployeeData
	for _, v := range res {
		res := &apiweb.EmployeeData{
			Id:             v.ID,
			WorkId:         v.WorkId.String,
			Name:           v.Name.String,
			Gender:         v.Gender.String,
			Picture:        v.Picture.String,
			IdCard:         v.IdCard.String,
			IdCardPicture:  v.IdCardPicture.String,
			Wedlock:        v.Wedlock.String,
			NationId:       v.NationId.Int64,
			NativePlace:    v.NativePlace.String,
			PoliticId:      v.PoliticId.Int64,
			Email:          v.Email.String,
			Phone:          v.Phone.String,
			Address:        v.Address.String,
			DepartmentId:   v.DepartmentId.Int64,
			JobLevelId:     v.JobLevelId.Int64,
			PosId:          v.PosId.Int64,
			EngageForm:     v.EngageForm.String,
			TiptopDegree:   v.TiptopDegree.String,
			Specialty:      v.Specialty.String,
			School:         v.School.String,
			BeginDate:      v.BeginDate.Time.Format(time.DateOnly),
			WorkStatus:     v.WorkStatus.Int32,
			ContractTerm:   v.ContractTerm.Int32,
			ConversionTime: v.ConversionTime.Time.Format(time.DateOnly),
			NotWorkDate:    v.NotWorkDate.Time.Format(time.DateOnly),
			BeginContract:  v.BeginContract.Time.Format(time.DateOnly),
			EndContract:    v.EndContract.Time.Format(time.DateOnly),
			WorkAge:        v.WorkAge.Int32,
		}
		employees = append(employees, res)
	}
	return employees, nil
}

func (uc *EmployeeUsecase) Nations(ctx context.Context) ([]string, error) {
	return uc.repo.Nations(ctx)
}

func (uc *EmployeeUsecase) Joblevels(ctx context.Context) ([]*apiweb.Joblevels, error) {
	res, err := uc.repo.Joblevels(ctx)
	if err != nil {
		return nil, err
	}
	var list []*apiweb.Joblevels
	for _, v := range res {
		list = append(list, &apiweb.Joblevels{
			Id:         v.ID,
			Name:       v.Name.String,
			TitleLevel: v.TitleLevel.String,
			CreateDate: v.CreateDate.Time.Format(time.DateOnly),
			Enabled:    v.Enabled.Bool,
		})
	}
	return list, nil
}

func (uc *EmployeeUsecase) Politicsstatus(ctx context.Context) ([]string, error) {
	return uc.repo.Politicsstatus(ctx)
}

func (uc *EmployeeUsecase) Deps(ctx context.Context) ([]*apiweb.Department, error) {
	res, err := uc.repo.Deps(ctx)
	if err != nil {
		return nil, err
	}

	var list []*apiweb.Department
	for _, v := range res {
		list = append(list, &apiweb.Department{
			Id:       v.ID,
			Name:     v.Name.String,
			ParentId: v.ParentID.Int64,
			DepPath:  v.DepPath.String,
			Enabled:  v.Enabled.Bool,
			IsParent: v.IsParent.Bool,
		})
	}
	return list, nil
}

func (uc *EmployeeUsecase) Positions(ctx context.Context) ([]*apiweb.Position, error) {
	res, err := uc.repo.Positions(ctx)
	if err != nil {
		return nil, err
	}
	var list []*apiweb.Position
	for _, v := range res {
		list = append(list, &apiweb.Position{
			Id:         v.ID,
			Name:       v.Name.String,
			CreateDate: v.CreateDate.Time.Format(time.DateOnly),
			Enabled:    v.Enabled.Bool,
		})
	}
	return list, nil
}

func (uc *EmployeeUsecase) MaxWorkID(ctx context.Context) (string, error) {
	return uc.repo.MaxWorkID(ctx)
}

func (uc *EmployeeUsecase) Export(ctx context.Context) (string, error) {
	return uc.repo.Export(ctx)
}
