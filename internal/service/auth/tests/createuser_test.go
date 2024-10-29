package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/client/db"
	dbMock "github.com/erikqwerty/auth/internal/client/db/mocks"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	repoMock "github.com/erikqwerty/auth/internal/repository/mocks"
	"github.com/erikqwerty/auth/internal/service/auth"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type authRepoMockFunc func(mc *minimock.Controller) repository.AuthRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.CreateUser
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		name           = gofakeit.Name()
		email          = gofakeit.Email()
		passhash       = gofakeit.Password(true, true, true, true, true, 4)
		roleID   int32 = 1

		req = &model.CreateUser{
			Name:         name,
			Email:        email,
			PasswordHash: passhash,
			RoleID:       roleID,
		}

		// repoErr = errors.New("репо ошибка")
	)

	tests := []struct {
		name             string
		args             args
		want             int64
		err              error
		authRepoMockFunc authRepoMockFunc
		dbMockFunc       txManagerMockFunc
	}{{
		name: "service create user success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: id,
		err:  nil,
		dbMockFunc: func(mc *minimock.Controller) db.TxManager {
			mock := dbMock.NewTxManagerMock(mc)
			mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
				return handler(ctx)
			})
			return mock
		},
		authRepoMockFunc: func(mc *minimock.Controller) repository.AuthRepository {
			mock := repoMock.NewAuthRepositoryMock(t)
			mock.CreateUserMock.Expect(ctx, req).Return(id, nil)
			mock.CreateLogMock.Expect(ctx, &model.Log{
				ActionType:    "CREATE",
				ActionDetails: "детальная информация отсутствует",
			}).Return(nil)
			return mock
		},
	},
		{
			name: "service create user error case",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authRepoMock := tt.authRepoMockFunc(mc)
			txManagerMock := tt.dbMockFunc(mc)

			servic := auth.NewService(authRepoMock, txManagerMock)

			ID, err := servic.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, ID)
		})
	}
}
