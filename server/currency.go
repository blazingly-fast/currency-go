package server

import (
	"context"
	"log"

	pb "github.com/blazingly-fast/currency-go/protos"
)

type Currency struct {
  pb.UnimplementedCurrencyServer
}

func NewCurrency() *Currency{
  return &Currency{}
}

func (c*Currency) GetRate(ctx context.Context, rr* pb.RateRequest) (*pb.RateResponse, error) {
  log.Println("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
  return &pb.RateResponse{Rate: 0.5}, nil
}



