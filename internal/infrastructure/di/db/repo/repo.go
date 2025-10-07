package repo

import (
	"project-go/internal/application/user/interfaces/repo"
	"project-go/internal/infrastructure/db"
	"project-go/internal/infrastructure/db/config"
	"project-go/internal/infrastructure/db/repo/reservations"
	"project-go/internal/infrastructure/db/repo/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.LoadConfig,
		db.NewPool,
		fx.Annotate(
			reservations.NewReservStorage,
			fx.As(new(repo.ReservStorage)),
		),

		fx.Annotate(
			user.NewStorage,
			fx.As(new(repo.UserStorage)),
		),
	),
	fx.Invoke(db.StartPostgres),
)
