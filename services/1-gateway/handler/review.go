package handler

import "github.com/gofiber/fiber/v2"

type ReviewHandler struct {
	*BaseHandler
}

func NewReviewHandler(baseURL string) *ReviewHandler {
	return &ReviewHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *ReviewHandler) FindSellerReviews(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ReviewHandler) FindProductReviews(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ReviewHandler) Add(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ReviewHandler) Update(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ReviewHandler) Remove(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
