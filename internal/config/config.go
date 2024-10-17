package config

import "github.com/joho/godotenv"

// Config - Содержит в себе конфигурацию приложения
type Config struct {
	GRPC GRPCConfig
	DB   PGConfig
}

// New - Создает новый обьект конфигурации
func New(path string) (*Config, error) {

	if err := load(path); err != nil {
		return nil, err
	}

	grpcConf, err := NewGRPCConfig()
	if err != nil {
		return nil, err
	}

	pgConf, err := NewPGConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		GRPC: grpcConf,
		DB:   pgConf,
	}, nil
}

// Load - Парсит файл и загружает переменные среды по указному пути
func load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
