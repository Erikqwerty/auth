package auth

import (
	"context"
	"errors"

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
	query := sq.
		Insert(tableUsers).
		Columns(
			nameColumn, emailColumn, passwordHashColumn,
			roleIDColumn, createdAtColumn, updatedAtColumn).
		Values(
			user.Name, user.Email, user.PasswordHash,
			user.RoleID, user.CreatedAt, user.UpdatedAt).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

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
