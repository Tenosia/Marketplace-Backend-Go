package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles graceful shutdown for services
func GracefulShutdown(cancel context.CancelFunc, cleanup func() error) {
	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Block until a signal is received
	sig := <-quit
	log.Printf("Received signal: %v. Starting graceful shutdown...", sig)

	// Cancel context to stop all operations
	cancel()

	// Give services time to finish current operations
	shutdownTimeout := 30 * time.Second
	if timeout := os.Getenv("SHUTDOWN_TIMEOUT"); timeout != "" {
		if parsed, err := time.ParseDuration(timeout); err == nil {
			shutdownTimeout = parsed
		}
	}

	// Create a context with timeout for cleanup
	ctx, cancelTimeout := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelTimeout()

	// Run cleanup in a goroutine
	done := make(chan error, 1)
	go func() {
		if cleanup != nil {
			done <- cleanup()
		} else {
			done <- nil
		}
	}()

	// Wait for cleanup or timeout
	select {
	case err := <-done:
		if err != nil {
			log.Printf("Cleanup error: %v", err)
		} else {
			log.Println("Graceful shutdown completed successfully")
		}
	case <-ctx.Done():
		log.Printf("Shutdown timeout exceeded (%v). Forcing exit.", shutdownTimeout)
	}

	os.Exit(0)
}
