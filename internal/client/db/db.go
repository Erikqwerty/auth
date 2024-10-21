package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// DB описывает интерфейс для работы с базой данных.
// Включает операции выполнения SQL-запросов, управления транзакциями и проверки соединений.
type DB interface {
	SQLExecer
	Pinger
	Transactor
	Close()
}

// Client определяет интерфейс клиента для взаимодействия с базой данных.
type Client interface {
	DB() DB
	Close() error
}

// Handler представляет функцию, которая выполняется в рамках транзакции.
type Handler func(ctx context.Context) error

// TxManager определяет интерфейс менеджера транзакций, который обрабатывает транзакции
// с использованием пользовательских функций-обработчиков.
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Transactor интерфейс для работы с транзакциями
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Query обертка над запросом, хранящая имя запроса и сам запрос
// Имя запроса используется для логирования и потенциально может использоваться еще где-то, например, для трейсинга
type Query struct {
	Name     string
	QueryRaw string
}

// SQLExecer комбинирует NamedExecer и QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer интерфейс для работы с именованными запросами с помощью тегов в структурах
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer интерфейс для работы с обычными запросами
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}
