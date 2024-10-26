package auth

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/model"
)

const (
	tableLogs = "user_log"

	actionType      = "action_type"
	actionDetails   = "action_details"
	actionTimestamp = "action_timestamp"
)

// CreateLog - записываем лог
func (pg *repo) CreateLog(ctx context.Context, log *model.Log) error {
	query := sq.Insert(tableLogs).
		Columns(actionType, actionDetails, actionTimestamp).
		Values(log.ActionType, log.ActionDetails, sq.Expr("NOW()")).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository_CreateLog",
		QueryRaw: sql,
	}

	_, err = pg.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
