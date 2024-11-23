package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/erikqwerty/erik-platform/clients/db"
	dbMock "github.com/erikqwerty/erik-platform/clients/db/mocks"
	"github.com/erikqwerty/erik-platform/clients/kafka"
	kafkaMock "github.com/erikqwerty/erik-platform/clients/kafka/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/autherrors"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	repoMock "github.com/erikqwerty/auth/internal/repository/mocks"
	"github.com/erikqwerty/auth/internal/service/auth"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	type authRepoMockFunc func(mc *minimock.Controller) repository.AuthRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) repository.UserCache
	type kafkaProducerMockFunc func(mc *minimock.Controller) kafka.Producer

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

		repoErr = errors.New("репо ошибка")
	)

	tests := []struct {
		name                  string
		args                  args
		want                  int64
		err                   error
		authRepoMockFunc      authRepoMockFunc
		dbMockFunc            txManagerMockFunc
		userCacheMockFunc     userCacheMockFunc
		kafkaProducerMockFunc kafkaProducerMockFunc
	}{{
		name: "service create user success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: id,
		err:  nil,
		dbMockFunc: func(_ *minimock.Controller) db.TxManager {
			mock := dbMock.NewTxManagerMock(t)

			mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
				return handler(ctx)
			})

			return mock
		},
		authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
			mock := repoMock.NewAuthRepositoryMock(t)

			mock.CreateUserMock.Expect(ctx, req).Return(id, nil)
			mock.CreateLogMock.Expect(ctx, &model.Log{
				ActionType:    "CREATE",
				ActionDetails: "детальная информация отсутствует",
			}).Return(nil)

			return mock
		},
		userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
			mock := repoMock.NewUserCacheMock(t)
			return mock
		},
		kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
			mock := kafkaMock.NewProducerMock(t)
			mock.SendMessageMock.Expect("CREATEUSER", req.Email).Return(0, 0, nil)
			return mock
		},
	},
		{
			name: "service create user error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
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

				mock.CreateUserMock.Expect(ctx, req).Return(0, repoErr)

				return mock
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
			kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
				return kafkaMock.NewProducerMock(t)
			},
		},
		{
			name: "service create user with invalid data",
			args: args{
				ctx: ctx,
				req: &model.CreateUser{
					Name:         "erik",
					Email:        "invalid email",
					PasswordHash: "123",
					RoleID:       0,
				},
			},
			want: 0,
			err:  autherrors.ErrInvalidEmail,
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				return repoMock.NewAuthRepositoryMock(t)
			},
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				return dbMock.NewTxManagerMock(mc)
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
			kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
				return kafkaMock.NewProducerMock(t)
			},
		},
		{
			name: "service transaction failure",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  errors.New("transaction failed"),
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(mc)

				mock.ReadCommittedMock.Set(func(_ context.Context, _ db.Handler) error {
					return errors.New("transaction failed")
				})

				return mock
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				return repoMock.NewAuthRepositoryMock(t)
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
			kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
				return kafkaMock.NewProducerMock(t)
			},
		},
		{
			name: "service log writing failure",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  errors.New("log writing failed"),
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(mc)

				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})

				return mock
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				mock := repoMock.NewAuthRepositoryMock(t)

				mock.CreateUserMock.Expect(ctx, req).Return(id, nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    "CREATE",
					ActionDetails: "детальная информация отсутствует",
				}).Return(errors.New("log writing failed"))

				return mock
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
			kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
				return kafkaMock.NewProducerMock(t)
			},
		},
		{
			name: "service structure CreateUser nil",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: 0,
			err:  autherrors.ErrCreateUserNil,
			dbMockFunc: func(mc *minimock.Controller) db.TxManager {
				return dbMock.NewTxManagerMock(mc)
			},
			authRepoMockFunc: func(_ *minimock.Controller) repository.AuthRepository {
				return repoMock.NewAuthRepositoryMock(t)
			},
			userCacheMockFunc: func(_ *minimock.Controller) repository.UserCache {
				mock := repoMock.NewUserCacheMock(t)
				return mock
			},
			kafkaProducerMockFunc: func(mc *minimock.Controller) kafka.Producer {
				return kafkaMock.NewProducerMock(t)
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
			kafkaProducerMock := tt.kafkaProducerMockFunc(mc)

			servic := auth.NewService(authRepoMock, txManagerMock, cacheMock, kafkaProducerMock)

			ID, err := servic.CreateUser(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, ID)
		})
	}

	t.Cleanup(mc.Finish)
}
