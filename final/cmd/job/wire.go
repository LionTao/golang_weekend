//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"

	"golang_weekend/final/internal/biz"
	"golang_weekend/final/internal/cache"
	"golang_weekend/final/internal/conf"
	"golang_weekend/final/internal/data"
	"golang_weekend/final/internal/mq"
	"golang_weekend/final/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Data, *conf.Cache, *conf.MessageQueue, log.Logger) (func(context.Context) error, func(), error) {
	panic(wire.Build(
		cache.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.JobProviderSet,
		mq.ProviderSet,
		newApp))
}
