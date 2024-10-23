package auth

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	"github.com/erikqwerty/auth/internal/repository/auth/convertor"

	modelRepo "github.com/erikqwerty/auth/internal/repository/auth/model"
)

var _ repository.AuthRepository = (*repo)(nil)

const (
	tableUsers = "users"

	idColumn           = "id"
	nameColumn         = "name"
	emailColumn        = "email"
	passwordHashColumn = "password_hash"
	roleIDColumn       = "role_id"
	createdAtColumn    = "created_at"
	updatedAtColumn    = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepo - Создает новый обьект repo, для работы с базой данных
func NewRepo(db db.Client) *repo {
	return &repo{db: db}
}

// CreateUser - создает нового пользователя (user)
func (pg *repo) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	query := sq.Insert(tableUsers)

	columns, values, err := prepareUserFields(user)
	if err != nil {
		return 0, err
	}

	query = query.
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	if len(columns) == 0 {
		return 0, fmt.Errorf("no valid fields to insert")
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository_CreateUser",
		QueryRaw: sql,
	}

	var id int64

	err = pg.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// ReadUser - считывает информацию о пользователе из бд
func (pg *repo) ReadUser(ctx context.Context, email string) (*model.User, error) {
	query := sq.
		Select(idColumn, nameColumn, emailColumn, passwordHashColumn, roleIDColumn, createdAtColumn, updatedAtColumn).
		From(tableUsers).
		Where(sq.Eq{emailColumn: email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository_ReadUser",
		QueryRaw: sql,
	}

	var user = &modelRepo.User{}

	err = pg.db.DB().ScanOneContext(ctx, user, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToAuthFromRepo(user), nil
}

// UpdateUser - обновляет информацию о пользователе.
func (pg *repo) UpdateUser(ctx context.Context, user *model.User) error {
	query := sq.Update(tableUsers)

	fieldsToUpdate := true

	if user.Name != "" {
		query = query.Set(nameColumn, user.Name)
		fieldsToUpdate = false
	}

	if user.RoleID != 0 {
		query = query.Set(roleIDColumn, user.RoleID)
		fieldsToUpdate = false
	}

	if !user.UpdatedAt.IsZero() {
		query = query.Set(updatedAtColumn, user.UpdatedAt)
		fieldsToUpdate = false
	}

	if fieldsToUpdate {
		return errors.New("ErrNothingToUpdate")
	}

	query.Where(sq.Eq{idColumn: user.ID}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository_UpdateUser",
		QueryRaw: sql,
	}

	_, err = pg.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser - удаляет пользователя из базы данных по его (id).
func (pg *repo) DeleteUser(ctx context.Context, id int64) error {
	query := sq.
		Delete(tableUsers).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "auth_repository_DeleteUser",
		QueryRaw: sql,
	}

	_, err = pg.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// prepareUserFields подготавливает колонки и значения для вставки нового пользователя
func prepareUserFields(user *model.User) ([]string, []interface{}, error) {
	columns := []string{}
	values := []interface{}{}

	if user.Name != "" {
		columns = append(columns, nameColumn)
		values = append(values, user.Name)
	}

	if user.Email != "" {
		columns = append(columns, emailColumn)
		values = append(values, user.Email)
	}

	if user.PasswordHash != "" {
		columns = append(columns, passwordHashColumn)
		values = append(values, user.PasswordHash)
	}

	if user.RoleID != 0 {
		columns = append(columns, roleIDColumn)
		values = append(values, user.RoleID)
	}

	if !user.CreatedAt.IsZero() {
		columns = append(columns, createdAtColumn)
		values = append(values, user.CreatedAt)
	}

	if !user.UpdatedAt.IsZero() {
		columns = append(columns, updatedAtColumn)
		values = append(values, user.UpdatedAt)
	}

	if len(columns) == 0 {
		return nil, nil, fmt.Errorf("no valid fields to insert")
	}

	return columns, values, nil
}
