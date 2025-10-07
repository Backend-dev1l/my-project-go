package usecase

import (
	"context"
	"project-go/internal/application/dto"

	"github.com/shopspring/decimal"
)

type UserUsecases interface {
	Deposit(ctx context.Context, request *dto.Deposit) error
	Reserve(ctx context.Context, request *dto.Reserve) error
	ConfirmRevenue(ctx context.Context, request *dto.ConfirmRevenue) error
	GetBalance(ctx context.Context, request *dto.GetBalance) (decimal.Decimal, error)
	Create(ctx context.Context, request *dto.Create) error
}
