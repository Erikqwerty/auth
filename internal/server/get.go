package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Get получение информации о пользователе по его идентификатору
func (a *Auth) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Получение информации о пользователе по его идентификатору: %v", req.Email)
	return nil, nil
}
