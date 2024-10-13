package server

import (
	"context"
	"time"

	"github.com/erikqwerty/auth/internal/db"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update обновление информации о пользователе по его идентификатору
func (a *Auth) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	localTime, _ := time.LoadLocation("Europe/Moscow")
	err := a.DB.UpdateUser(ctx, db.User{
		ID:        req.Id,
		Name:      req.Name.Value,
		RoleID:    int(req.Role),
		UpdatedAt: time.Now().In(localTime),
	})

	return nil, err
}
