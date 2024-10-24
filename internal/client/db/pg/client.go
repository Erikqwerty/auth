package pg

import (
	"context"

	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

// pgClient представляет клиента базы данных, который использует masterDBS для работы с базой данных.
type pgClient struct {
	masterDBS db.DB
}

// New создает новый клиент базы данных (pgClient).
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pgClient{
		masterDBS: &pg{dbc: dbc},
	}, nil
}

// DB возвращает объект базы данных, который использует pgClient для выполнения запросов.
func (c *pgClient) DB() db.DB {
	return c.masterDBS
}

// Close закрывает подключение к базе данных.
func (c *pgClient) Close() error {
	if c.masterDBS != nil {
		c.masterDBS.Close()
	}

	return nil
}
