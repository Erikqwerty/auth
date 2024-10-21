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

// Log - структура для логирования действий в БД
type Log struct {
	ID              int64
	ActionType      string
	ActionDetails   string
	ActionTimestamp time.Time
}
