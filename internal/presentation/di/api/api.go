package api

import (
	"project-go/internal/presentation/di/api/handlers"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.api",
	handlers.Module,
)
