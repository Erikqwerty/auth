package service

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// AuthService - интерфейс для управления пользователями.
// Обеспечивает функционал создания, получения, обновления и удаления пользователей.
type AuthService interface {
	// CreateUser - регистрирует нового пользователя в системе.
	CreateUser(ctx context.Context, user *model.CreateUser) (int64, error)

	// GetUser - возвращает информацию о пользователе по email.
	GetUser(ctx context.Context, email string) (*model.UserInfo, error)

	// UpdateUser - обновляет данные существующего пользователя.
	UpdateUser(ctx context.Context, user *model.UpdateUser) error

	// DeleteUser - удаляет пользователя по его идентификатору.
	DeleteUser(ctx context.Context, id int64) error
}
