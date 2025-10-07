package errors

import (
	"errors"
)

var (
	ErrInvalidAmount       = errors.New("amount must be positive")
	ErrReservationExists   = errors.New("reservation already exists")
	ErrNoReservation       = errors.New("no reservation found")
	ErrReservationNotFound = errors.New("reservation not found for this user")
	ErrAmountMismatch      = errors.New("amount does not match reservation")
	ErrInvalidUserID       = errors.New("invalid user ID")
	ErrInsufficientFunds   = errors.New("insufficient funds")
)
