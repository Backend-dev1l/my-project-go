package usecases

import (
	"context"
	"fmt"
	"log/slog"
	"project-go/internal/application/dto"
	"project-go/internal/application/user/interfaces/repo"

	"project-go/internal/domain/entities"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/shopspring/decimal"
)

type UserUsecases struct {
	userStorage     repo.UserStorage
	reservatStorage repo.ReservStorage
	logger          *slog.Logger
	tm              trm.Manager
}

func NewUserUsecases(userStorage repo.UserStorage, logger *slog.Logger, tm trm.Manager, reservatStorage repo.ReservStorage) *UserUsecases {
	return &UserUsecases{
		userStorage:     userStorage,
		logger:          logger,
		tm:              tm,
		reservatStorage: reservatStorage,
	}
}

func (u *UserUsecases) Create(ctx context.Context, request *dto.Create) error {
	u.logger.Info("start Create user usecase")

	user := &entities.User{
		ID:      request.UserID,
		Balance: request.Balance,
		Revenue: request.Revenue,
	}

	if err := u.userStorage.CreateUser(ctx, user); err != nil {
		u.logger.Info("can't create user")
		return err
	}
	return nil
}

func (u *UserUsecases) Deposit(ctx context.Context, request *dto.Deposit) error {
	u.logger.Info("Starting deposit usecase")

	return u.tm.Do(ctx, func(ctx context.Context) error {
		user, err := u.userStorage.GetUserByID(ctx, request.UserID)
		if err != nil {
			u.logger.Error("user don't created")
			return err
		}

		user.UpdateBalance(request.Amount)

		if err := u.userStorage.Update(ctx, user); err != nil {
			u.logger.Error("failed to save user balance")
			return err
		}

		return nil
	})
}

func (u *UserUsecases) Reserve(ctx context.Context, request *dto.Reserve) error {
	u.logger.Info("Starting reserve usecase")

	return u.tm.Do(ctx, func(ctx context.Context) error {
		user, err := u.userStorage.GetUserByID(ctx, request.UserID)
		if err != nil {
			u.logger.Error("Failed to get user by ID", "error", err)
			return err
		}

		if user.Balance.LessThan(request.Amount) {
			u.logger.Error("Insufficient balance", "balance", user.Balance, "amount", request.Amount)
			return fmt.Errorf("insufficient balance")
		}

		user.UpdateBalance(request.Amount)

		if err := u.userStorage.Update(ctx, user); err != nil {
			u.logger.Error("Failed to update user balance", "error", err)
			return err
		}

		reservation := &entities.Reservation{
			OrderID: request.OrderID,
			UserID:  request.UserID,
			Amount:  request.Amount,
		}

		if err := u.reservatStorage.Save(ctx, reservation); err != nil {
			u.logger.Error("Failed to save reservation", "error", err)
			return err
		}

		u.logger.Info("Reservation saved successfully")
		return nil
	})

}

func (u *UserUsecases) ConfirmRevenue(ctx context.Context, request *dto.ConfirmRevenue) error {
	u.logger.Info("Starting confirm revenue usecase")

	return u.tm.Do(ctx, func(ctx context.Context) error {
		res, err := u.reservatStorage.GetByOrderID(ctx, request.OrderID)
		if err != nil {
			u.logger.Error("Failed to get reservation by OrderID", "error", err)
			return err
		}

		user, err := u.userStorage.GetUserByID(ctx, request.UserID)
		if err != nil {
			u.logger.Error("Failed to get user by ID", "error", err)
			return err
		}

		err = u.reservatStorage.Delete(ctx, request.OrderID)
		if err != nil {
			u.logger.Error("Failed to delete reservation", "error", err)
			return err
		}

		user.UpdateRevenue(res.Amount)

		if err := u.userStorage.Update(ctx, user); err != nil {
			u.logger.Error("failed to save user revenue")
			return err
		}

		return nil
	})
}

func (u *UserUsecases) GetBalance(ctx context.Context, request *dto.GetBalance) (decimal.Decimal, error) {
	u.logger.Info("Starting GetBalance usecase")

	user, err := u.userStorage.GetUserByID(ctx, request.UserID)
	if err != nil {
		u.logger.Error("Failed to get user by ID", "error", err)
		return decimal.Zero, err
	}
	return user.Balance, nil
}
