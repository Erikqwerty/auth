package server

import (
	"context"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete удаление пользователя из системы по его идентификатору
func (a *Auth) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := a.DB.DeleteUser(ctx, req.Id)
	return nil, err
}
