package api

import (
	"github.com/erikqwerty/auth/internal/service"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

type Implementation struct {
	desc.UnimplementedUserAPIV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
