package model

import "time"

// User - структура, представляющая пользователя в базе данных.
type User struct {
	ID           int64     `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	RoleID       int32     `db:"role_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Role - структура, представляющая роль пользователя в базе данных.
type Role struct {
	ID       int64  `db:"id"`
	RoleName string `db:"role_name"`
}
