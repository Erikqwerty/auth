package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// UpdateUser - обновить информацию о пользователе
func (s *service) UpdateUser(ctx context.Context, user *model.UpdateUser) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTX := s.authRepository.UpdateUser(ctx, user)
		if errTX != nil {
			return errTX
		}

		if errTx := s.createLog(ctx, actionTypeUpdate); errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
