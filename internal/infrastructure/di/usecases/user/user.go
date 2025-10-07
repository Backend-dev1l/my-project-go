package user

import (
	"project-go/internal/application/user/interfaces/usecase"
	"project-go/internal/application/user/usecases"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		usecases.NewUserUsecases,
		fx.As(new(usecase.UserUsecases)),
	),
)
