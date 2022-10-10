package main

import (
	"net"

	pb "github.com/blazingly-fast/currency-go/protos"
	"github.com/blazingly-fast/currency-go/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
  log := hclog.Default()

  gs := grpc.NewServer()
  cs := server.NewCurrency()
  pb.RegisterCurrencyServer(gs, cs)

  reflection.Register(gs)

  l, err := net.Listen("tcp", ":9092")
  log.Info("server listening at: ","address", l.Addr())
  if err !=nil {
    log.Error("Unable to listen", "error", err)
  }
  gs.Serve(l)
}
