package main

import (
	"log"

	"github.com/erikqwerty/auth/internal/server"
)

const grpcPort = 50052

func main() {
	auth := &server.Auth{}
	srv := server.NewServer(grpcPort, auth)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
