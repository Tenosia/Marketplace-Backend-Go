package main

import (
	"context"
	"log"
	"os"

	"github.com/marketplace-go-backend/services/3-auth/handler"
	"github.com/marketplace-go-backend/services/3-auth/types"
	"github.com/marketplace-go-backend/services/3-auth/util"
	"github.com/marketplace-go-backend/services/common/env"
	"github.com/marketplace-go-backend/services/common/shutdown"
)

func main() {
	// Load environment variables
	env.Load()

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database
	db, err := NewStore()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create PostgreSQL extensions
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Printf("Warning: Failed to create uuid-ossp extension: %v", err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pg_trgm";`).Error; err != nil {
		log.Printf("Warning: Failed to create pg_trgm extension: %v", err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`).Error; err != nil {
		log.Printf("Warning: Failed to create pgcrypto extension: %v", err)
	}

	// Initialize Cloudinary
	cld := util.NewCloudinary()

	// Initialize gRPC clients
	ccs := handler.NewGRPCClients()
	if err := ccs.AddClient(types.USER_SERVICE, env.RequireEnv("USER_GRPC_PORT")); err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	if err := ccs.AddClient(types.NOTIFICATION_SERVICE, env.RequireEnv("NOTIFICATION_GRPC_PORT")); err != nil {
		log.Fatalf("Failed to connect to notification service: %v", err)
	}

	// Start HTTP server in goroutine
	go NewHttpServer(db, cld, ccs)

	// Start gRPC server
	grpcServer := NewGRPCServer(env.RequireEnv("AUTH_GRPC_PORT"))
	go func() {
		if err := grpcServer.Run(db); err != nil {
			log.Printf("gRPC server error: %v", err)
			cancel()
		}
	}()

	// Setup graceful shutdown
	shutdown.GracefulShutdown(cancel, func() error {
		log.Println("Shutting down auth service...")
		// Add any cleanup logic here
		return nil
	})
}
