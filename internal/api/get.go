package api

import (
	"context"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.authService.Get(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return convertor.ToGetResponseFromModelUser(user), nil
}
