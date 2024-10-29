package port

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	GetUserByEmail(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}

type AuthService interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}
