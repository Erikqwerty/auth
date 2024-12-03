package api

import (
	"context"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// CreateUser - обрабатывает получаемый запрос от клиента gRPC на создание пользователя
func (i *ImplServAuthUser) CreateUser(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.authService.CreateUser(ctx, convertor.ToCreateUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, err
}
