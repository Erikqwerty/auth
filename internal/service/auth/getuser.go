package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// GetUser - получение информации о пользователе
func (s *service) GetUser(ctx context.Context, email string) (*model.UserInfo, error) {
	user := &model.UserInfo{}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		user, errTx = s.authRepository.ReadUser(ctx, email)
		if errTx != nil {
			return errTx
		}

		errTx = s.authRepository.CreateLog(ctx, &model.Log{
			ActionType:    actionTypeGet,
			ActionDetails: details(ctx),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return user, err
}
