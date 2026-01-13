package handler

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
	*BaseHandler
}

func NewUserHandler(baseURL string) *UserHandler {
	return &UserHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *UserHandler) GetMyBuyerInfo(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) FindBuyerByUsername(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) UpdateBuyer(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) GetMySellerInfo(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) FindSellerByID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) FindSellerByUsername(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) GetRandomSellers(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *UserHandler) UpdateSeller(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
