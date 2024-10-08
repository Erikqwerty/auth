package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/erikqwerty/auth/pkg/UserAPI_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50052

type auth struct {
	desc.UnimplementedUserAPIV1Server
}

func (a *auth) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Создание нового пользователя в системе: %v, %v, %v, %v, %v", req.Name, req.Email, req.Password, req.PasswordConfirm, req.Role)
	return nil, nil
}

func (a *auth) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Получение информации о пользователе по его идентификатору: %v", req.Id)
	return nil, nil
}
func (a *auth) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Обновление информации о пользователе по его идентификатору %v", req)
	return nil, nil
}

func (a *auth) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Eдаление пользователя из системы по его идентификатору: %v", req.Id)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, &auth{})

	log.Printf("server listening at :%v", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Faider to server: ", err)
	}
}
