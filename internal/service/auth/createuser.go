package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Create - создание пользователя
func (s *service) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	if err := prepareUserForCreate(user); err != nil {
		return 0, err
	}

	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.authRepository.CreateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		if errTx := s.createLog(ctx, actionTypeCreate); errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
