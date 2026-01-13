package handler

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	clients map[string]*grpc.ClientConn
}

func NewGRPCClients() *GRPCClients {
	return &GRPCClients{
		clients: make(map[string]*grpc.ClientConn),
	}
}

func (ccs *GRPCClients) AddClient(serviceName, addr string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s service: %v", serviceName, err)
	}
	ccs.clients[serviceName] = conn
	log.Printf("Connected to %s service at %s", serviceName, addr)
}

func (ccs *GRPCClients) GetClient(serviceName string) *grpc.ClientConn {
	return ccs.clients[serviceName]
}
