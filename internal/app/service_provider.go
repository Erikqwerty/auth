package app

import (
	"context"
	"log"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/client/db/pg"
	"github.com/erikqwerty/auth/internal/client/db/transaction"
	"github.com/erikqwerty/auth/internal/closer"
	"github.com/erikqwerty/auth/internal/config"
	"github.com/erikqwerty/auth/internal/repository"
	"github.com/erikqwerty/auth/internal/service"

	authrepository "github.com/erikqwerty/auth/internal/repository/auth"
	authservice "github.com/erikqwerty/auth/internal/service/auth"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	authRepository repository.AuthRepository

	authService service.AuthService

	authImpl *api.ImplServAuthUser
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig - инициализирует конфигурацию базы данных
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("ошибка загрущки конфигурации базы данных: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig - инициализирует конфигурацию gRPC сервера
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("ошибка загрузки конфигурации gRPC сервера: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// DBClient - создает клиента для подключения к базе данных
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("ошибка подключения к базе данных: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping до базы данных не проходит: %v", err)
		}

		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager - инициализирует менеджер транзакций
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// AuthRepository инициализирует репозиторий auth для работы с бд
func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authrepository.NewRepo(s.DBClient(ctx))
	}

	return s.authRepository
}

// AuthService - инициализирует сервисный слой сервиса auth
func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authservice.NewService(s.AuthRepository(ctx), s.TxManager(ctx))
	}

	return s.authService
}

// AuthImpl - инициализирует иплиментацию gRPC сервера auth
func (s *serviceProvider) AuthImpl(ctx context.Context) *api.ImplServAuthUser {
	if s.authImpl == nil {
		s.authImpl = api.NewImplementationServAuthUser(s.AuthService(ctx))
	}

	return s.authImpl
}
