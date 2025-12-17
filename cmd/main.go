package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/techies/ecom/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetEnv("GOOSE_DBSTRING", "host=localhost port=5442 user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	db, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	defer db.Close(ctx)
	slog.Info("connected to database", slog.String("dsn", cfg.db.dsn))

	api := application{
		config: cfg,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
