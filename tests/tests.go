package tests

import (
	"context"
	"employee/migrations"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/peterldowns/pgtestdb"
	"github.com/peterldowns/pgtestdb/migrators/golangmigrator"
)

func GetPostgresPool(t *testing.T, ctx context.Context) (*pgxpool.Pool, error) {
	t.Helper()

	var (
		dbconf = pgtestdb.Config{
			DriverName: "pgx",
			User:       "testuser",
			Password:   "testpassword",
			Host:       "localhost",
			Port:       "5433",
			Database:   "testdb",
			Options:    "sslmode=disable",
		}
		gm     = golangmigrator.New(".", golangmigrator.WithFS(migrations.FS))
		config = pgtestdb.Custom(t, dbconf, gm)
	)

	pool, err := pgxpool.New(ctx, config.URL())
	if err != nil {
		return nil, fmt.Errorf("error on pgxpool.New(): %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error on pgxpool.Ping(): %w", err)
	}

	return pool, nil
}
