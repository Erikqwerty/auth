package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

// DeleteUser - удаляет пользователя из базы данных по его (id).
func (pg *PG) DeleteUser(ctx context.Context, id int64) error {

	status, err := checkIDExists(ctx, pg, id)
	if !status {
		return fmt.Errorf("пользователя с таким ID не существует")
	}
	if err != nil {
		return err
	}

	query := pg.sb.
		Delete("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return errSQLCreateQwery(err)
	}

	_, err = pg.pool.Exec(ctx, sql, args...)
	if err != nil {
		return errSQLQwery(err)
	}

	return nil
}

func checkIDExists(ctx context.Context, pg *PG, id int64) (bool, error) {
	query := pg.sb.
		Select("1").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return false, errSQLCreateQwery(err)
	}

	var exists int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&exists)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, errSQLQwery(err)
	}

	return true, nil
}
