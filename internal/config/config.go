package config

import (
	"time"

	"github.com/joho/godotenv"
)

// GRPCConfig - интерфейс, представляющий конфигурацию для GRPC-сервера.
// Интерфейс определяет метод Address, который возвращает адрес GRPC-сервера
// в формате "host:port".
type GRPCConfig interface {
	Address() string
}

// PGConfig - интерфейс для конфигурации подключения к PostgreSQL.
type PGConfig interface {
	DSN() string // DSN возвращает строку подключения к PostgreSQL
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

// Load - Парсит файл и загружает переменные среды по указному пути
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
