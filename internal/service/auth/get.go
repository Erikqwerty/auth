package authservice

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Get - получение информации о пользователе
func (s *service) Get(ctx context.Context, email string) (*model.User, error) {

	user := &model.User{}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTX error
		user, errTX = s.authRepository.ReadUser(ctx, email)
		if errTX != nil {
			return errTX
		}
		//
		return nil
	})

	return user, err
}
