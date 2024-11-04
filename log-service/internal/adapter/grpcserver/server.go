package grpcserver

import (
	"log"
	"net"

	"github.com/bedirhangull/hrcubo/log-service/internal/core/port"
	"google.golang.org/grpc"

	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/config"
	pb "github.com/bedirhangull/hrcubo/log-service/internal/adapter/proto"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	logService port.LogService
	grpcServer *grpc.Server
}

func NewServer(logService port.LogService) *Server {
	grpcServer := grpc.NewServer()
	server := &Server{
		logService: logService,
		grpcServer: grpcServer,
	}

	pb.RegisterLogServiceServer(grpcServer, NewLogHandler(logService))

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
