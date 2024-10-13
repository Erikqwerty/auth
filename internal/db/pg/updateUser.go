package pg

import (
	"context"
	"fmt"

	"github.com/erikqwerty/auth/internal/db"
)

// UpdateUser - обновляет информацию о пользователе в базе данных.
func (pg *PG) UpdateUser(_ context.Context, user db.User) error {
	fmt.Println(user)
	return nil
}
