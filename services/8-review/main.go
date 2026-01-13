package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := NewStore()

	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	go NewHttpServer(db)

	grpcServer := NewGRPCServer(os.Getenv("REVIEW_GRPC_PORT"))
	err = grpcServer.Run(db)
	if err != nil {
		log.Fatal("Error listen GRPC")
	}
}
