package http

import (
	"github.com/marketplace-go-backend/services/4-user/service"
	"github.com/gofiber/fiber/v2"
)

type UserHttpHandler struct {
	service *service.UserService
	ccs     interface{} // GRPCClients placeholder
}

func NewUserHttpHandler(svc *service.UserService, ccs interface{}) *UserHttpHandler {
	return &UserHttpHandler{
		service: svc,
		ccs:     ccs,
	}
}

// TODO: Implement HTTP handlers
