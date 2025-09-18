package config

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Deps struct {
	Postgres *pgxpool.Pool
}

func WithPostgres(
	login string,
	password string,
	host string,
	port string,
	database string,
) *Deps {
	dsn := "postgres://" + login + ":" + password + "@" + host + ":" + port + "/" + database
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(context.TODO(), config)
	if err != nil {
		panic(err)
	}

	return &Deps{
		Postgres: pool,
	}
}
