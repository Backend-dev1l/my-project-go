package db

import (
	"context"

	"log/slog"
	"project-go/internal/infrastructure/db/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"go.uber.org/fx"
)

func NewPool(cfg *config.Config, log *slog.Logger) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		log.Error("Failed to parse DSN", "error", err)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Error("Failed to create connection pool", "error", err)

	}

	if err := db.Ping(context.Background()); err != nil {
		log.Error("Failed to ping database", "error", err)
	}

	log.Info("Successfully connected to database")

	return db, nil

}

func StartPostgres(lc fx.Lifecycle, pool *pgxpool.Pool, logger *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := pool.Ping(ctx); err != nil {
				logger.Error("Failed to ping database", "error", err)
				return err
			}
			logger.Info("Connected to database")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("close database pool")
			pool.Close()
			return nil
		},
	})
}
