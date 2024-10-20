package model

import (
	"time"
)

// User - структура, представляющая пользователя в базе данных.
type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	PasswordHash string
	RoleID       int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Role - структура, представляющая роль пользователя в базе данных.
type Role struct {
	ID       int64
	RoleName string
}
