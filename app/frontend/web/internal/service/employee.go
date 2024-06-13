package service

import (
	"context"

	apiweb "vhrgo/api/frontend/web/v1"
	"vhrgo/app/frontend/web/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type EmployeeService struct {
	apiweb.EmployeeHTTPServer

	uc  *biz.EmployeeUsecase
	log *log.Helper
}

func NewEmployeeService(
	uc *biz.EmployeeUsecase,
	logger log.Logger) *EmployeeService {
	return &EmployeeService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "core", "service/menu")),
	}
}

func (s *EmployeeService) List(ctx context.Context,
	req *apiweb.GetBasicRequest) (*apiweb.EmployeeReply, error) {
	res, err := s.uc.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return &apiweb.EmployeeReply{
		Data: res,
	}, nil
}

func (s *EmployeeService) Nations(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetNationsReply, error) {
	res, err := s.uc.Nations(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetNationsReply{
		Nations: res,
	}, nil
}

func (s *EmployeeService) Joblevels(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetJoblevelsReply, error) {
	res, err := s.uc.Joblevels(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetJoblevelsReply{
		Data: res,
	}, nil
}

func (s *EmployeeService) Politicsstatus(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetPoliticsstatusReply, error) {
	res, err := s.uc.Politicsstatus(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetPoliticsstatusReply{
		Politicsstatus: res,
	}, nil
}

func (s *EmployeeService) Deps(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetDepsReply, error) {
	res, err := s.uc.Deps(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetDepsReply{
		Data: res,
	}, nil
}

func (s *EmployeeService) Positions(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetPositionsReply, error) {
	res, err := s.uc.Positions(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetPositionsReply{
		Data: res,
	}, nil
}

func (s *EmployeeService) MaxWorkID(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetMaxWorkIDReply, error) {
	res, err := s.uc.MaxWorkID(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetMaxWorkIDReply{
		MaxWorkId: res,
	}, nil
}

func (s *EmployeeService) Export(ctx context.Context,
	req *apiweb.GetEmptyRequest) (*apiweb.GetExportReply, error) {
	res, err := s.uc.Export(ctx)
	if err != nil {
		return nil, err
	}
	return &apiweb.GetExportReply{
		Export: res,
	}, nil
}
