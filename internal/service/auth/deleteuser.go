package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Delete - удалить пользователя
func (s *service) DeleteUser(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.authRepository.DeleteUser(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.authRepository.CreateLog(ctx, &model.Log{
			ActionType:    actionTypeDelete,
			ActionDetails: details(ctx),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
