package grpcserver

import (
	"log"
	"net"

	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/config"
	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/auth-service/internal/core/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	authService port.AuthService
	grpcServer  *grpc.Server
}

func NewServer(authService port.AuthService) *Server {
	grpcServer := grpc.NewServer()
	server := &Server{
		grpcServer:  grpcServer,
		authService: authService,
	}

	pb.RegisterAuthServiceServer(grpcServer, NewAuthHandler(authService))

	reflection.Register(grpcServer)

	return server
}

func (s *Server) Start(config *config.Server) error {
	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		panic(err)
	}

	log.Printf("gRPC server listening on port %s", config.Port)
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
