package services

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Debug bool `env:"DEBUG" env-default:"false"`
	DB    struct {
		Name     string `env:"DB_NAME"`
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
	}
	Minio struct {
		Endpoint  string `env:"MINIO_ENDPOINT" env-default:"localhost:9000"`
		AccessKey string `env:"MINIO_ACCESS_KEY"`
		SecretKey string `env:"MINIO_SECRET_KEY"`
		UseSSL    bool   `env:"MINIO_USE_SSL" env-default:"false"`
	}
}

func NewConfig() (*Config, error) {
	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
