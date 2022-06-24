package client

import (
	"context"

	pb "golang_weekend/final/api/fare/v1"
	"golang_weekend/final/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewFareClient(conf *conf.Service) pb.FareServiceHTTPClient {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(conf.ServiceMap["Fare"]),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewFareServiceHTTPClient(conn)
}
