package auth

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/erikqwerty/erik-platform/clients/db"

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
func (pg *repo) CreateUser(ctx context.Context, user *model.CreateUser) (int64, error) {
	query := sq.
		Insert(tableUsers).
		Columns(
			nameColumn, emailColumn, passwordHashColumn, roleIDColumn, createdAtColumn).
		Values(
			user.Name, user.Email, user.PasswordHash, user.RoleID, sq.Expr("NOW()")).
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
func (pg *repo) ReadUser(ctx context.Context, email string) (*model.UserInfo, error) {
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

	return convertor.ToUserInfoFromRepo(user), nil
}

// UpdateUser - обновляет информацию о пользователе.
func (pg *repo) UpdateUser(ctx context.Context, user *model.UpdateUser) error {
	if user.Email == nil {
		return errors.New("отсутствует email пользователя данные которого нужно обновить")
	}

	query := sq.Update(tableUsers)

	fieldsToUpdate := false

	if user.Name != nil {
		query = query.Set(nameColumn, user.Name)
		fieldsToUpdate = true
	}

	if user.RoleID != nil {
		query = query.Set(roleIDColumn, user.RoleID)
		fieldsToUpdate = true
	}

	if !fieldsToUpdate {
		return errors.New("отсутствует информация для обновления")
	}

	query = query.Set(updatedAtColumn, sq.Expr("NOW()"))

	query = query.Where(sq.Eq{emailColumn: user.Email}).PlaceholderFormat(sq.Dollar)

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
