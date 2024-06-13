//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"vhrgo/app/frontend/web/internal/biz"
	"vhrgo/app/frontend/web/internal/conf"
	"vhrgo/app/frontend/web/internal/data"
	"vhrgo/app/frontend/web/internal/server"
	"vhrgo/app/frontend/web/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Microservices, *conf.Management, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
