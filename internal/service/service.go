package service

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// AuthService - интерфейс сервисного слоя
type AuthService interface {
	// CreateUser - создание пользователя
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	// GetUser - получение информации о пользователе
	GetUser(ctx context.Context, email string) (*model.User, error)
	// UpdateUser - обновить информацию о пользователе
	UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser - удалить пользователя
	DeleteUser(ctx context.Context, id int64) error
}
