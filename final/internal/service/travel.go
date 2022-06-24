package service

import (
	"context"

	"golang_weekend/final/api/common/v1"
	ts "golang_weekend/final/api/travel/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func NewTravelService(logger log.Logger) *TravelService {
	return &TravelService{
		log: log.NewHelper(logger),
	}
}

func (t *TravelService) Query(ctx context.Context, req *ts.TravelRequest) (*ts.TravelResponse, error) {
	t.log.Infof("input data %v", req)
	return &ts.TravelResponse{Result: &common.Result{Code: "0"},
		OrgAirport:     req.OrgAirport,
		ArrAirport:     req.ArrAirport,
		FlightDatetime: req.FlightDatetime,
		TravelMessage:  "欢迎来到Go的世界!!!",
	}, nil
}
