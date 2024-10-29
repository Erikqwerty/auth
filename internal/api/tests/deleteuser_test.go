package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/service"
	serviceMock "github.com/erikqwerty/auth/internal/service/mocks"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

func TestDeleteUser(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		req = &desc.DeleteRequest{
			Id: id,
		}

		tempErr = errors.New("service delete error")
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name                string
		args                args
		want                error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "api delete user success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.DeleteUserMock.Expect(ctx, req.Id).Return(nil)
				return mock
			},
		},
		{
			name: "api delete user error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: tempErr,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.DeleteUserMock.Expect(ctx, req.Id).Return(tempErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			authServiceMock := tt.authServiceMockFunc(mc)
			api := api.NewImplementationServAuthUser(authServiceMock)

			_, err := api.DeleteUser(tt.args.ctx, tt.args.req)

			if tt.want != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.want.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
