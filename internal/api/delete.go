package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Delete - обрабатывает получаемый запрос от клиента gRPC на удаление пользователя
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.authService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
