package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	base_url string
}

func NewProductHandler(base_url string) *ProductHandler {
	return &ProductHandler{
		base_url: base_url,
	}
}

func (ph *ProductHandler) HealthCheck(c *fiber.Ctx) error {
	route := ph.base_url + "/health-check"
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	if err != nil {
		log.Printf("PRODUCT - health check error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (ph *ProductHandler) GetPopularProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/popular")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "get popular products")
}

func (ph *ProductHandler) FindProductByID(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/id/%s", c.Params("id"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "find product by id")
}

func (ph *ProductHandler) FindProductsByCategory(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/category/%s/%s/%s", c.Params("category"), c.Params("page"), c.Params("size"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "find products by category")
}

func (ph *ProductHandler) FindSimilarProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/similar/%s/%s/%s", c.Params("productId"), c.Params("page"), c.Params("size"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "find similar products")
}

func (ph *ProductHandler) ProductQuerySearch(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/search/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "product query search")
}

func (ph *ProductHandler) FindSellerActiveProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/sellers/active/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "find seller active products")
}

func (ph *ProductHandler) FindSellerInactiveProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/sellers/inactive/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "find seller inactive products")
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "create product")
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "update product")
}

func (ph *ProductHandler) ActivateProductStatus(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/update-status/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "activate product status")
}

func (ph *ProductHandler) DeactivateProductStatus(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "PRODUCT", "deactivate product status")
}
