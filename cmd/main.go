package main

import (
	"project-go/internal/infrastructure/di"
	"project-go/internal/presentation/di/api"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		di.Module,
		api.Module,
	)
	app.Run()
}
