package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"
	"github.com/demeero/shop-sandbox/users/config"
	"github.com/demeero/shop-sandbox/users/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg config.Config
	err := envconfig.Process("users", &cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterUserServiceServer(grpcServer, service.New())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve GRPC: %v", err)
	}
}
