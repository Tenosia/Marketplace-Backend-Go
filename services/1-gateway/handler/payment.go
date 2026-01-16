package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	base_url string
}

func NewPaymentHandler(base_url string) *PaymentHandler {
	return &PaymentHandler{
		base_url: base_url,
	}
}

func (ph *PaymentHandler) HealthCheck(c *fiber.Ctx) error {
	route := ph.base_url + "/health-check"
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	if err != nil {
		log.Printf("PAYMENT - health check error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (ph *PaymentHandler) ProcessPayment(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/process")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PAYMENT", "process payment")
}

func (ph *PaymentHandler) FindPaymentByID(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/%s", c.Params("paymentId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PAYMENT", "find payment by id")
}

func (ph *PaymentHandler) HandleStripeWebhook(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/stripe/webhook")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PAYMENT", "handle stripe webhook")
}
