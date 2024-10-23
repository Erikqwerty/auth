package api

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// UpdateUserInfo - обрабатывает получаемый запрос от клиента gRPC, на обновление информации о пользователе
func (i *ImplServAuthUser) UpdateUserInfo(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if req.Email == "" {
		return nil, errors.New("не указан пользователь данные которого нужно обновить")
	}

	updateScope := 0

	if req.Name.Value != "" {
		updateScope++
	}

	if req.Role.String() != "" {
		updateScope++
	}

	if updateScope == 0 {
		return nil, errors.New("не переданны данные для обновления")
	}

	err := i.authService.UpdateUser(ctx, convertor.ToUpdateUserFromUpdateRequest(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
