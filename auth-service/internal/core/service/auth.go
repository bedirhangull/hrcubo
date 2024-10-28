package service

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/auth-service/internal/core/port"
	"github.com/bedirhangull/hrcubo/auth-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	repo port.AuthRepository
}

func NewUserService(repo port.AuthRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hashedPassword

	user, err := s.repo.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email and password are required")
	}
	user, err := s.repo.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	existingUser, err := s.repo.GetUserByEmail(ctx, &pb.LoginRequest{Email: *req.Email})
	if err != nil {
		return nil, err
	}

	var hashedPassword string
	if req.Password != nil {
		hashedPassword, err = util.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
	}

	if hashedPassword != "" {
		req.Password = &hashedPassword
	}

	if req.Profile == nil {
		req.Profile = existingUser.Profile
	}

	if req.Role == nil {
		req.Role = &existingUser.Role
	}

	user, err := s.repo.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := s.repo.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}
