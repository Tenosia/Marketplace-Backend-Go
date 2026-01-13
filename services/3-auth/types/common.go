package types

import "github.com/golang-jwt/jwt/v5"

const (
	USER_SERVICE         = "user"
	NOTIFICATION_SERVICE = "notification"
)

type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
