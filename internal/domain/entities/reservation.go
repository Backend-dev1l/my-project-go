package entities

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Reservation struct {
	OrderID int64
	UserID  uuid.UUID
	Amount  decimal.Decimal
}
