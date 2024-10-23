package auth

import (
	"context"
	"time"

	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	dev "github.com/erikqwerty/auth/internal/service"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/peer"
)

var _ dev.AuthService = (*service)(nil)

const (
	actionTypeCreate = "CREATE"
	actionTypeGet    = "GET"
	actionTypeUpdate = "UPDATE"
	actionTypeDelete = "DELETE"
)

type service struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

// NewService - создает экземляр сервиса
func NewService(authRepository repository.AuthRepository, txManager db.TxManager) dev.AuthService {
	return &service{
		authRepository: authRepository,
		txManager:      txManager,
	}
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
