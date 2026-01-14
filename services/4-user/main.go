package main

import (
	"log"
	"os"

	"github.com/marketplace-go-backend/services/4-user/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := NewStore()

	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "pg_trgm";`)

	ccs := handler.NewGRPCClients()

	go NewHttpServer(db)

	grpcServer := NewGRPCServer(os.Getenv("USER_GRPC_PORT"))
	err = grpcServer.Run(db)
	if err != nil {
		log.Fatal("Error listen GRPC")
	}
}
