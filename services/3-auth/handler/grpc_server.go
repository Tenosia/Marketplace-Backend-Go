package handler

import (
	"context"
	"log"

	"github.com/marketplace-go-backend/services/3-auth/service"
	pb "github.com/marketplace-go-backend/services/common/genproto/auth"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthGrpcHandler struct {
	authSvc service.AuthServiceImpl
	pb.UnimplementedAuthServiceServer
}

func NewAuthGRPCHandler(grpc *grpc.Server, authSvc service.AuthServiceImpl) {
	gRPCHandler := &AuthGrpcHandler{
		authSvc: authSvc,
	}

	pb.RegisterAuthServiceServer(grpc, gRPCHandler)
}

func (h *AuthGrpcHandler) FindUserByUserID(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	log.Println("FindUserByUserID receive data", req)

	u, err := h.authSvc.FindUserByIDIncPassword(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.FindUserResponse{
		Id:                     u.ID,
		Username:               u.Username,
		Email:                  u.Email,
		Password:               u.Password,
		Country:                u.Country,
		ProfilePicture:         u.ProfilePicture,
		ProfilePublicID:        u.ProfilePublicID,
		EmailVerified:          u.EmailVerified,
		EmailVerificationToken: u.EmailVerificationToken.String,
		PasswordResetToken:     u.PasswordResetToken.String,
		PasswordResetExpires:   timestamppb.New(*u.PasswordResetExpires),
		CreatedAt:              timestamppb.New(u.CreatedAt),
	}, nil
}
