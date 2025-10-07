package repo

import (
	"context"
	"project-go/internal/domain/entities"
)

type ReservStorage interface {
	Save(ctx context.Context, r *entities.Reservation) error
	GetByOrderID(ctx context.Context, orderID int64) (*entities.Reservation, error)
	Delete(ctx context.Context, orderID int64) error
}
