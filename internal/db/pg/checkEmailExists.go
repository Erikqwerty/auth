package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
)

// сheckEmailExists - возвращает true если email есть в базе данных
func (pg *PG) СheckEmailExists(ctx context.Context, email string) (bool, error) {
	query := pg.sb.
		Select("1").
		From("users").
		Where(squirrel.Eq{"email": email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return false, errSQLtoSring(err)
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
