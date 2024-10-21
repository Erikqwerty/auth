package authservice

import (
	"context"
	"time"

	"github.com/erikqwerty/auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Create - создание пользователя
func (s *service) Create(ctx context.Context, user *model.User) (int64, error) {

	passHash, err := s.hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = passHash

	location := time.FixedZone("UTC+3", 3*60*60)
	user.CreatedAt = time.Now().In(location)
	user.UpdatedAt = time.Now().In(location)
	var id int64
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.authRepository.CreateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.authRepository.ReadUser(ctx, user.Email)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// hashPassword - создает хеш из пароля
func (s *service) hashPassword(password string) (string, error) {
	// Используем bcrypt для хеширования пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
