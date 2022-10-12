package main

import (
	"log"
	"net"

	pb "github.com/blazingly-fast/currency-go/protos"
	"github.com/blazingly-fast/currency-go/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){

  gs := grpc.NewServer()
  cs := server.NewCurrency()
  pb.RegisterCurrencyServer(gs, cs)

  reflection.Register(gs)

  l, err := net.Listen("tcp", ":9092")
  log.Println("server listening at: ","address", l.Addr())
  if err !=nil {
    log.Fatal("Unable to listen", "error", err)
  }
  gs.Serve(l)
}
