package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Postgres struct {
		Login string `env:"USER, default=postgres"`
		Pass  string `env:"PASS, default=postgres"`
		Host  string `env:"HOST, default=localhost"`
		Port  string `env:"PORT, default=5432"`
		DB    string `env:"DB, default=employees_db"`
	} `env:", prefix=PG_"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
