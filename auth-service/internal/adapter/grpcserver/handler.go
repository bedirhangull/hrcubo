package grpcserver

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/auth-service/internal/core/port"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) pb.AuthServiceServer {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp, err := h.authService.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := h.authService.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *AuthHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	resp, err := h.authService.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *AuthHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	resp, err := h.authService.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
