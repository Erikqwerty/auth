package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/service"
	serviceMock "github.com/erikqwerty/auth/internal/service/mocks"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		pass  = gofakeit.Password(true, true, true, true, true, 8)

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        pass,
			PasswordConfirm: pass,
			Role:            desc.Role_ROLE_USER,
		}

		createUser = &model.CreateUser{
			Name:         name,
			Email:        email,
			PasswordHash: pass,
			RoleID:       int32(desc.Role_ROLE_USER),
		}

		res = &desc.CreateResponse{Id: id}

		tempErr = errors.New("service create error")
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.CreateResponse
		err                 error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.CreateUserMock.Expect(ctx, createUser).Return(id, nil)
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
				mock.CreateUserMock.Expect(ctx, createUser).Return(0, tempErr)
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

			newID, err := api.CreateUser(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, newID)
		})
	}

}
