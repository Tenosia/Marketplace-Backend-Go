package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load loads environment variables from .env file
// It's safe to call multiple times and won't fail if .env doesn't exist
func Load() {
	// Try to load .env file, but don't fail if it doesn't exist
	// This allows the app to run with environment variables set externally
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v (using system environment variables)", err)
	}
}

// GetEnv retrieves an environment variable or returns the default value
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// RequireEnv retrieves an environment variable and logs a fatal error if not set
func RequireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}
