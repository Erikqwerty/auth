package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/autherrors"
	"github.com/erikqwerty/auth/internal/model"
)

// Create - создание пользователя
func (s *service) CreateUser(ctx context.Context, user *model.CreateUser) (int64, error) {
	if user == nil {
		return 0, autherrors.ErrCreateUserNil
	}

	if err := user.Validate(); err != nil {
		return 0, err
	}

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

		if errTx := s.writeLog(ctx, actionTypeCreate); errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
