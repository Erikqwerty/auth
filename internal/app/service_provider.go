package app

import (
	"context"
	"log"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/client/db"
	"github.com/erikqwerty/auth/internal/client/db/pg"
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
	authRepository repository.AuthRepository

	authService service.AuthService

	authImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

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

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authrepository.NewRepo(s.DBClient(ctx))
	}
	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authservice.NewService(s.AuthRepository(ctx))
	}
	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *api.Implementation {
	if s.authImpl == nil {
		s.authImpl = api.NewImplementation(s.AuthService(ctx))
	}
	return s.authImpl
}
