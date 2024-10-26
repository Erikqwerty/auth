package auth

import (
	"context"
)

// Delete - удалить пользователя
func (s *service) DeleteUser(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTX := s.authRepository.DeleteUser(ctx, id)
		if errTX != nil {
			return errTX
		}

		if errTx := s.writeLog(ctx, actionTypeDelete); errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
