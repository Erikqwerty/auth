package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/service"
	serviceMock "github.com/erikqwerty/auth/internal/service/mocks"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

func TestUpdateUserInfo(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		email   = gofakeit.Email()
		name    = gofakeit.Name()
		role    = desc.Role_ROLE_ADMIN
		tempErr = errors.New("service updateUser error")

		req = &desc.UpdateRequest{
			Email: email,
			Name:  wrapperspb.String(name),
			Role:  role,
		}

		rol = int32(role)

		updateUser = &model.UpdateUser{
			Email:  &email,
			Name:   &name,
			RoleID: &rol,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name                string
		args                args
		want                *emptypb.Empty
		err                 error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.UpdateUserMock.Expect(ctx, updateUser).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  tempErr,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.UpdateUserMock.Expect(ctx, updateUser).Return(tempErr)
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

			got, err := api.UpdateUserInfo(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
