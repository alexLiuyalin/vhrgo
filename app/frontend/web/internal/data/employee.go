package data

import (
	"context"
	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"
	"vhrgo/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.EmployeeRepo = (*employeeRepo)(nil)

type employeeRepo struct {
	data *Data
	log  *log.Helper
}

func NewEmployeeRepo(data *Data, logger log.Logger) biz.EmployeeRepo {
	return &employeeRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "web", "data/employee")),
	}
}

func (r employeeRepo) List(ctx context.Context, req *apiweb.GetBasicRequest) (res []*model.Employee, err error) {
	// var employees []*model.Employee
	sqlStr := "select * from employee "
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}

	return
}

func (r employeeRepo) Nations(ctx context.Context) (res []string, err error) {
	sqlStr := "selet name from nation"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}
	return
}

func (r employeeRepo) Joblevels(ctx context.Context) (res []*model.JobLevel, err error) {
	sqlStr := "selet * from joblevel"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}
	return
}

func (r employeeRepo) Politicsstatus(ctx context.Context) (res []string, err error) {
	sqlStr := "selet name from politicsstatus"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}
	return
}

func (r employeeRepo) Deps(ctx context.Context) (res []*model.Department, err error) {
	sqlStr := "selet * from department"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}
	return
}

func (r employeeRepo) Positions(ctx context.Context) (res []*model.Position, err error) {
	sqlStr := "selet * from positions"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return nil, err
	}
	return
}

func (r employeeRepo) MaxWorkID(ctx context.Context) (res string, err error) {
	sqlStr := "selet max(workID) from employee"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return "", err
	}
	return
}

func (r employeeRepo) Export(ctx context.Context) (res string, err error) {
	sqlStr := "selet name from nation"
	if err = r.data.db.Select(&res, sqlStr); err != nil {
		return "", err
	}
	return
}
