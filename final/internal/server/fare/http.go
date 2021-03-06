package fare

import (
	v1 "golang_weekend/final/api/fare/v1"
	"golang_weekend/final/internal/conf"
	"golang_weekend/final/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, tracer trace.TracerProvider, fare *service.FareService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tracer)),
			// logging.Server(logger),
			// validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterFareServiceHTTPServer(srv, fare)
	return srv
}
