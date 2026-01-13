package handler

import "github.com/gofiber/fiber/v2"

type PaymentHandler struct {
	*BaseHandler
}

func NewPaymentHandler(baseURL string) *PaymentHandler {
	return &PaymentHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *PaymentHandler) ProcessPayment(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *PaymentHandler) FindPaymentByID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *PaymentHandler) HandleStripeWebhook(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
