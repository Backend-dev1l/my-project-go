package entities

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	ID      uuid.UUID
	Name    string
	Email   string
	Balance decimal.Decimal
	Revenue decimal.Decimal
}

func (u *User) UpdateBalance(amount decimal.Decimal) {
	if amount.Cmp(decimal.Zero) <= 0 {
		return
	}

	u.Balance = u.Balance.Add(amount)
}

func (u *User) UpdateRevenue(amount decimal.Decimal) {
	if amount.Cmp(decimal.Zero) <= 0 {
		return
	}

	u.Revenue = u.Revenue.Add(amount)
}
