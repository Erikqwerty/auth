package app

import (
	"context"
	"log"

	"github.com/erikqwerty/erik-platform/clients/db"
	"github.com/erikqwerty/erik-platform/clients/db/pg"
	"github.com/erikqwerty/erik-platform/clients/db/transaction"
	"github.com/erikqwerty/erik-platform/closer"
	"github.com/gomodule/redigo/redis"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/client/cache"
	redisCL "github.com/erikqwerty/auth/internal/client/cache/redis"
	"github.com/erikqwerty/auth/internal/config"
	"github.com/erikqwerty/auth/internal/config/env"
	"github.com/erikqwerty/auth/internal/repository"
	authrepository "github.com/erikqwerty/auth/internal/repository/auth"
	cacherepository "github.com/erikqwerty/auth/internal/repository/cache"
	"github.com/erikqwerty/auth/internal/service"
	authservice "github.com/erikqwerty/auth/internal/service/auth"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig
	httpConfig  config.HTTPConfig

	dbClient       db.Client
	txManager      db.TxManager
	authRepository repository.AuthRepository
	userCache      repository.UserCache

	redisPool   *redis.Pool
	redisClient cache.RedisClient

	authService service.AuthService

	authImpl *api.ImplServAuthUser
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig - инициализирует конфигурацию базы данных
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
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
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("ошибка загрузки конфигурации gRPC сервера: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("ошибка загрузки конфигурации redis: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
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

func (s *serviceProvider) RedisPool() *redis.Pool {
	if s.redisPool == nil {
		s.redisPool = &redis.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redis.Conn, error) {
				return redis.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) RedisClient() cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redisCL.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) UserCache() repository.UserCache {
	if s.userCache == nil {
		s.userCache = cacherepository.NewCache(s.RedisClient())
	}
	return s.userCache
}

// AuthService - инициализирует сервисный слой сервиса auth
func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authservice.NewService(s.AuthRepository(ctx), s.TxManager(ctx), s.UserCache())
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
