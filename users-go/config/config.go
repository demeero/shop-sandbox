package config

type Config struct {
	GRPC GRPC
}

type GRPC struct {
	Port int
}
