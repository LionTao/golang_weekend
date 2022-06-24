//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"golang_weekend/final/internal/conf"
	shop "golang_weekend/final/internal/server/shop"
	"golang_weekend/final/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Service, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		shop.ProviderSet,
		service.ShopProviderSet,
		newApp))
}
