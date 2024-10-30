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

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	type authRepoMockFunc func(mc *minimock.Controller) repository.AuthRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx  context.Context
		user *model.UpdateUser
	}

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		name             = gofakeit.Name()
		email            = gofakeit.Email()
		updateTime       = utils.TimeNowUTC3()
		roleID     int32 = 1
		repoErr          = errors.New("repo error")

		user = &model.UpdateUser{
			Email:     &email,
			Name:      &name,
			RoleID:    &roleID,
			UpdatedAt: &updateTime,
		}
	)

	tests := []struct {
		name             string
		args             args
		err              error
		dbMockFunc       txManagerMockFunc
		authRepoMockFunc authRepoMockFunc
	}{{
		name: name,
		args: args{
			ctx:  ctx,
			user: user,
		},
		err: nil,
		dbMockFunc: func(mc *minimock.Controller) db.TxManager {
			mock := dbMock.NewTxManagerMock(mc)
			mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
				return handler(ctx)
			})
			return mock
		},
		authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
			mock := repoMock.NewAuthRepositoryMock(t)
			mock.UpdateUserMock.Expect(ctx, user).Return(nil)
			mock.CreateLogMock.Expect(ctx, &model.Log{
				ActionType:    "UPDATE",
				ActionDetails: "детальная информация отсутствует",
			}).Return(nil)
			return mock
		}},
		{
			name: name,
			args: args{
				ctx:  ctx,
				user: user,
			},
			err: repoErr,
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				mock := repoMock.NewAuthRepositoryMock(t)
				mock.UpdateUserMock.Expect(ctx, user).Return(repoErr)
				return mock
			}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			authRepoMock := tt.authRepoMockFunc(mc)
			txManagerMock := tt.dbMockFunc(mc)

			servic := auth.NewService(authRepoMock, txManagerMock)
			err := servic.UpdateUser(ctx, user)
			require.Equal(t, tt.err, err)
		})
	}
}
