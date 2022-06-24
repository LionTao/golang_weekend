//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"golang_weekend/final/internal/conf"
	travel "golang_weekend/final/internal/server/travel"
	"golang_weekend/final/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		travel.ProviderSet,
		service.TravelProviderSet,
		newApp))
}
