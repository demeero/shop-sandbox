package config

type Config struct {
	GRPC  GRPC
	Mongo Mongo
}

type GRPC struct {
	Port int
}

type Mongo struct {
	URI string
}
