package service

import (
	"context"
	"time"

	"github.com/marketplace-go-backend/services/3-auth/types"
	"github.com/marketplace-go-backend/services/common/genproto/user"
	"gorm.io/gorm"
)

type AuthServiceImpl interface {
	FindUserByUsernameOrEmail(ctx context.Context, str string) (*types.AuthExcludePassword, error)
	FindUserByUsernameOrEmailIncPassword(ctx context.Context, str string) (*types.Auth, error)
	Create(ctx context.Context, data *types.SignUp, userGrpcClient user.UserServiceClient) (*types.AuthExcludePassword, error)
	FindUserByID(ctx context.Context, id string) (*types.AuthExcludePassword, error)
	FindUserByIDIncPassword(ctx context.Context, id string) (*types.Auth, error)
	FindUserByVerificationToken(ctx context.Context, token string) (*types.AuthExcludePassword, error)
	FindUserByPasswordToken(ctx context.Context, token string) (*types.Auth, error)
	UpdateEmailVerification(ctx context.Context, userId string, emailStatus bool, emailVerifToken ...string) (*types.AuthExcludePassword, error)
	UpdatePasswordToken(ctx context.Context, userId string, token string, tokenExpiration time.Time) error
	UpdatePassword(ctx context.Context, userId string, password string) error
}

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthServiceImpl {
	return &AuthService{
		db: db,
	}
}

// TODO: Implement auth service methods
