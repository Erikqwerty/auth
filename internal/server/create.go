package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/erikqwerty/auth/internal/db"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"golang.org/x/crypto/bcrypt"
)

// Create создание нового пользователя в системе
func (a *Auth) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, fmt.Errorf("пароль не совпадает с полем подтверждения пароля, создание пользователя не возможно")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	localTime, _ := time.LoadLocation("Europe/Moscow")
	user := db.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		RoleID:       int(req.Role),
		CreatedAt:    time.Now().In(localTime),
		UpdatedAt:    time.Now().In(localTime),
	}

	id, err := a.DB.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	log.Printf("Создание нового пользователя в системе: %v, %v, %v, %v, %v", req.Name, req.Email, req.Password, req.PasswordConfirm, req.Role)
	return &desc.CreateResponse{
		Id: id,
	}, nil
}
