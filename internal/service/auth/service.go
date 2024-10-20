package authservice

import (
	"github.com/erikqwerty/auth/internal/repository"
	dev "github.com/erikqwerty/auth/internal/service"
)

var _ dev.AuthService = (*service)(nil)

type service struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) dev.AuthService {
	return &service{
		authRepository: authRepository,
	}
}
