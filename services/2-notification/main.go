package main

import (
	"log"
	"os"

	"github.com/marketplace-go-backend/services/2-notification/handler"
	"github.com/marketplace-go-backend/services/2-notification/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	grpcServer := NewGRPCServer(os.Getenv("NOTIFICATION_GRPC_PORT"))
	err = grpcServer.Run()
	if err != nil {
		log.Fatal("Error listen GRPC")
	}
}
