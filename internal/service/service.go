package service

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// AuthService - интерфейс сервисного слоя
type AuthService interface {
	// Create - создание пользователя
	Create(ctx context.Context, user *model.User) (int64, error)
	// Get - получение информации о пользователе
	Get(ctx context.Context, email string) (*model.User, error)
	// Update - обновить информацию о пользователе
	Update(ctx context.Context, user *model.User) error
	// Delete - удалить пользователя
	Delete(ctx context.Context, id int64) error
}
