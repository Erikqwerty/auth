package cache

import (
	"context"
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"

	clredis "github.com/erikqwerty/auth/internal/client/cache"
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository"
	"github.com/erikqwerty/auth/internal/repository/cache/convertor"
	cachemodel "github.com/erikqwerty/auth/internal/repository/cache/model"
)

var _ repository.UserCache = (*cache)(nil)

type cache struct {
	redisClient clredis.RedisClient
}

// NewCache создает новый объект кеша
func NewCache(client clredis.RedisClient) repository.UserCache {
	return &cache{
		redisClient: client,
	}
}

// SetUser сохраняет пользователя в кеш с email в качестве ключа
func (c *cache) SetUser(ctx context.Context, email string, user *model.UserCache) error {
	if user == nil {
		return errors.New("пустой пользователь, нечего записывать в кеш")
	}

	// Конвертируем объект `user` в формат, подходящий для хранения в Redis
	u := convertor.ToUserCacheModelFromServiceUserCache(user)

	// Используем email как ключ
	err := c.redisClient.HashSet(ctx, email, u)
	if err != nil {
		return err
	}

	// Устанавливаем TTL (время жизни) для кеша
	err = c.redisClient.Expire(ctx, email, time.Hour)
	if err != nil {
		return err
	}

	return nil
}

// GetUser получает пользователя из кеша по email
func (c *cache) GetUser(ctx context.Context, email string) (*model.UserCache, error) {
	// Получаем все поля, сохранённые в хеше по ключу email
	values, err := c.redisClient.HGetAll(ctx, email)
	if err != nil {
		return nil, err
	}

	// Если значений нет, значит пользователь не найден в кеше
	if len(values) == 0 {
		return nil, errors.New("пользователь не найден в кеше")
	}

	// Создаем объект `UserCache` для декодирования из Redis
	var user cachemodel.UserCache
	err = redis.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	// Конвертируем объект в модель `model.UserCache` и возвращаем его
	return convertor.ToServiceUserCacheFromUserCacheModel(&user), nil
}
