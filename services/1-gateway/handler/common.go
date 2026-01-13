package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BaseHandler struct {
	baseURL string
}

func NewBaseHandler(baseURL string) *BaseHandler {
	return &BaseHandler{
		baseURL: baseURL,
	}
}

func (h *BaseHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Service is healthy and OK")
}
