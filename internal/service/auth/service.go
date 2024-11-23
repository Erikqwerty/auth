package auth

import (
	"context"

	"github.com/erikqwerty/erik-platform/clients/db"
	"github.com/erikqwerty/erik-platform/clients/kafka"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/peer"

	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	dev "github.com/erikqwerty/auth/internal/service"
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
	userCache      repository.UserCache
	producer       kafka.Producer
	txManager      db.TxManager
}

// NewService - создает экземляр сервиса
func NewService(
	authRepository repository.AuthRepository,
	txManager db.TxManager,
	userCache repository.UserCache,
	producer kafka.Producer) dev.AuthService {

	return &service{
		authRepository: authRepository,
		userCache:      userCache,
		txManager:      txManager,
		producer:       producer,
	}
}

// prepareUserForCreate - добавляем в model.User хеш пароля
func prepareUserForCreate(user *model.CreateUser) error {
	passHash, err := hashPassword(user.PasswordHash)
	if err != nil {
		return err
	}

	user.PasswordHash = passHash

	return nil
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
	if peer != nil {
		info += peer.Addr.String()
	} else {
		info = "детальная информация отсутствует"
	}

	return info
}
