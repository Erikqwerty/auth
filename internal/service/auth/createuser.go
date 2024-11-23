package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/autherrors"
	"github.com/erikqwerty/auth/internal/model"
)

// Create - создаeт пользователя и записывает событие в kafka
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

		errTx = s.authRepository.CreateLog(ctx, &model.Log{
			ActionType:    actionTypeCreate,
			ActionDetails: details(ctx),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	_, _, err = s.producer.SendMessage("CREATEUSER", user.Email)
	if err != nil {
		return id, err
	}

	return id, nil
}
