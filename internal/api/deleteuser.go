package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// DeleteUser - обрабатывает получаемый запрос от клиента gRPC на удаление пользователя
func (i *ImplServAuthUser) DeleteUser(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := validateRequest(req); err != nil {
		return nil, err
	}

	err := i.authService.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
