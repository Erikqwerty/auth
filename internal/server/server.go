package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Server представляет gRPC сервер
type Server struct {
	grpcPort int
	auth     *Auth
}

// NewServer создает новый экземпляр Server
func NewServer(grpcPort int, auth *Auth) *Server {
	return &Server{
		grpcPort: grpcPort,
		auth:     auth,
	}
}

// Start запускает gRPC сервер
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterUserAPIV1Server(grpcServer, s.auth)

	log.Printf("server listening at :%v", s.grpcPort)

	return grpcServer.Serve(lis)
}
