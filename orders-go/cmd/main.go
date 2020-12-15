package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/demeero/shop-sandbox/orders/internal/config"
	"github.com/demeero/shop-sandbox/orders/internal/service"
)

func main() {
	var cfg config.Config
	err := envconfig.Process("orders", &cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	s := service.New(nil)
	pb.RegisterOrderServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve GRPC: %v", err)
	}
}
