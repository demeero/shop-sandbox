package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/demeero/shop-sandbox/orders/internal/config"
	"github.com/demeero/shop-sandbox/orders/internal/service"
	"github.com/demeero/shop-sandbox/orders/internal/storage/gorm"
	"github.com/demeero/shop-sandbox/orders/internal/storage/sql"
	"github.com/demeero/shop-sandbox/orders/internal/storage/squirrel"
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

	repo, err := createRepo(cfg.Repository)
	if err != nil {
		log.Fatalf("failed create repository: %+v", err)
	}
	s := service.New(repo)
	pb.RegisterOrderServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve GRPC: %v", err)
	}
}

func createRepo(cfg config.Repository) (service.Repository, error) {
	switch strings.ToLower(cfg.Name) {
	case "sql":
		return sql.New(cfg.Datasource)
	case "gorm":
		return gorm.New(cfg.Datasource)
	case "squirrel":
		return squirrel.New(cfg.Datasource)
	}
	return nil, errors.New("unknown repo name")
}
