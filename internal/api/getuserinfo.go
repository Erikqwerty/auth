package api

import (
	"context"
	"errors"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// GetUserInfo - обрабатывает получаемый запрос от клиента gRPC, на получение информации о пользователе
func (i *ImplServAuthUser) GetUserInfo(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req.Email == "" {
		return nil, errors.New("не указан email")
	}

	user, err := i.authService.GetUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return convertor.ToGetResponseFromModelUser(user), nil
}
