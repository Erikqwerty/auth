package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Auth используется для реализации методов UserAPIV1.
type Auth struct {
	desc.UnimplementedUserAPIV1Server
}

// Create создание нового пользователя в системе
func (a *Auth) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Создание нового пользователя в системе: %v, %v, %v, %v, %v", req.Name, req.Email, req.Password, req.PasswordConfirm, req.Role)
	return nil, nil
}

// Get получение информации о пользователе по его идентификатору
func (a *Auth) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Получение информации о пользователе по его идентификатору: %v", req.Id)
	return nil, nil
}

// Update обновление информации о пользователе по его идентификатору
func (a *Auth) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Обновление информации о пользователе по его идентификатору %v", req)
	return nil, nil
}

// Delete удаление пользователя из системы по его идентификатору
func (a *Auth) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Удаление пользователя из системы по его идентификатору: %v", req.Id)
	return nil, nil
}
