package api

import (
	"context"
	"log/slog"
	"project-go/internal/config"

	"project-go/internal/presentation/api/handlers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewHTTPServer(logger *slog.Logger, handler *handlers.Handler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))
	RegisterRoutes(e, handler)

	return e
}

func RunHTTPServer(lc fx.Lifecycle, e *echo.Echo, cfg *config.Config, logger *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting HTTP server", "address", cfg.Address())
				if err := e.Start(cfg.Address()); err != nil {
					logger.Error("Error when start http server", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down http server")
			return e.Shutdown(ctx)
		},
	})

}
