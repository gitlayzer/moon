// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package houyi

import (
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/biz"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/data"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/data/repoimpl"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/houyiconf"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/server"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *houyiconf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(bootstrap)
	httpServer := server.NewHTTPServer(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData)
	cacheRepo := repoimpl.NewCacheRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, cacheRepo)
	greeterService := service.NewGreeterService(greeterUsecase)
	metric := repoimpl.NewMetricRepository(dataData)
	metricBiz := biz.NewMetricBiz(metric)
	metricService := service.NewMetricService(metricBiz)
	healthService := service.NewHealthService()
	serverServer := server.RegisterService(grpcServer, httpServer, greeterService, metricService, healthService)
	app := newApp(serverServer, logger)
	return app, func() {
		cleanup()
	}, nil
}
