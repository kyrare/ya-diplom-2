package services

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Debug bool `env:"DEBUG" env-default:"false"`
	DB    struct {
		Name     string `env:"DB_NAME"`
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
	}
	GRPC struct {
		Address string `env:"GRPC_ADDRESS" env-default:"localhost:50051"`
	}
	Minio struct {
		Endpoint  string `env:"MINIO_ENDPOINT" env-default:"localhost:9000"`
		AccessKey string `env:"MINIO_ACCESS_KEY"`
		SecretKey string `env:"MINIO_SECRET_KEY"`
		UseSSL    bool   `env:"MINIO_USE_SSL" env-default:"false"`
	}
	Jwt struct {
		Secret   string        `env:"JWT_TOKEN_SECRET"`
		Duration time.Duration `env:"JWT_TOKEN_DURATION" env-default:"48h"` // Valid time units ns, us, ms, s, m, h
	}
}

func NewConfig() (*Config, error) {
	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
