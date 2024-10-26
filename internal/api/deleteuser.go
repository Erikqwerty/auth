package api

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// DeleteUser - обрабатывает получаемый запрос от клиента gRPC на удаление пользователя
func (i *ImplServAuthUser) DeleteUser(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := validateDataDeleteUser(req); err != nil {
		return nil, err
	}

	err := i.authService.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// validateDataDeleteUser необходима для проверки переданных данных и их валидации перед обработкой в сервисном слое
func validateDataDeleteUser(req *desc.DeleteRequest) error {
	if req.Id == 0 {
		return errors.New("id удаляемого пользователя не может быть равен 0")
	}

	return nil
}
