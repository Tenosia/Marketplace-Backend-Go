package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	base_url string
}

func NewReviewHandler(base_url string) *ReviewHandler {
	return &ReviewHandler{
		base_url: base_url,
	}
}

func (rh *ReviewHandler) HealthCheck(c *fiber.Ctx) error {
	route := rh.base_url + "/health-check"
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	if err != nil {
		log.Printf("REVIEW - health check error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"response": string(body),
	})
}

func (rh *ReviewHandler) FindSellerReviews(c *fiber.Ctx) error {
	route := rh.base_url + fmt.Sprintf("/api/v1/reviews/seller/%s", c.Params("sellerId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "REVIEW", "find seller reviews")
}

func (rh *ReviewHandler) FindProductReviews(c *fiber.Ctx) error {
	route := rh.base_url + fmt.Sprintf("/api/v1/reviews/product/%s", c.Params("productId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "REVIEW", "find product reviews")
}

func (rh *ReviewHandler) Add(c *fiber.Ctx) error {
	route := rh.base_url + fmt.Sprintf("/api/v1/reviews")
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "REVIEW", "add review")
}

func (rh *ReviewHandler) Update(c *fiber.Ctx) error {
	route := rh.base_url + fmt.Sprintf("/api/v1/reviews/%s", c.Params("reviewId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "REVIEW", "update review")
}

func (rh *ReviewHandler) Remove(c *fiber.Ctx) error {
	route := rh.base_url + fmt.Sprintf("/api/v1/reviews/%s", c.Params("reviewId"))
	statusCode, body, err := sendHttpReqToAnotherService(c, route)
	return handleServiceResponse(c, statusCode, body, err, "REVIEW", "remove review")
}
