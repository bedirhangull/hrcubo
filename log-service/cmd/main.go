package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/config"
	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/grpcserver"
	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/storage/mongo"
	"github.com/bedirhangull/hrcubo/log-service/internal/adapter/storage/mongo/repository"
	"github.com/bedirhangull/hrcubo/log-service/internal/core/service"
	"github.com/bedirhangull/hrcubo/log-service/pkg/logger"
)

func main() {
	newConfig, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := mongo.New(newConfig.DB, ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(ctx)

	dbLog := logger.NewLogItem("INFO", "DB Connection is successful")
	logger.Log(dbLog)

	// Test the database connection
	logRepo := repository.NewRepository(db)
	logService := service.NewLogService(logRepo)
	grpcServer := grpcserver.NewServer(logService)

	errChan := make(chan error, 1)
	go func() {
		err := grpcServer.Start(newConfig.Server)
		if err != nil {
			errChan <- fmt.Errorf("Failed to start gRPC server: %v", err)
		}
	}()

	select {
	case err := <-errChan:
		log.Fatalf("gRPC server error: %v", err)
	case <-time.After(2 * time.Second):
		log.Printf("gRPC server started successfully on port %s", newConfig.Server.Port)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.Stop()
	log.Println("Server stopped")
}
