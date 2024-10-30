package env

import (
	"errors"
	"os"

	"github.com/erikqwerty/auth/internal/config"
)

const (
	// dsnEnvName - имя переменной окружения для строки подключения к PostgreSQL.
	dsnEnvName = "PG_DSN"
)

// pgConfig - структура, реализующая интерфейс PGConfig для хранения DSN.
type pgConfig struct {
	dsn string // Строка подключения к PostgreSQL
}

// NewPGConfig - создает новую конфигурацию для PostgreSQL, читая DSN из переменной окружения.
// Возвращает ошибку, если переменная окружения отсутствует.
func NewPGConfig() (config.PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

// DSN - метод для получения строки подключения к PostgreSQL.
func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
