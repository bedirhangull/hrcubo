package repository

import (
	"context"
	"time"

	pb "github.com/bedirhangull/hrcubo/log-service/internal/adapter/proto"
	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/storage/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogRepostiory struct {
	db *mongo.DB
}

func NewRepository(db *mongo.DB) *LogRepostiory {
	return &LogRepostiory{db: db}
}

func (l *LogRepostiory) CreateLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	id := uuid.New().String()
	now := time.Now().UTC().String()

	coll := l.db.Client.Database("logs").Collection("log")

	doc := pb.Log{
		Id:        id,
		Message:   req.Message,
		Level:     req.Level,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return &pb.LogResponse{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		Message:   req.Message,
		Level:     req.Level,
		CreatedAt: now,
		UpdatedAt: now,
		Success:   true,
	}, nil
}

func (l *LogRepostiory) GetLogById(ctx context.Context, req *pb.GetLogRequest) (*pb.GetLogResponse, error) {
	coll := l.db.Client.Database("logs").Collection("log")
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	var log pb.Log
	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&log)
	if err != nil {
		return nil, err
	}

	return &pb.GetLogResponse{
		Id:        log.Id,
		Message:   log.Message,
		Level:     log.Level,
		CreatedAt: log.CreatedAt,
		UpdatedAt: log.UpdatedAt,
		Success:   true,
	}, nil
}

func (l *LogRepostiory) DeleteLog(ctx context.Context, req *pb.DeleteLogRequest) (*pb.DeleteLogResponse, error) {
	coll := l.db.Client.Database("logs").Collection("log")
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteLogResponse{
		Id:      req.Id,
		Success: true,
	}, nil
}

func (l *LogRepostiory) ListAllLogs(ctx context.Context, req *pb.ListLogRequest) (*pb.ListLogResponse, error) {
	coll := l.db.Client.Database("logs").Collection("log")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var logs []*pb.Log
	for cursor.Next(ctx) {
		var log pb.Log
		err := cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}

	return &pb.ListLogResponse{
		Logs:    logs,
		Success: true,
	}, nil
}
