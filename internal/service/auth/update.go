package authservice

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// Update - обновить информацию о пользователе
func (s *service) Update(ctx context.Context, user *model.User) error {
	return s.authRepository.UpdateUser(ctx, user)
}
