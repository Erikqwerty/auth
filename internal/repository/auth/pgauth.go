package authrepository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	"github.com/erikqwerty/auth/internal/repository/auth/convertor"
	"github.com/jackc/pgx/v4/pgxpool"

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
	db *pgxpool.Pool
}

// NewRepo - Создает новый обьект repo, для работы с базой данных
func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

// CreateUser - создает нового пользователя (user)
func (pg *repo) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	query := sq.Insert(tableUsers).Columns(
		nameColumn, emailColumn, passwordHashColumn,
		roleIDColumn, createdAtColumn, updatedAtColumn).
		Values(
			user.Name, user.Email, user.PasswordHash,
			user.RoleID, user.CreatedAt, user.UpdatedAt).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	querySQL, args, err := query.ToSql() // Переименовал переменную sql
	if err != nil {
		return 0, fmt.Errorf("failed to generate SQL query: %w", err)
	}

	var id int64
	err = pg.db.QueryRow(ctx, querySQL, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return id, nil
}

// ReadUser - считывает информацию о пользователе из бд
func (pg *repo) ReadUser(ctx context.Context, email string) (*model.User, error) {
	query := sq.
		Select(idColumn, nameColumn, emailColumn, passwordHashColumn, roleIDColumn, createdAtColumn, updatedAtColumn).
		From(tableUsers).
		Where(sq.Eq{emailColumn: email}).Limit(1).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user = &modelRepo.User{}

	err = pg.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return convertor.ToAuthFromRepo(user), nil
}

// UpdateUser - обновляет информацию о пользователе.
func (pg *repo) UpdateUser(ctx context.Context, user *model.User) error {
	query := sq.
		Update(tableUsers).
		Set(nameColumn, user.Name).
		Set(roleIDColumn, user.RoleID).
		Set(updatedAtColumn, user.UpdatedAt).
		Where(sq.Eq{idColumn: user.ID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = pg.db.Exec(ctx, sql, args...)
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

	_, err = pg.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
