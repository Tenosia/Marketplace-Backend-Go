package main

import (
	"context"
	"log"

	"github.com/marketplace-go-backend/services/4-user/handler"
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

	// Initialize gRPC clients
	ccs := handler.NewGRPCClients()

	// Start HTTP server in goroutine
	go NewHttpServer(db)

	// Start gRPC server
	grpcServer := NewGRPCServer(env.RequireEnv("USER_GRPC_PORT"))
	go func() {
		if err := grpcServer.Run(db); err != nil {
			log.Printf("gRPC server error: %v", err)
			cancel()
		}
	}()

	// Setup graceful shutdown
	shutdown.GracefulShutdown(cancel, func() error {
		log.Println("Shutting down user service...")
		_ = ctx // Use context if needed for cleanup
		return nil
	})
}
