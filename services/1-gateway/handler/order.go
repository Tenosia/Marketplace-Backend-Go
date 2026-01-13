package handler

import "github.com/gofiber/fiber/v2"

type OrderHandler struct {
	*BaseHandler
}

func NewOrderHandler(baseURL string) *OrderHandler {
	return &OrderHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *OrderHandler) FindOrderByID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) FindOrdersByBuyerID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) FindOrdersBySellerID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) FindMyOrdersNotifications(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
