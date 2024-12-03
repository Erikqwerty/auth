package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// DeleteUser - обрабатывает получаемый запрос от клиента gRPC на удаление пользователя
func (i *ImplServAuthUser) DeleteUser(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
