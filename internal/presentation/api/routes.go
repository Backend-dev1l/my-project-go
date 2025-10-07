package api

import (
	"project-go/internal/presentation/api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handler *handlers.Handler) {
	v1 := e.Group("api/v1")

	v1.POST("/deposit", handler.Deposit)
	v1.POST("/create", handler.Create)
	v1.POST("/confirm-revenue", handler.ConfirmRevenue)
	v1.POST("/reserve", handler.Reserve)
	v1.GET("/balance", handler.GetBalance)

}
