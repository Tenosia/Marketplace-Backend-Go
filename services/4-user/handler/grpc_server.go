package handler

import (
	"context"

	"github.com/marketplace-go-backend/services/4-user/service"
	pb "github.com/marketplace-go-backend/services/common/genproto/user"
	"google.golang.org/grpc"
)

type UserGRPCHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserGRPCHandler(grpcServer *grpc.Server, svc *service.UserService) {
	handler := &UserGRPCHandler{
		service: svc,
	}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *UserGRPCHandler) SaveBuyerData(ctx context.Context, req *pb.SaveBuyerRequest) (*pb.SaveBuyerResponse, error) {
	// TODO: Implement
	return nil, nil
}

func (h *UserGRPCHandler) FindSeller(ctx context.Context, req *pb.FindSellerRequest) (*pb.FindSellerResponse, error) {
	// TODO: Implement
	return nil, nil
}

func (h *UserGRPCHandler) UpdateSellerBalance(ctx context.Context, req *pb.UpdateSellerBalanceRequest) (*pb.UpdateSellerBalanceResponse, error) {
	// TODO: Implement
	return nil, nil
}

func (h *UserGRPCHandler) FindBuyer(ctx context.Context, req *pb.FindBuyerRequest) (*pb.FindBuyerResponse, error) {
	// TODO: Implement
	return nil, nil
}
