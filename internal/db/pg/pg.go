package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/auth/internal/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Подтверждение того, что PG реализует интерфейс db.DB
var _ db.DB = (*PG)(nil)

// PG - структура для работы с базой данных PostgreSQL через pgxpool и squirrel.
type PG struct {
	pool *pgxpool.Pool                 // Пул соединений с базой данных
	sb   squirrel.StatementBuilderType // Объект для построения SQL-запросов с помощью squirrel
}

// New - создает новый экземпляр PG для работы с базой данных PostgreSQL.
// Принимает DSN (Data Source Name) и возвращает ошибку, если подключение невозможно.
// New - создает новый объект для работы с базой данных PostgreSQL.
func New(dsn string) (*PG, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &PG{pool: pool, sb: sb}, nil
}

// GetUser - получает информацию о пользователе по его ID.
// Возвращает структуру db.User и ошибку, если пользователь не найден.
func (pg *PG) GetUser(_ context.Context, id int64) (*db.User, error) {
	fmt.Println(id)
	return &db.User{}, nil
}

// UpdateUser - обновляет информацию о пользователе в базе данных.
func (pg *PG) UpdateUser(_ context.Context, user db.User) error {
	fmt.Println(user)
	return nil
}

// DeleteUser - удаляет пользователя из базы данных по его ID.
func (pg *PG) DeleteUser(_ context.Context, id int64) error {
	fmt.Println(id)
	return nil
}
