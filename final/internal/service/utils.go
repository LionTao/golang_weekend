package service

import (
	v1 "golang_weekend/final/api/common/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func genKey(org, arr string, flightDatetime *timestamppb.Timestamp, psgType v1.PassageTypes) string {
	return org + arr + flightDatetime.AsTime().String() + psgType.String() + ":price:flight"
}
