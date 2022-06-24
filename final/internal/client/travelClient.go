package client

import (
	"context"

	pb "golang_weekend/final/api/travel/v1"
	"golang_weekend/final/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewTravelClient(conf *conf.Service) pb.TravelServiceHTTPClient {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint(conf.ServiceMap["Travel"]),
	)
	if err != nil {
		panic(err)
	}

	return pb.NewTravelServiceHTTPClient(conn)
}
