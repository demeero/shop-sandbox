package config

type Config struct {
	GRPC GRPC
	SQL  SQL
}

type GRPC struct {
	Port int
}

type SQL struct {
	Driver     string
	Datasource string
}
