package model

import (
	"time"
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
	RoleID       int64
	CreatedAt    time.Time
}

// UserInfo - структра представляющая информацию о пользователе
type UserInfo struct {
	ID int64
	CreateUser
	UpdatedAt time.Time
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
