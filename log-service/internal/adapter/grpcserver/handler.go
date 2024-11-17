package grpcserver

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/log-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/log-service/internal/core/port"
)

type LogHandler struct {
	pb.UnimplementedLogServiceServer
	logService port.LogService
}

func NewLogHandler(logService port.LogService) pb.LogServiceServer {
	return &LogHandler{
		logService: logService,
	}
}

func (h *LogHandler) CreateLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	resp, err := h.logService.CreateLog(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *LogHandler) GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	resp, err := h.logService.GetLog(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *LogHandler) DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogResponse, error) {
	resp, err := h.logService.DeleteLog(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *LogHandler) ListLog(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogResponse, error) {
	resp, err := h.logService.ListLogs(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
