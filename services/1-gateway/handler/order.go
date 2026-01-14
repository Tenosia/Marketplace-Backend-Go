package handler

import (
	"fmt"

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
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - health check error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (oh *OrderHandler) FindOrderByID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s", c.Params("id"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - find order by id error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) FindOrdersByBuyerID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/buyer/my-orders")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - find orders by buyer id error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) FindOrdersBySellerID(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/seller/my-orders")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - find orders by seller id error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) FindMyOrdersNotifications(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/buyer/my-orders-notifications")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - find my orders notifications error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - create order error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s/status", c.Params("orderId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - update order status error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (oh *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	route := oh.base_url + fmt.Sprintf("/api/v1/orders/%s/cancel", c.Params("orderId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("ORDER - cancel order error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}
