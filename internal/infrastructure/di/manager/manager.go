package manager

import (
	"project-go/internal/infrastructure/db/manager"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		manager.NewPgxTrManager,
		manager.NewPgxTrManagerDefaultCtxGetter,
	),
)
