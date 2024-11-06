package model

import (
	"time"

	"github.com/erikqwerty/auth/internal/autherrors"
	"github.com/erikqwerty/auth/pkg/utils"
)

// User - структура, представляющая пользователя в базе данных.
type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	RoleID   int32
}

// CreateUser - структра представляющая создание пользователя
type CreateUser struct {
	Name         string
	Email        string
	PasswordHash string
	RoleID       int32
}

// Validate - используйте для проверки СreateUser на ожидаемые поля
func (u *CreateUser) Validate() error {
	switch {
	case u.Name == "":
		return autherrors.ErrNameNotSpecified
	case u.Email == "":
		return autherrors.ErrEmailNotSpecified
	case !utils.IsValidEmail(u.Email):
		return autherrors.ErrInvalidEmail
	case u.PasswordHash == "":
		return autherrors.ErrPasswordNotSpecified
	case u.RoleID == 0:
		return autherrors.ErrRoleNotSpecified
	case u.RoleID > 2:
		return autherrors.ErrInvalidRole
	}

	return nil
}

// UserInfo - структра представляющая информацию о пользователе
type UserInfo struct {
	ID int64
	CreateUser
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// UpdateUser -  структра представляющая обновление пользователя
type UpdateUser struct {
	Email     *string // не обновляется, используется как условие (кого обновлять)
	Name      *string
	RoleID    *int32
	UpdatedAt *time.Time
}

// Role - структура, представляющая роль пользователя в базе данных.
type Role struct {
	ID       int32
	RoleName string
}

// Log - структура для логирования действий в БД
type Log struct {
	ActionType      string
	ActionDetails   string
	ActionTimestamp time.Time
}

// UserCache - структура для кеширования пользователя в redis
// type UserCache struct {
// 	ID           string
// 	Email        string
// 	Name         string
// 	PasswordHash string
// 	RoleID       int32
// 	CreatedAt    time.Time
// 	UpdatedAt    *time.Time
// }
