package api

import (
	"github.com/erikqwerty/auth/internal/service"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// ImplServAuthUser - имплементирует gRPC методы
type ImplServAuthUser struct {
	desc.UnimplementedUserAPIV1Server
	authService service.AuthService
}

// NewImplementationServAuthUser - Создает новый обьект имплементирующий gRPC сервер
func NewImplementationServAuthUser(authService service.AuthService) *ImplServAuthUser {
	return &ImplServAuthUser{authService: authService}
}
