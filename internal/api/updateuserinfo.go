package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// UpdateUserInfo - обрабатывает получаемый запрос от клиента gRPC, на обновление информации о пользователе
func (i *ImplServAuthUser) UpdateUserInfo(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.authService.UpdateUser(ctx, convertor.ToModelUserFromUpdateRequest(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
