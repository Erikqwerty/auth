package db

import (
	"context"
	"time"
)

// User - структура, представляющая пользователя в базе данных.
// Содержит основные данные пользователя, такие как имя, электронная почта,
// хеш пароля, ID роли, а также время создания и обновления записи.
type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	RoleID       int32     // Идентификатор роли пользователя
	CreatedAt    time.Time // Время создания записи
	UpdatedAt    time.Time // Время последнего обновления записи
}

// Role - структура, представляющая роль пользователя в базе данных.
// Определяет права доступа и разрешения для пользователя.
type Role struct {
	ID       int64
	RoleName string // Название роли, описывающее её функционал
}

// DB - интерфейс для взаимодействия с базой данных.
// Описывает основные операции CRUD для структуры User.
type DB interface {
	// CreateUser - создает нового пользователя в базе данных и возвращает его ID.
	CreateUser(ctx context.Context, user User) (int64, error)

	// ReadUser - получает информацию о пользователе по его ID.
	// Возвращает структуру db.User и ошибку, если пользователь не найден.
	ReadUser(ctx context.Context, email string) (*User, error)

	// UpdateUser - обновляет информацию о пользователе в базе данных.
	UpdateUser(ctx context.Context, user User) error

	// DeleteUser - удаляет пользователя из базы данных по его ID.
	DeleteUser(ctx context.Context, id int64) error

	// CheckEmailExists - возвращает true если email есть в базе данных
	СheckEmailExists(ctx context.Context, email string) (bool, error)
}
