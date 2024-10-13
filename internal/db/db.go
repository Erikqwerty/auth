package db

import (
	"context"
	"time"
)

// User - структура, представляющая пользователя в базе данных.
// Содержит основные данные пользователя, такие как имя, электронная почта,
// хеш пароля, ID роли, а также время создания и обновления записи.
type User struct {
	ID           int64     // Уникальный идентификатор пользователя
	Name         string    // Имя пользователя
	Email        string    // Электронная почта пользователя
	PasswordHash string    // Хеш пароля пользователя
	RoleID       int       // Идентификатор роли пользователя
	CreatedAt    time.Time // Время создания записи
	UpdatedAt    time.Time // Время последнего обновления записи
}

// Role - структура, представляющая роль пользователя в базе данных.
// Определяет права доступа и разрешения для пользователя.
type Role struct {
	ID       int64  // Уникальный идентификатор роли
	RoleName string // Название роли, описывающее её функционал
}

// DB - интерфейс для взаимодействия с базой данных.
// Описывает основные операции CRUD для структуры User.
type DB interface {
	CreateUser(ctx context.Context, user User) (int64, error) // Создает нового пользователя и возвращает его ID
	GetUser(ctx context.Context, id int64) (*User, error)     // Возвращает пользователя по его ID
	UpdateUser(ctx context.Context, user User) error          // Обновляет данные пользователя
	DeleteUser(ctx context.Context, id int64) error           // Удаляет пользователя по его ID
}
