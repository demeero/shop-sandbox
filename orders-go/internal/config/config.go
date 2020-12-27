package config

type Config struct {
	GRPC       GRPC
	Repository Repository
}

type GRPC struct {
	Port int
}

type Repository struct {
	Name       string
	Datasource string
}
