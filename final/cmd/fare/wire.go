//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"golang_weekend/final/internal/biz"
	"golang_weekend/final/internal/cache"
	"golang_weekend/final/internal/conf"
	"golang_weekend/final/internal/data"
	"golang_weekend/final/internal/mq"
	server "golang_weekend/final/internal/server/fare"
	"golang_weekend/final/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Cache, *conf.MessageQueue, *conf.Service, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		cache.ProviderSet,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		mq.ProviderSet,
		newApp))
}
