package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewStore creates a new GORM database connection with optimized settings
func NewStore() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Optimized GORM configuration
	config := &gorm.Config{
		SkipDefaultTransaction: true, // Skip default transaction for better performance
		PrepareStmt:            true, // Use prepared statements for better performance
		Logger:                  logger.Default.LogMode(logger.Silent), // Disable logging in production
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Optimize connection pool settings
	sqlDB.SetMaxIdleConns(10)                    // Maximum idle connections
	sqlDB.SetMaxOpenConns(100)                   // Maximum open connections
	sqlDB.SetConnMaxLifetime(time.Hour)          // Maximum connection lifetime
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)   // Maximum idle time before closing

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return db, nil
}
