package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DB     Postgres
	Server Server
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	SSLMode  string
	DBName   string
}

type Server struct {
	Port int
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
