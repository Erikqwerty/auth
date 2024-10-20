package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/erikqwerty/auth/internal/api"
	"github.com/erikqwerty/auth/internal/config"
	authrepository "github.com/erikqwerty/auth/internal/repository/auth"
	authservice "github.com/erikqwerty/auth/internal/service/auth"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", conf.GRPC.Address())
	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.Connect(ctx, conf.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	authRepo := authrepository.NewRepo(pool)

	authServ := authservice.NewService(authRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, api.NewImplementation(authServ))
	log.Printf("server start and listen: %v", conf.GRPC.Address())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server faild start: %v", err)
	}
}
