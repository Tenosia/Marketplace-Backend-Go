package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	base_url string
}

func NewUserHandler(base_url string) *UserHandler {
	return &UserHandler{
		base_url: base_url,
	}
}

func (uh *UserHandler) HealthCheck(c *fiber.Ctx) error {
	route := uh.base_url + "/health-check"
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	if err != nil {
		log.Printf("USER - health check error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (uh *UserHandler) GetMyBuyerInfo(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/buyers/my-info")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "get my buyer info")
}

func (uh *UserHandler) FindBuyerByUsername(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/buyers/%s", c.Params("username"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "find buyer by username")
}

func (uh *UserHandler) GetMySellerInfo(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers/my-info")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "get my seller info")
}

func (uh *UserHandler) FindSellerByID(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers/id/%s", c.Params("id"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "find seller by id")
}

func (uh *UserHandler) FindSellerByUsername(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers/username/%s", c.Params("username"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "find seller by username")
}

func (uh *UserHandler) GetRandomSellers(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers/random/%s", c.Params("count"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "get random sellers")
}

func (uh *UserHandler) Create(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "create seller")
}

func (uh *UserHandler) UpdateSeller(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/sellers")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "update seller")
}

func (uh *UserHandler) UpdateBuyer(c *fiber.Ctx) error {
	route := uh.base_url + fmt.Sprintf("/api/v1/users/buyers")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "USER", "update buyer")
}
