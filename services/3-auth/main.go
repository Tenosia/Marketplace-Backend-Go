package main

import (
	"log"
	"os"

	"github.com/marketplace-go-backend/services/3-auth/handler"
	"github.com/marketplace-go-backend/services/3-auth/types"
	"github.com/marketplace-go-backend/services/3-auth/util"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := NewStore()
	cld := util.NewCloudinary()

	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "pg_trgm";`)
	db.Debug().Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	ccs := handler.NewGRPCClients()
	ccs.AddClient(types.USER_SERVICE, os.Getenv("USER_GRPC_PORT"))
	ccs.AddClient(types.NOTIFICATION_SERVICE, os.Getenv("NOTIFICATION_GRPC_PORT"))

	go NewHttpServer(db, cld, ccs)

	grpcServer := NewGRPCServer(os.Getenv("AUTH_GRPC_PORT"))
	err = grpcServer.Run(db)
	if err != nil {
		log.Fatal("Error listen GRPC")
	}
}
