package handler

import (
	"fmt"

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
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - health check error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (ph *ProductHandler) GetPopularProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/popular")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - get popular products error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) FindProductByID(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/id/%s", c.Params("id"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - find product by id error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) FindProductsByCategory(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/category/%s/%s/%s", c.Params("category"), c.Params("page"), c.Params("size"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - find products by category error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) FindSimilarProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/similar/%s/%s/%s", c.Params("productId"), c.Params("page"), c.Params("size"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - find similar products error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) ProductQuerySearch(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/search/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - product query search error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) FindSellerActiveProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/sellers/active/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - find seller active products error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) FindSellerInactiveProducts(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/sellers/inactive/%s/%s", c.Params("page"), c.Params("size"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - find seller inactive products error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products")
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - create product error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - update product error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) ActivateProductStatus(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/update-status/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - activate product status error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}

func (ph *ProductHandler) DeactivateProductStatus(c *fiber.Ctx) error {
	route := ph.base_url + fmt.Sprintf("/api/v1/products/%s/%s", c.Params("sellerId"), c.Params("productId"))
	statusCode, body, errs := sendHttpReqToAnotherService(c, route)
	if len(errs) > 0 {
		fmt.Println("PRODUCT - deactivate product status error", errs)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	return c.Status(statusCode).Send(body)
}
