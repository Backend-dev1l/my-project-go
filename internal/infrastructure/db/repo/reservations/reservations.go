package reservations

import (
	"context"
	"fmt"
	"log/slog"
	"project-go/internal/domain/entities"
	repoErrors "project-go/internal/infrastructure/db/repo/repo_errors"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReservStorage struct {
	db     *pgxpool.Pool
	getter *trmpgx.CtxGetter
	log    *slog.Logger
}

func NewReservStorage(db *pgxpool.Pool, getter *trmpgx.CtxGetter, log *slog.Logger) *ReservStorage {
	return &ReservStorage{db: db, log: log, getter: getter}
}

func (s *ReservStorage) Save(ctx context.Context, r *entities.Reservation) error {
	s.log.Info("Starting SaveReservation")

	q := s.getter.DefaultTrOrDB(ctx, s.db)
	_, err := q.Exec(ctx, "INSERT INTO identity.reservations(order_id, user_id, amount) VALUES($1, $2, $3)", r.OrderID, r.UserID, r.Amount)
	if err != nil {
		if repoErrors.IsUniqueViolation(err) {
			s.log.Error("Reservation already exists (unique violation)", "orderID", r.OrderID)
			return fmt.Errorf("reservation already exists: %w", err)
		}
		s.log.Error("Failed to execute insert query", "error", err)
		return fmt.Errorf("execute insert query: %w", err)
	}
	return nil
}

func (s *ReservStorage) Delete(ctx context.Context, orderID int64) error {
	s.log.Info("Starting DeleteReservation")
	q := s.getter.DefaultTrOrDB(ctx, s.db)
	_, err := q.Exec(ctx, "DELETE FROM identity.reservations WHERE order_id = $1", orderID)
	if err != nil {
		s.log.Error("Failed to execute delete query", "error", err)
		return fmt.Errorf("execute delete query: %w", err)
	}
	return nil
}

func (s *ReservStorage) GetByOrderID(ctx context.Context, orderID int64) (*entities.Reservation, error) {
	s.log.Info("Starting GetReservationByOrderID")

	var res entities.Reservation

	q := s.getter.DefaultTrOrDB(ctx, s.db)
	err := q.QueryRow(ctx, "SELECT order_id, user_id, amount FROM identity.reservations WHERE order_id = $1 FOR UPDATE", orderID).
		Scan(&res.OrderID, &res.UserID, &res.Amount)
	if err != nil {
		s.log.Error("Failed to execute select query", "error", err)
		return nil, fmt.Errorf("execute select query: %w", err)
	}
	return &res, nil
}
