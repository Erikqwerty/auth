package api

import (
	"github.com/erikqwerty/auth/internal/service"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Implementation - имплементирует gRPC методы
type Implementation struct {
	desc.UnimplementedUserAPIV1Server
	authService service.AuthService
}

// NewImplementation - Создает новый обьект имплементирующий gRPC сервер
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
