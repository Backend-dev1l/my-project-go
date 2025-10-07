package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Deposit struct {
	UserID uuid.UUID       `json:"user_id" validate:"required,uuid_not_nil"`
	Amount decimal.Decimal `json:"amount"  validate:"required,decimal_gt0"`
}

type Reserve struct {
	UserID  uuid.UUID       `json:"user_id"  validate:"required,uuid_not_nil"`
	Amount  decimal.Decimal `json:"amount"   validate:"required,decimal_gt0"`
	OrderID int64           `json:"order_id" validate:"required"`
}

type ConfirmRevenue struct {
	UserID  uuid.UUID `json:"user_id"  validate:"required,uuid_not_nil"`
	OrderID int64     `json:"order_id" validate:"required"`
}

type GetBalance struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid_not_nil"`
}

type Create struct {
	UserID  uuid.UUID       `json:"user_id"  validate:"required,uuid_not_nil"`
	Revenue decimal.Decimal `json:"revenue"`
	Balance decimal.Decimal `json:"balance"  validate:"required,decimal_gt0"`
}
