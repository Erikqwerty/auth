package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	repoMock "github.com/erikqwerty/auth/internal/repository/mocks"
	"github.com/erikqwerty/auth/internal/service/auth"
	"github.com/erikqwerty/auth/pkg/db"
	dbMock "github.com/erikqwerty/auth/pkg/db/mocks"
	"github.com/erikqwerty/auth/pkg/utils"
)

func TestGetUser(t *testing.T) {
	t.Parallel()

	type authRepoMockFunc func(mc *minimock.Controller) repository.AuthRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		req     = gofakeit.Email()
		repoErr = errors.New("repo error")
		time    = utils.TimeNowUTC3()

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		passhash = gofakeit.Password(true, true, true, true, true, 4)

		expectedRes = &model.UserInfo{
			ID: id,
			CreateUser: model.CreateUser{
				Name:         name,
				Email:        req,
				PasswordHash: passhash,
			},
			CreatedAt: time,
			UpdatedAt: &time,
		}
	)

	tests := []struct {
		name              string
		args              args
		want              *model.UserInfo
		err               error
		dbMockFunc        txManagerMockFunc
		authRepoMockFunc  authRepoMockFunc
		userCacheMockFunc userCacheMockFunc
	}{
		{
			name: "service get user info success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: expectedRes,
			err:  nil,
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				mock := repoMock.NewAuthRepositoryMock(t)
				mock.ReadUserMock.Expect(ctx, req).Return(expectedRes, nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    "GET",
					ActionDetails: "детальная информация отсутствует",
				}).Return(nil)
				return mock
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
		},
		{
			name: "service get user info error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				mock := repoMock.NewAuthRepositoryMock(t)
				mock.ReadUserMock.Expect(ctx, req).Return(nil, repoErr)
				return mock
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			authRepoMock := tt.authRepoMockFunc(mc)
			txManagerMock := tt.dbMockFunc(mc)
			cacheMock := tt.userCacheMockFunc(mc)

			servic := auth.NewService(authRepoMock, txManagerMock, cacheMock)

			user, err := servic.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, user)
		})
	}
}
