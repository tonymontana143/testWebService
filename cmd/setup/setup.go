package setup

import (
	"fmt"
	"os"

	"context"
	"database/sql"
	"employee/cmd/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrations(ctx context.Context) error {
	conf, err := config.NewConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Postgres.Login,
		conf.Postgres.Pass,
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create db driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+os.Getenv("MIGRATIONS_PATH"),
		conf.Postgres.DB, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}
