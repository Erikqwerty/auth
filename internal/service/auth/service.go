package authservice

import (
	"github.com/erikqwerty/auth/internal/client/db"
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
	txManager      db.TxManager
}

// NewService - создает экземляр сервиса
func NewService(authRepository repository.AuthRepository, txManager db.TxManager) dev.AuthService {
	return &service{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
