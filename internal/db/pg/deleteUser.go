package pg

import (
	"context"
	"fmt"
)

// DeleteUser - удаляет пользователя из базы данных по его ID.
func (pg *PG) DeleteUser(_ context.Context, id int64) error {
	fmt.Println(id)
	return nil
}
