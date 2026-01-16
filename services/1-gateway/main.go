package main

import (
	"context"
	"log"
	"time"

	"github.com/marketplace-go-backend/services/1-gateway/config"
	"github.com/marketplace-go-backend/services/common/env"
	"github.com/marketplace-go-backend/services/common/shutdown"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load environment variables
	env.Load()

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := env.GetEnv("PORT", ":8080")
	config.NewGoogleAuthConfig(port)

	// Create Fiber app with optimized configuration
	app := fiber.New(fiber.Config{
		BodyLimit:            5 * 1024 * 1024, // 5MB
		CaseSensitive:        true,
		StrictRouting:        true,
		DisableStartupMessage: false,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
		IdleTimeout:         120 * time.Second,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     env.GetEnv("CLIENT_URL", "*"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
	}))
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Setup routes
	MainRouter(app)

	// Start server in goroutine
	go func() {
		log.Printf("Gateway service starting on port %s", port)
		if err := app.Listen(port); err != nil {
			log.Printf("Gateway server error: %v", err)
			cancel()
		}
	}()

	// Setup graceful shutdown
	shutdown.GracefulShutdown(cancel, func() error {
		log.Println("Shutting down gateway service...")
		return app.Shutdown()
	})
}
