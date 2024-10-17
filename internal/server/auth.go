package server

import (
	"fmt"

	"github.com/erikqwerty/auth/internal/config"
	"github.com/erikqwerty/auth/internal/db"
	"github.com/erikqwerty/auth/internal/db/pg"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Auth - используется для реализации методов UserAPIV1.
type Auth struct {
	desc.UnimplementedUserAPIV1Server
	Config *config.Config
	DB     db.DB
}

// NewAuthApp - Создает структуру приложения аутентификации, загружая конфигурации
func NewAuthApp(path string) (*Auth, error) {
	a := &Auth{}
	conf, err := config.New(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации %v", err)
	}
	a.Config = conf

	database, err := pg.New(conf.DB.DSN())
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения dsn для подключения к базе данных %v", err)
	}
	a.DB = database

	return a, nil
}
