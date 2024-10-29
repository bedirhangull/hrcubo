package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/storage/postgres"
	"github.com/bedirhangull/hrcubo/auth-service/pkg/util"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{db: db}
}

func formatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}

func (u *UserRepository) CreateUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id := uuid.New().String()
	now := time.Now().UTC()

	query := u.db.QueryBuilder.Insert("users").
		Columns(
			"id", "email", "password", "first_name", "last_name",
			"phone", "image_id", "is_user_premium", "role",
			"created_at", "updated_at",
		).
		Values(
			id, req.Email, req.Password, req.Profile.FirstName,
			req.Profile.LastName, req.Profile.Phone, req.Profile.ImageId,
			req.Profile.IsUserPremium, req.Role.String(), now, now,
		).
		Suffix("RETURNING id, email, first_name, last_name, phone, image_id, is_user_premium, role, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	var response pb.RegisterResponse
	var profile pb.UserProfile
	var roleStr string
	var createdAt, updatedAt time.Time

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&response.Id,
		&response.Email,
		&profile.FirstName,
		&profile.LastName,
		&profile.Phone,
		&profile.ImageId,
		&profile.IsUserPremium,
		&roleStr,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if errCode := u.db.ErrorCode(err); errCode == "23505" {
			return nil, grpc.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response.Profile = &profile
	response.Role = pb.UserRole(pb.UserRole_value[roleStr])
	response.CreatedAt = formatTimestamp(createdAt)
	response.UpdatedAt = formatTimestamp(updatedAt)
	response.Success = true

	return &response, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	query := u.db.QueryBuilder.
		Select("id, email, password, first_name, last_name, phone, image_id, is_user_premium, role, created_at, updated_at").
		From("users").
		Where(sq.Eq{"email": req.Email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	var response pb.LoginResponse
	var profile pb.UserProfile
	var roleStr string
	var createdAt, updatedAt time.Time
	var hashedPassword string

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&response.Id,
		&response.Email,
		&hashedPassword,
		&profile.FirstName,
		&profile.LastName,
		&profile.Phone,
		&profile.ImageId,
		&profile.IsUserPremium,
		&roleStr,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	err = util.ComparePassword(req.Password, hashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid email or password")
	}

	response.Profile = &profile
	response.Role = pb.UserRole(pb.UserRole_value[roleStr])
	response.CreatedAt = formatTimestamp(createdAt)
	response.UpdatedAt = formatTimestamp(updatedAt)
	response.Success = true

	return &response, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	query := u.db.QueryBuilder.Update("users")

	if req.Email != nil {
		query = query.Set("email", req.Email)
	}
	if req.Password != nil {
		query = query.Set("password", req.Password)
	}
	if req.Profile != nil {
		if req.Profile.FirstName != "" {
			query = query.Set("first_name", req.Profile.FirstName)
		}
		if req.Profile.LastName != "" {
			query = query.Set("last_name", req.Profile.LastName)
		}
		if req.Profile.Phone != nil && *req.Profile.Phone != "" {
			query = query.Set("phone", req.Profile.Phone)
		}
		if req.Profile.ImageId != nil && *req.Profile.ImageId != "" {
			query = query.Set("image_id", req.Profile.ImageId)
		}
		query = query.Set("is_user_premium", req.Profile.IsUserPremium)
	}
	if req.Role != nil {
		query = query.Set("role", req.Role.String())
	}

	query = query.
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": req.Id}).
		Suffix("RETURNING id, email, first_name, last_name, phone, image_id, is_user_premium, role, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	var response pb.UpdateUserResponse
	var profile pb.UserProfile
	var roleStr string
	var createdAt, updatedAt time.Time
	var hashedPassword string

	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&response.Id,
		&response.Email,
		&hashedPassword,
		&profile.FirstName,
		&profile.LastName,
		&profile.Phone,
		&profile.ImageId,
		&profile.IsUserPremium,
		&roleStr,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if errCode := u.db.ErrorCode(err); errCode == "23505" {
			return nil, grpc.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	response.Profile = &profile
	response.Role = pb.UserRole(pb.UserRole_value[roleStr])
	response.CreatedAt = formatTimestamp(createdAt)
	response.UpdatedAt = formatTimestamp(updatedAt)
	response.Success = true

	return &response, nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	query := u.db.QueryBuilder.Delete("users").
		Where(sq.Eq{"id": req.Id}).
		Suffix("RETURNING id")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	var deletedUserId string
	err = u.db.QueryRow(ctx, sql, args...).Scan(&deletedUserId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}
