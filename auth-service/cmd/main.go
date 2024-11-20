package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/client"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/config"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/grpcserver"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/storage/postgres"
	"github.com/bedirhangull/hrcubo/auth-service/internal/adapter/storage/postgres/repository"
	"github.com/bedirhangull/hrcubo/auth-service/internal/core/service"
	"github.com/bedirhangull/hrcubo/auth-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Config initialization
	newConfig, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// DB connection and migration
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := postgres.New(ctx, newConfig.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	dbLog := logger.NewLogItem("INFO", "DB Connection is successful")
	logger.Log(dbLog)

	// Test the database connection
	var result int
	err = db.QueryRow(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		dbError := logger.NewLogItem("ERROR", fmt.Sprintf("Failed to execute test query: %v", err))
		logger.Log(dbError)
	}

	err = db.Migrate()
	if err != nil {
		migrationError := logger.NewLogItem("ERROR", fmt.Sprintf("Migration error: %v", err))
		logger.Log(migrationError)
	} else {
		migrationSuccess := logger.NewLogItem("INFO", "Migrations completed successfully")
		logger.Log(migrationSuccess)
	}

	logger.Log(logger.NewLogItem("INFO", "Migration is successful"))

	// service url from env
	sl := config.NewServiceList()

	// Log connection
	logConn, err := grpc.NewClient(sl.GetServiceURL("log"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to log service: %v", err)
	}
	defer logConn.Close()

	// Dependency injection
	logClient := client.NewClient(logConn)
	userRepo := repository.NewRepository(db)
	userService := service.NewUserService(userRepo, logClient)
	grpcServer := grpcserver.NewServer(userService)

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
