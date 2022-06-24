package service

import (
	f "golang_weekend/final/api/fare/v1"
	s "golang_weekend/final/api/shop/v1"
	t "golang_weekend/final/api/travel/v1"
	"golang_weekend/final/internal/biz"
	c "golang_weekend/final/internal/cache"
	m "golang_weekend/final/internal/mq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFareService)

var ShopProviderSet = wire.NewSet(NewShoppingService)

var TravelProviderSet = wire.NewSet(NewTravelService)

var JobProviderSet = wire.NewSet(NewJobService)

type FareService struct {
	f.UnimplementedFareServiceServer

	log   *log.Helper
	fare  *biz.FareUsecase
	cache c.Cache
	mq    m.MessageQueue
}

type ShoppingService struct {
	s.UnimplementedShopServiceServer

	log    *log.Helper
	fare   f.FareServiceHTTPClient
	travel t.TravelServiceHTTPClient
}

type TravelService struct {
	t.UnimplementedTravelServiceServer
	log *log.Helper
}

type JobService struct {
	ttl   int
	fare  *biz.FareUsecase
	log   *log.Helper
	cache c.Cache
	mq    m.MessageQueue
}
