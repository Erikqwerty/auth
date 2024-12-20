package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// UpdateUser - обновить информацию о пользователе
func (s *service) UpdateUser(ctx context.Context, user *model.UpdateUser) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.authRepository.UpdateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.authRepository.CreateLog(ctx, &model.Log{
			ActionType:    actionTypeUpdate,
			ActionDetails: details(ctx),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
