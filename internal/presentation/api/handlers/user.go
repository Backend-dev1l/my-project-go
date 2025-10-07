package handlers

import (
	"log/slog"
	"net/http"
	"project-go/internal/application/dto"
	"project-go/internal/application/user/interfaces/usecase"
	httpErrors "project-go/pkg/http_errors"
	customvalidator "project-go/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase  usecase.UserUsecases
	logger   *slog.Logger
	validate *validator.Validate
}

func NewHandler(usecase usecase.UserUsecases, logger *slog.Logger) *Handler {
	v := customvalidator.New()
	return &Handler{usecase: usecase, logger: logger, validate: v}
}

func (h *Handler) Create(c echo.Context) error {
	h.logger.Info("start Create handler")

	var req dto.Create
	if err := c.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", slog.Any("err", err))
		return httpErrors.Error(c, httpErrors.NewBadRequest("invalid request body"))
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("validate.struct")
		return httpErrors.Error(c, httpErrors.NewBadRequest("validate struct error"))
	}

	if err := h.usecase.Create(c.Request().Context(), &req); err != nil {
		h.logger.Error("failed to create user", slog.Any("err", err))
		return httpErrors.Error(c, httpErrors.NewInternal("handler create user error"))

	}
	return c.JSON(http.StatusOK, req)
}

func (h *Handler) Deposit(c echo.Context) error {
	h.logger.Info("Start Deposit handler")

	var req dto.Deposit
	if err := c.Bind(&req); err != nil {
		h.logger.Error("failed to bind request", slog.Any("err", err))
		return httpErrors.Error(c, httpErrors.NewBadRequest("invalid request body"))
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("validate.struct")
		return httpErrors.Error(c, httpErrors.NewBadRequest("validate struct error"))
	}

	if err := h.usecase.Deposit(c.Request().Context(), &req); err != nil {
		h.logger.Error("failed to deposit")
		return httpErrors.Error(c, httpErrors.NewInternal("handler deposit error"))
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) ConfirmRevenue(c echo.Context) error {
	h.logger.Info("start handler ConfirmRevenue")

	var req dto.ConfirmRevenue
	if err := c.Bind(&req); err != nil {
		h.logger.Error("c.Bind error")
		return httpErrors.Error(c, httpErrors.NewBadRequest("bad request"))
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("invalid request")
		return httpErrors.Error(c, httpErrors.NewBadRequest("bad request"))
	}

	if err := h.usecase.ConfirmRevenue(c.Request().Context(), &req); err != nil {
		h.logger.Error("failed to confirm revenue")
		return httpErrors.Error(c, httpErrors.NewInternal("confirm revenue error"))
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) Reserve(c echo.Context) error {
	h.logger.Info("start reserve handler")

	var req dto.Reserve
	if err := c.Bind(&req); err != nil {
		h.logger.Error("bind error")
		return httpErrors.Error(c, httpErrors.NewBadRequest("invalid request"))
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("invalid request")
		return httpErrors.Error(c, httpErrors.NewBadRequest("bad request"))
	}

	if err := h.usecase.Reserve(c.Request().Context(), &req); err != nil {
		h.logger.Error("can't reserve")
		return httpErrors.Error(c, httpErrors.NewInternal("reserve error"))
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetBalance(c echo.Context) error {
	h.logger.Info("start get balance handler")

	var req *dto.GetBalance

	if err := c.Bind(&req); err != nil {
		h.logger.Error("can't decode body")
		return httpErrors.Error(c, httpErrors.NewBadRequest("invalid body request"))
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("invalid request")
		return httpErrors.Error(c, httpErrors.NewBadRequest("validate error"))
	}

	balance, err := h.usecase.GetBalance(c.Request().Context(), req)
	if err != nil {
		h.logger.Error("failed to get balance", slog.Any("error", err))
		return httpErrors.Error(c, httpErrors.NewInternal("get balance error"))
	}

	return c.JSON(http.StatusOK, balance)

}
