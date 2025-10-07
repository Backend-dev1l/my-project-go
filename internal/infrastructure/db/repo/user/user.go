package user

import (
	"context"
	"fmt"
	"log/slog"
	"project-go/internal/domain/entities"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	log    *slog.Logger
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewStorage(log *slog.Logger, getter *trmpgx.CtxGetter, db *pgxpool.Pool) (*UserStorage, error) {
	return &UserStorage{log: log, getter: getter, db: db}, nil
}

func (s *UserStorage) Update(ctx context.Context, user *entities.User) error {
	s.log.Info("Starting SaveUser")

	q := s.getter.DefaultTrOrDB(ctx, s.db)

	_, err := q.Exec(ctx, `UPDATE identity.users SET balance = $2, revenue = $3 WHERE id = $1`, user.ID, user.Balance, user.Revenue)
	if err != nil {
		s.log.Error("Failed to execute update query", "error", err)
		return fmt.Errorf("execute update query: %w", err)
	}

	return nil
}

func (s *UserStorage) GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	s.log.Info("Starting GetUserByID")

	var user entities.User

	q := s.getter.DefaultTrOrDB(ctx, s.db)
	err := q.QueryRow(ctx, "SELECT id, balance, revenue FROM identity.users WHERE id = $1 FOR UPDATE", userID).Scan(&user.ID, &user.Balance, &user.Revenue)
	if err != nil {
		s.log.Error("Failed to execute select query", "error", err)
		return nil, fmt.Errorf("execute select query: %w", err)
	}
	return &user, nil
}

func (s *UserStorage) CreateUser(ctx context.Context, user *entities.User) error {
	s.log.Info("Starting create user", "user_id", user.ID)
	q := s.getter.DefaultTrOrDB(ctx, s.db)

	_, err := q.Exec(ctx, "INSERT INTO identity.users(id, balance, revenue) VALUES($1, $2, $3)", user.ID, user.Balance, user.Revenue)
	if err != nil {
		s.log.Error("can't create user", "error", err)
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
