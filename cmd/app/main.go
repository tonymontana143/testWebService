package main

import (
	"context"
	"employee/cmd/config"
	"employee/cmd/setup"
	"employee/internal/users/adapters"
	"employee/internal/users/ports"
	"employee/internal/users/service"
	"log/slog"
	"net/http"
)

func main() {
	ctx := context.Background()

	if err := setup.RunMigrations(ctx); err != nil {
		slog.Error("Migration failed", "error", err)
		return
	}

	conf, err := config.NewConfig(ctx)
	if err != nil {
		slog.Error("Unable to get configuration", "error", err)
		return
	}

	deps, err := config.WithPostgres(
		conf.Postgres.Login,
		conf.Postgres.Pass,
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.DB,
	), nil
	if err != nil {
		slog.Error("Unable to connect to Postgres", "error", err)
		return
	}

	employeeRepo := adapters.NewEmployeesRepository(deps.Postgres)
	employeeSvc := service.NewEmployeesService(employeeRepo)
	handler := ports.NewHandler(employeeSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", handler.CreateEmployee)
	slog.Info("starting HTTP server", "address", "8080")
	http.ListenAndServe(":8080", mux)

}
