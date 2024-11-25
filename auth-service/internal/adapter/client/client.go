package client

import (
	"context"
	"time"

	pb "github.com/bedirhangull/hrcubo/auth-service/internal/adapter/proto/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.LogServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{client: pb.NewLogServiceClient(conn)}
}

func (c *Client) CreateLog(ctx context.Context, message string, level pb.LogLevel) (*pb.LogResponse, error) {
	now := time.Now().Format(time.RFC3339)

	return c.client.CreateLog(ctx, &pb.LogRequest{
		Id:        uuid.New().String(),
		Message:   message,
		Level:     level,
		CreatedAt: now,
		UpdatedAt: now,
	})
}
