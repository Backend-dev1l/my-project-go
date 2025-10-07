package logger

import (
	"project-go/internal/infrastructure/logger"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		logger.LoadConfig,
		logger.New,
	),
	fx.Invoke(logger.Init),
)
