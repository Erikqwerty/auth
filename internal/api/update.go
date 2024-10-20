package api

import (
	"context"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update - обрабатывает получаемый запрос от клиента gRPC, на обновление информации о пользователе
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	err := i.authService.Update(ctx, convertor.ToModelUserFromUpdateRequest(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
