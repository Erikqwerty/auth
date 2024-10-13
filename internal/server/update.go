package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update обновление информации о пользователе по его идентификатору
func (a *Auth) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Обновление информации о пользователе по его идентификатору %v", req)
	return nil, nil
}
