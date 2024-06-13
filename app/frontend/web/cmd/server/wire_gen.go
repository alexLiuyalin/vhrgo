// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"vhrgo/app/frontend/web/internal/biz"
	"vhrgo/app/frontend/web/internal/conf"
	"vhrgo/app/frontend/web/internal/data"
	"vhrgo/app/frontend/web/internal/server"
	"vhrgo/app/frontend/web/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, microservices *conf.Microservices, management *conf.Management, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	universalClient := data.NewRedis(confServer, confData, logger)
	client := server.NewClient(management)
	registrar := server.NewRegistrar(client, management)
	dataData, cleanup, err := data.NewData(db, universalClient, registrar, confServer, microservices, management, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	captchaUseCase := biz.NewCaptchaUseCase(userRepo, logger)
	captchaService := service.NewCaptchaService(captchaUseCase, logger)
	userUsecase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUsecase, logger)
	hrRepo := data.NewHrRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(logger, hrRepo)
	authService := service.NewAuthService(authUseCase, logger)
	menuRepo := data.NewMenuRepo(dataData, logger)
	menuUsecase := biz.NewMenuUseCase(menuRepo, logger)
	menuService := service.NewMenuService(menuUsecase, logger)
	employeeRepo := data.NewEmployeeRepo(dataData, logger)
	employeeUsecase := biz.NewEmployeeUseCase(employeeRepo, logger)
	employeeService := service.NewEmployeeService(employeeUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, captchaService, userService, authService, menuService, employeeService, logger)
	app := newApp(logger, confServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
