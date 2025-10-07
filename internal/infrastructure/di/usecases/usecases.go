package usecases

import (
	"project-go/internal/infrastructure/di/usecases/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
)
