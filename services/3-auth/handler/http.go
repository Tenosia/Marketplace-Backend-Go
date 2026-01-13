package handler

import (
	"github.com/marketplace-go-backend/services/3-auth/service"
	"github.com/marketplace-go-backend/services/3-auth/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHttpHandler struct {
	service *service.AuthService
	cld     *util.Cloudinary
	ccs     *GRPCClients
}

func NewAuthHttpHandler(svc *service.AuthService, cld *util.Cloudinary, ccs *GRPCClients) *AuthHttpHandler {
	return &AuthHttpHandler{
		service: svc,
		cld:     cld,
		ccs:     ccs,
	}
}

func (h *AuthHttpHandler) SignIn(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) SignUp(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) SendForgotPasswordURL(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) ResetPassword(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) GetUserInfo(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) RefreshToken(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) SendVerifyEmailURL(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) VerifyEmail(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHttpHandler) ChangePassword(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
