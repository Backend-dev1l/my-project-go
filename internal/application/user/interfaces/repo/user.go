package repo

import (
	"context"
	"project-go/internal/domain/entities"

	"github.com/google/uuid"
)

type UserStorage interface {
	Update(ctx context.Context, user *entities.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
}
