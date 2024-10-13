package main

import (
	"flag"
	"log"

	"github.com/erikqwerty/auth/internal/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	auth, err := server.NewAuthApp(configPath)
	if err != nil {
		log.Fatalf("Ошибка инициализации приложения %v", err)
	}

	srv := server.NewServer(auth)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
