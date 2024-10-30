package repository

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// AuthRepository - интерфейс для взаимодействия с базой данных.
// Описывает основные операции CRUD
type AuthRepository interface {
	// CreateUser - создает нового пользователя в базе данных
	CreateUser(ctx context.Context, user *model.CreateUser) (int64, error)

	// ReadUser - получает информацию о пользователе
	ReadUser(ctx context.Context, email string) (*model.UserInfo, error)

	// UpdateUser - обновляет информацию о пользователе
	UpdateUser(ctx context.Context, user *model.UpdateUser) error

	// DeleteUser - удаляет пользователя
	DeleteUser(ctx context.Context, id int64) error

	RepoLoger
}

// RepoLoger - интерфейс для взаимодействия с базой данных таблицой логов.
type RepoLoger interface {
	// CreateLog - Записывает действие в бд в лог таблицу
	CreateLog(ctx context.Context, log *model.Log) error
}

type UserCache interface {
	SetUser(ctx context.Context, email string, user *model.UserCache) error
	GetUser(ctx context.Context, email string) (*model.UserCache, error)
}
