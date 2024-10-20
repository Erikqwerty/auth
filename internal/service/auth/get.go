package authservice

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Get - получение информации о пользователе
func (s *service) Get(ctx context.Context, email string) (*model.User, error) {
	return s.authRepository.ReadUser(ctx, email)
}
