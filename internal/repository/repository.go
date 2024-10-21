package repository

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// AuthRepository - интерфейс для взаимодействия с базой данных.
// Описывает основные операции CRUD
type AuthRepository interface {
	// CreateUser - создает нового пользователя в базе данных
	CreateUser(ctx context.Context, user *model.User) (int64, error)

	// ReadUser - получает информацию о пользователе
	ReadUser(ctx context.Context, email string) (*model.User, error)

	// UpdateUser - обновляет информацию о пользователе
	UpdateUser(ctx context.Context, user *model.User) error

	// DeleteUser - удаляет пользователя
	DeleteUser(ctx context.Context, id int64) error

	RepoLoger
}
type RepoLoger interface {
	// CreateLog - Записывает действие в бд в лог таблицу
	CreateLog(ctx context.Context, log *model.Log) error
}
