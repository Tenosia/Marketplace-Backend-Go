package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	base_url string
}

func NewOrderHandler(base_url string) *OrderHandler {
	return &OrderHandler{
		base_url: base_url,
	}
}

func (oh *OrderHandler) HealthCheck(c *fiber.Ctx) error {
	route := oh.base_url + "/health-check"
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	if err != nil {
		log.Printf("ORDER - health check error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (oh *OrderHandler) FindOrderByID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s", c.Params("id"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "find order by id")
}

func (oh *OrderHandler) FindOrdersByBuyerID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/buyer/my-orders")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "find orders by buyer id")
}

func (oh *OrderHandler) FindOrdersBySellerID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/seller/my-orders")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "find orders by seller id")
}

func (oh *OrderHandler) FindMyOrdersNotifications(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/buyer/my-orders-notifications")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "find my orders notifications")
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "create order")
}

func (oh *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s/status", c.Params("orderId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "update order status")
}

func (oh *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s/cancel", c.Params("orderId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "ORDER", "cancel order")
}
