package server

import (
	"context"
	"log"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаление пользователя из системы по его идентификатору
func (a *Auth) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Удаление пользователя из системы по его идентификатору: %v", req.Id)
	return nil, nil
}
