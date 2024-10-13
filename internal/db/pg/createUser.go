package pg

import (
	"context"
	"fmt"

	"github.com/erikqwerty/auth/internal/db"
)

// CreateUser - создает нового пользователя в базе данных и возвращает его ID.
func (pg *PG) CreateUser(ctx context.Context, user db.User) (int64, error) {
	query := pg.sb.Insert("users").Columns("name", "email", "password_hash", "role_id", "created_at", "updated_at").
		Values(user.Name, user.Email, user.PasswordHash, user.RoleID, user.CreatedAt, user.UpdatedAt).
		Suffix("RETURNING id")

	// Конвертируем построенный запрос в SQL-строку и список аргументов.
	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build CreateUser query: %w", err)
	}

	var id int64

	// Выполняем SQL-запрос и сохраняем возвращенный ID в переменную id.
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute CreateUser query: %w", err)
	}

	return id, nil
}
