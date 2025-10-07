package db

import (
	"project-go/internal/infrastructure/di/db/repo"

	"go.uber.org/fx"
)

var Module = fx.Options(
	repo.Module,
)
