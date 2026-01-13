package handler

import (
	"context"

	"github.com/marketplace-go-backend/services/2-notification/service"
	pb "github.com/marketplace-go-backend/services/common/genproto/notification"
)

type GRPCServer struct {
	pb.UnimplementedNotificationServiceServer
	service *service.NotificationService
}

func NewGRPCServer(grpcServer interface{}, svc *service.NotificationService) {
	// TODO: Register the service
}

func (s *GRPCServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	// TODO: Implement
	return &pb.SendEmailResponse{
		Success: false,
		Message: "Not implemented",
	}, nil
}
