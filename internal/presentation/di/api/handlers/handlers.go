package handlers

import (
	"project-go/internal/config"
	"project-go/internal/presentation/api"
	"project-go/internal/presentation/api/handlers"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		handlers.NewHandler,
		config.LoadConfig,
		api.NewHTTPServer,
	),
	fx.Invoke(api.RunHTTPServer),
)
