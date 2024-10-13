package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/auth/internal/db"
)

// GetUser - получает информацию о пользователе по его ID.
// Возвращает структуру db.User и ошибку, если пользователь не найден.
func (pg *PG) SelectUser(ctx context.Context, email string) (*db.User, error) {
	query := pg.sb.
		Select("id", "name", "email", "password_hash", "role_id", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"email": email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errSQLtoSring(err)
	}

	var user db.User

	err = pg.pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, errSQLQwery(err)
	}

	return &user, nil
}
