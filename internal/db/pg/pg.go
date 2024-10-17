package pg

import (
	"context"

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
func New(dsn string) (*PG, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errDSN(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errDBConect(err)
	}

	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &PG{pool: pool, sb: sb}, nil
}
