package httperrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ErrBadRequest    = "Bad request"
	ErrAlreadyExists = "Already exists"
	ErrNoSuchUser    = "User not found"
)

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewBadRequest(message string) *APIError {
	return &APIError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInternal(message string) *APIError {
	return &APIError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func Error(c echo.Context, apiErr *APIError) error {
	return c.JSON(apiErr.Code, apiErr)
}
