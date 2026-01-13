package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	BASE_PATH = "/api/v1/products"
)

func MainRouter(app *fiber.App, db *gorm.DB) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Product Service is healthy and OK.")
	})

	api := app.Group(BASE_PATH)
	api.Use(verifyGatewayReq)

	// TODO: Add routes
}

func verifyGatewayReq(c *fiber.Ctx) error {
	// TODO: Implement gateway verification
	return c.Next()
}
