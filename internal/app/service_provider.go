package app

import (
	"context"
	"log"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/closer"
	"github.com/erikqwerty/auth/internal/config"
	"github.com/erikqwerty/auth/internal/repository"
	authrepository "github.com/erikqwerty/auth/internal/repository/auth"
	"github.com/erikqwerty/auth/internal/service"
	authservice "github.com/erikqwerty/auth/internal/service/auth"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
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
			log.Fatalf("ошибка загрущки конфигурации gRPC сервера: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {

	if s.pgConfig == nil {
		s.PGConfig()
	}

	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.pgConfig.DSN())
		if err != nil {
			log.Fatalf("ошибка подключения к базе данных: %v", err)
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping до базы данных не проходит: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}
	return s.pgPool
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authrepository.NewRepo(s.PgPool(ctx))
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
