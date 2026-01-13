package main

import (
	"github.com/marketplace-go-backend/services/4-user/handler"
	"github.com/marketplace-go-backend/services/4-user/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	BASE_PATH = "/api/v1/users"
)

func MainRouter(app *fiber.App, db *gorm.DB, ccs *handler.GRPCClients) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("User Service is healthy and OK.")
	})

	api := app.Group(BASE_PATH)
	api.Use(verifyGatewayReq)

	us := service.NewUserService(db)
	uh := handler.NewUserHttpHandler(us, ccs)

	// TODO: Add routes
	_ = uh
}

func verifyGatewayReq(c *fiber.Ctx) error {
	// TODO: Implement gateway verification
	return c.Next()
}
