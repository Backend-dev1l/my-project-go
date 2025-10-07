package di

import (
	"project-go/internal/infrastructure/di/db"
	"project-go/internal/infrastructure/di/logger"
	"project-go/internal/infrastructure/di/manager"

	"project-go/internal/infrastructure/di/usecases"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	db.Module,
	usecases.Module,
	logger.Module,
	manager.Module,
)
