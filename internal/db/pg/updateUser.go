package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/auth/internal/db"
)

// UpdateUser обновляет информацию о пользователе в базе данных по его ID.
func (pg *PG) UpdateUser(ctx context.Context, user db.User) error {
	query := pg.sb.
		Update("users").
		Set("name", user.Name).
		Set("role_id", user.RoleID).
		Set("updated_at", user.UpdatedAt).
		Where(squirrel.Eq{"id": user.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		return errSQLtoSring(err)
	}

	_, err = pg.pool.Exec(ctx, sql, args...)
	if err != nil {
		return errSQLQwery(err)
	}

	return nil
}
