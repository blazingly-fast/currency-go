package server

import (
	"context"

	pb "github.com/blazingly-fast/currency-go/protos"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
  log hclog.Logger
  pb.UnimplementedCurrencyServer
}

func NewCurrency() *Currency{
  return &Currency{}
}

func (c*Currency) GetRate(ctx context.Context, rr*pb.RateRequest) (*pb.RateResponse, error) {
  c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
  return &pb.RateResponse{Rate: 0.5}, nil
}



