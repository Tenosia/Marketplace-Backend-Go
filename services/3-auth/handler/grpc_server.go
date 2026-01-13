package handler

import (
	"context"

	"github.com/marketplace-go-backend/services/3-auth/service"
	pb "github.com/marketplace-go-backend/services/common/genproto/auth"
	"google.golang.org/grpc"
)

type AuthGRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthGRPCHandler(grpcServer *grpc.Server, svc *service.AuthService) {
	handler := &AuthGRPCHandler{
		service: svc,
	}
	pb.RegisterAuthServiceServer(grpcServer, handler)
}

func (h *AuthGRPCHandler) FindUserByUserID(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	// TODO: Implement
	return nil, nil
}
