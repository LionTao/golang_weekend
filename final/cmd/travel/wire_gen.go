// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"golang_weekend/final/internal/conf"
	"golang_weekend/final/internal/server/travel"
	"golang_weekend/final/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(server *conf.Server, traceTracerProvider trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	travelService := service.NewTravelService(logger)
	httpServer := travel.NewHTTPServer(server, traceTracerProvider, travelService)
	grpcServer := travel.NewGRPCServer(server, traceTracerProvider, travelService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
