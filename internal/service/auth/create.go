package authservice

import (
	"context"
	"time"

	"github.com/erikqwerty/auth/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/peer"
)

// Create - создание пользователя
func (s *service) Create(ctx context.Context, user *model.User) (int64, error) {

	if err := prepareUserForCreate(user); err != nil {
		return 0, err
	}

	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {

		var errTx error
		id, errTx = s.authRepository.CreateUser(ctx, user)
		if errTx != nil {
			return errTx
		}

		if errTx := s.createLog(ctx, actionTypeCreate); errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// prepareUserForCreate - добавляем в model.User хеш пароля и задаем время создания и обновления пользователя
func prepareUserForCreate(user *model.User) error {
	// Хэшируем пароль
	passHash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.PasswordHash = passHash

	// Устанавливаем временные метки
	t := timeNowUTC3()
	user.CreatedAt = t
	user.UpdatedAt = t

	return nil
}

// timeNowUTC3 + возвращает время +3
func timeNowUTC3() time.Time {
	return time.Now().In(time.FixedZone("UTC+3", 3*60*60))
}

// hashPassword - создает хеш из пароля
func hashPassword(password string) (string, error) {
	// Используем bcrypt для хеширования пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// details - информация о пользователе
func details(ctx context.Context) string {
	info := "Адрес:"
	peer, _ := peer.FromContext(ctx)
	info += peer.Addr.String()
	return info
}
