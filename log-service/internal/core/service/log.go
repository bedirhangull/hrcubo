package service

import (
	"context"

	pb "github.com/bedirhangull/hrcubo/log-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/log-service/internal/core/port"
)

type LogService struct {
	repo port.LogRepository
}

func NewLogService(repo port.LogRepository) *LogService {
	return &LogService{
		repo: repo,
	}
}

func (s *LogService) CreateLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	log, err := s.repo.CreateLog(ctx, req)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *LogService) GetLog(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	log, err := s.repo.GetLogById(ctx, req)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *LogService) DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogResponse, error) {
	log, err := s.repo.DeleteLog(ctx, req)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (s *LogService) ListLogs(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogResponse, error) {
	logs, err := s.repo.ListAllLogs(ctx, req)
	if err != nil {
		return nil, err
	}

	return logs, nil
}
