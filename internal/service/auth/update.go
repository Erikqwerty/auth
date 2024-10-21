package authservice

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Update - обновить информацию о пользователе
func (s *service) Update(ctx context.Context, user *model.User) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTX := s.authRepository.UpdateUser(ctx, user)
		if errTX != nil {
			return errTX
		}
		//
		return nil
	})

	return err
}
