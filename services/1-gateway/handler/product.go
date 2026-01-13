package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler struct {
	*BaseHandler
}

func NewProductHandler(baseURL string) *ProductHandler {
	return &ProductHandler{
		BaseHandler: NewBaseHandler(baseURL),
	}
}

func (h *ProductHandler) GetPopularProducts(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) FindProductByID(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) FindProductsByCategory(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) FindSimilarProducts(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) ProductQuerySearch(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) FindSellerActiveProducts(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) FindSellerInactiveProducts(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) ActivateProductStatus(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}

func (h *ProductHandler) DeactivateProductStatus(c *fiber.Ctx) error {
	// TODO: Implement
	return c.SendStatus(501)
}
