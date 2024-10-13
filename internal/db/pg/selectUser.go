package pg

import (
	"context"
	"fmt"

	"github.com/erikqwerty/auth/internal/db"
)

// GetUser - получает информацию о пользователе по его ID.
// Возвращает структуру db.User и ошибку, если пользователь не найден.
func (pg *PG) SelectUser(_ context.Context, email string) (*db.User, error) {
	fmt.Println(email)
	return &db.User{}, nil
}
