package handler

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	*BaseHandler
}

func NewAuthHandler(baseURL string) *AuthHandler {
	return &AuthHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *AuthHandler) AuthWithGoogle(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SignUpWithGoogle(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SignInWithGoogle(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SendForgotPasswordURL(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) SendVerifyEmailURL(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) VerifyEmail(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
