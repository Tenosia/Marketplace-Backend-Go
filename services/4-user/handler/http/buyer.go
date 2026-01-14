package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	svc "github.com/marketplace-go-backend/services/4-user/service"
	"github.com/marketplace-go-backend/services/4-user/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BuyerHandler struct {
	buyerSvc *svc.BuyerService
}

func NewBuyerHandler(buyerSvc *svc.BuyerService) *BuyerHandler {
	return &BuyerHandler{
		buyerSvc: buyerSvc,
	}
}

func (bh *BuyerHandler) GetMyBuyerInfo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	userInfo, ok := c.UserContext().Value("current_user").(*types.JWTClaims)
	if !ok {
		return fiber.NewError(http.StatusUnauthorized, "invalid data. Please re-signin")
	}

	// TODO: Implement FindBuyerByID in service
	_ = ctx
	_ = userInfo
	return c.SendStatus(501)
}

func (bh *BuyerHandler) FindBuyerByUsername(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	// TODO: Implement FindBuyerByEmailOrUsername in service
	_ = ctx
	_ = c.Params("username")
	return c.SendStatus(501)
}

func (bh *BuyerHandler) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	userInfo, ok := c.UserContext().Value("current_user").(*types.JWTClaims)
	if !ok {
		return fiber.NewError(http.StatusUnauthorized, "invalid data. Please re-signin")
	}

	// TODO: Implement Update in service
	_ = ctx
	_ = userInfo
	_ = errors.Is(gorm.ErrRecordNotFound, nil)
	return c.SendStatus(501)
}
