package config

import (
	"errors"
	"net"
	"os"
)

// Имена переменных окружения для GRPC-конфигурации
const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

// GRPCConfig - интерфейс, представляющий конфигурацию для GRPC-сервера.
// Интерфейс определяет метод Address, который возвращает адрес GRPC-сервера
// в формате "host:port".
type GRPCConfig interface {
	Address() string
}

// grpcConfig - структура, реализующая интерфейс GRPCConfig.
type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig - функция для создания новой конфигурации GRPC-сервера.
// Читает хост и порт из переменных окружения GRPC_HOST и GRPC_PORT.
// Возвращает ошибку, если одна из переменных не найдена.
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Address - метод, реализующий интерфейс GRPCConfig для структуры grpcConfig.
// Возвращает адрес в формате "host:port", соединяя хост и порт с помощью функции net.JoinHostPort.
func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
