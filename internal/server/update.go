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

	location := time.FixedZone("UTC+3", 3*60*60)
	err := a.DB.UpdateUser(ctx, db.User{
		ID:        req.Id,
		Name:      req.Name.Value,
		RoleID:    int32(req.Role),
		UpdatedAt: time.Now().In(location),
	})

	return nil, err
}
