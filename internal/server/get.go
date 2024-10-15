package server

import (
	"context"
	"fmt"
	"log"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Get получение информации о пользователе по его идентификатору
func (a *Auth) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Получение информации о пользователе по его идентификатору: %v", req.Email)

	user, err := a.DB.ReadUser(ctx, req.Email)
	if err != nil {
		return &desc.GetResponse{}, fmt.Errorf("пользователь с таким email не существует %v", err)
	}

	return &desc.GetResponse{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      desc.Role(user.RoleID),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt)},
		nil
}
