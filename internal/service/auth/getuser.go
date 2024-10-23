package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// GetUser - получение информации о пользователе
func (s *service) GetUser(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error

		user, errTX = s.authRepository.ReadUser(ctx, email)
		if errTX != nil {
			return errTX
		}

		if errTx := s.createLog(ctx, actionTypeGet); errTx != nil {
			return errTx
		}

		return nil
	})

	return user, err
}