package pg

import (
	"context"
	"fmt"

	"github.com/erikqwerty/auth/internal/db"
)

// GetUser - получает информацию о пользователе по его ID.
// Возвращает структуру db.User и ошибку, если пользователь не найден.
func (pg *PG) SelectUser(_ context.Context, id int64) (*db.User, error) {
	fmt.Println(id)
	return &db.User{}, nil
}
