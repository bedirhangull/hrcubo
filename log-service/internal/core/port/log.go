package port

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/log-service/internal/adapter/proto"
)

type LogRepository interface {
	CreateLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error)
	GetLogById(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error)
	DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogResponse, error)
	ListAllLogs(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogResponse, error)
}

type LogService interface {
	CreateLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error)
	GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error)
	DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogResponse, error)
	ListLogs(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogResponse, error)
}
