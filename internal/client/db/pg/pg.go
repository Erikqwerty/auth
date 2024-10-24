package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/client/db/prettier"
)

type key string

const (
	TxKey key = "tx" // TxKey - ключ по которому достаются транзакции
)

// pg представляет реализацию интерфейса db.DB для работы с базой данных Postgres через pgxpool.
type pg struct {
	dbc *pgxpool.Pool
}

// NewDB создает новый объект для работы с базой данных Postgres через pgxpool.
// dbc - это пул соединений, который будет использоваться для выполнения запросов.
func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOneContext выполняет SQL-запрос и сканирует одну запись в dest.
// ctx - контекст для запроса, dest - куда будут помещены результаты, q - запрос, args - аргументы запроса.
func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAllContext выполняет SQL-запрос и сканирует все записи в dest.
// ctx - контекст для запроса, dest - куда будут помещены результаты, q - запрос, args - аргументы запроса.
func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, row)
}

// ExecContext выполняет SQL-запрос на выполнение (например, INSERT, UPDATE, DELETE).
func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext выполняет SQL-запрос и возвращает pgx.Rows с результатами.
func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext выполняет SQL-запрос и возвращает одну строку результата.
func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// MakeContextTx - Добавление транзакции к контексту
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

// Ping выполняет пинг к базе данных для проверки доступности.
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close закрывает пул соединений с базой данных.
func (p *pg) Close() {
	p.dbc.Close()
}

// logQuery логирует SQL-запрос с его именем и параметрами.
func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)

	log.Println(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}
