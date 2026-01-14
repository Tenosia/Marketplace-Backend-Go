package handler

import (
	"fmt"

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
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PAYMENT - health check error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (ph *PaymentHandler) ProcessPayment(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/process")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PAYMENT - process payment error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *PaymentHandler) FindPaymentByID(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/%s", c.Params("paymentId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PAYMENT - find payment by id error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *PaymentHandler) HandleStripeWebhook(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/payments/stripe/webhook")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PAYMENT - handle stripe webhook error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}
