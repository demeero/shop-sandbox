package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/demeero/shop-sandbox/users/config"
	"github.com/demeero/shop-sandbox/users/internal/service"
	"github.com/demeero/shop-sandbox/users/internal/storage/mongo"
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

	repo, err := mongo.New(context.Background(), options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		log.Fatal(err)
	}
	s := service.New(repo)
	pb.RegisterUserServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve GRPC: %v", err)
	}
}
