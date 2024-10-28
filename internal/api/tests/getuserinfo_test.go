package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/service"
	serviceMock "github.com/erikqwerty/auth/internal/service/mocks"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"github.com/erikqwerty/auth/pkg/utils"
)

func TestGetUserInfo(t *testing.T) {
	t.Parallel()
	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email   = gofakeit.Email()
		tempErr = errors.New("service getuserinfo error")
		time    = utils.TimeNowUTC3()
		id      = gofakeit.Int64()
		name    = gofakeit.Name()

		req = &desc.GetRequest{
			Email: email,
		}

		userInfo = &model.UserInfo{
			ID: id,
			CreateUser: model.CreateUser{
				Name:      name,
				Email:     email,
				RoleID:    int32(desc.Role_ROLE_USER),
				CreatedAt: time,
			},
			UpdatedAt: time,
		}

		resp = &desc.GetResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      desc.Role_ROLE_USER,
			CreatedAt: timestamppb.New(time),
			UpdatedAt: timestamppb.New(time),
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name                string
		args                args
		want                *desc.GetResponse
		err                 error
		authServiceMockFunc authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			authServiceMockFunc: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.GetUserMock.Expect(ctx, email).Return(userInfo, nil)
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
				mock.GetUserMock.Expect(ctx, email).Return(nil, tempErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			authServiceMock := tt.authServiceMockFunc(mc)
			api := api.NewImplementationServAuthUser(authServiceMock)

			got, err := api.GetUserInfo(tt.args.ctx, tt.args.req)

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
