package service

import (
	"context"
	"log"
	"time"

	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/client"
	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
	proto "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto/log"
	"github.com/bedirhangull/hrcubo/auth-service/internal/core/port"
	"github.com/bedirhangull/hrcubo/auth-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	userCacheDuration = 1 * time.Hour
	userCachePrefix   = "user:"
)

type UserService struct {
	repo      port.AuthRepository
	cache     port.CacheRepository
	logClient *client.Client
}

func NewUserService(repo port.AuthRepository, cache port.CacheRepository, logClient *client.Client) *UserService {
	return &UserService{
		repo:      repo,
		cache:     cache,
		logClient: logClient,
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

	cacheKey := util.GenerateCacheKey(userCachePrefix, user.Email)
	userData, err := util.Serialize(user)
	if err != nil {
		log.Printf("Failed to serialize user data: %v", err)
	} else {
		if err := s.cache.Set(ctx, cacheKey, userData, userCacheDuration); err != nil {
			log.Printf("Failed to cache user data: %v", err)
		}
	}

	logRes, err := s.logClient.CreateLog(ctx, "User registered: "+req.Email, proto.LogLevel_INFO)
	if err != nil {
		log.Printf("Failed to log registration: %v", err)
		return user, nil
	}
	if !logRes.Success {
		log.Printf("Log creation unsuccessful")
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email and password are required")
	}

	cacheKey := util.GenerateCacheKey(userCachePrefix, req.Email)
	cachedData, err := s.cache.Get(ctx, cacheKey)
	if err == nil && cachedData != nil {
		var user pb.LoginResponse
		if err := util.Deserialize(cachedData, &user); err == nil {
			log.Printf("User data retrieved from cache")
			return &user, nil
		}
	}

	user, err := s.repo.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	userData, err := util.Serialize(user)
	if err == nil {
		if err := s.cache.Set(ctx, cacheKey, userData, userCacheDuration); err != nil {
			log.Printf("Failed to cache user data: %v", err)
		}
	}

	logRes, err := s.logClient.CreateLog(ctx, "User login: "+req.Email, proto.LogLevel_INFO)
	if err != nil {
		log.Printf("Failed to log login: %v", err)
		return user, nil
	}
	if !logRes.Success {
		log.Printf("Log creation unsuccessful")
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

	cacheKey := util.GenerateCacheKey(userCachePrefix, *req.Email)
	userData, err := util.Serialize(user)
	if err == nil {
		if err := s.cache.Set(ctx, cacheKey, userData, userCacheDuration); err != nil {
			log.Printf("Failed to update user cache: %v", err)
		}
	}

	logRes, err := s.logClient.CreateLog(ctx, "User updated: "+req.Id, proto.LogLevel_INFO)
	if err != nil {
		log.Printf("Failed to log update: %v", err)
		return user, nil
	}
	if !logRes.Success {
		log.Printf("Log creation unsuccessful")
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := s.repo.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}

	cacheKey := util.GenerateCacheKey(userCachePrefix, req.Id)
	if err := s.cache.Delete(ctx, cacheKey); err != nil {
		log.Printf("Failed to delete user from cache: %v", err)
	}

	logRes, err := s.logClient.CreateLog(ctx, "User deleted: "+req.Id, proto.LogLevel_INFO)
	if err != nil {
		log.Printf("Failed to log deletion: %v", err)
		return &pb.DeleteUserResponse{Success: false}, nil
	}
	if !logRes.Success {
		log.Printf("Log creation unsuccessful")
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}
