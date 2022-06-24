package service

import (
	"context"

	"golang_weekend/final/api/common/v1"
	f "golang_weekend/final/api/fare/v1"
	sp "golang_weekend/final/api/shop/v1"
	t "golang_weekend/final/api/travel/v1"
	"golang_weekend/final/internal/client"
	"golang_weekend/final/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
)

func NewShoppingService(conf *conf.Service, logger log.Logger) *ShoppingService {
	return &ShoppingService{
		fare:   client.NewFareClient(conf),
		travel: client.NewTravelClient(conf),
		log:    log.NewHelper(logger),
	}
}

func (s *ShoppingService) Shopping(ctx context.Context, req *sp.ShoppingRequest) (*sp.ShoppingResponse, error) {
	s.log.Infof("input data %v", req)

	g, ctx := errgroup.WithContext(ctx)

	var (
		trsp *t.TravelResponse
		e    error
	)

	g.Go(func() error {
		treq := &t.TravelRequest{OrgAirport: req.ArrAirport, ArrAirport: req.ArrAirport, FlightDatetime: req.FlightDatetime}
		trsp, e = s.travel.Query(ctx, treq)
		return e
	})

	var (
		prsp *f.PricingResponse
		ep   error
	)
	g.Go(func() error {
		preq := &f.PricingRequest{OrgAirport: req.OrgAirport, ArrAirport: req.ArrAirport, FlightDatetime: req.FlightDatetime, PassageType: req.PassageType}
		prsp, ep = s.fare.Pricing(ctx, preq)
		return ep
	})

	var (
		pcrsp *f.PriceCalendarResponse
		epc   error
	)
	g.Go(func() error {
		pcreq := &f.PriceCalendarRequest{OrgAirport: req.OrgAirport, ArrAirport: req.ArrAirport, FlightDatetime: req.FlightDatetime, PassageType: req.PassageType}
		pcrsp, epc = s.fare.PriceCalendar(ctx, pcreq)
		return epc
	})

	if err := g.Wait(); err != nil {
		s.log.Errorf("error: %v\n", err)
		return &sp.ShoppingResponse{Result: &common.Result{Code: "1"}}, err
	}

	res := &sp.ShoppingResponse{Result: &common.Result{Code: "0"}}
	res.TravelMessage = trsp.TravelMessage
	res.PriceCalendar = pcrsp.PriceCalendar
	res.OrgAirport = req.OrgAirport
	res.ArrAirport = req.ArrAirport
	res.FlightDatetime = req.FlightDatetime
	res.PassageType = req.PassageType
	res.Amount = prsp.Amount
	return res, nil
}
