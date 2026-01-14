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

type SellerHandler struct {
	sellerSvc *svc.SellerService
	buyerSvc  *svc.BuyerService
}

func NewSellerHandler(buyerSvc *svc.BuyerService, sellerSvc *svc.SellerService) *SellerHandler {
	return &SellerHandler{
		sellerSvc: sellerSvc,
		buyerSvc:  buyerSvc,
	}
}

func (sh *SellerHandler) GetMySellerInfo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	userInfo, ok := c.UserContext().Value("current_user").(*types.JWTClaims)
	if !ok {
		return fiber.NewError(http.StatusUnauthorized, "sign-in first")
	}

	// TODO: Implement FindSellerByUsername in service
	_ = ctx
	_ = userInfo
	_ = errors.Is(gorm.ErrRecordNotFound, nil)
	return c.SendStatus(501)
}

func (sh *SellerHandler) FindSellerByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	// TODO: Implement FindSellerByID in service
	_ = ctx
	_ = c.Params("id")
	_ = errors.Is(gorm.ErrRecordNotFound, nil)
	return c.SendStatus(501)
}

func (sh *SellerHandler) FindSellerByUsername(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	// TODO: Implement FindSellerByUsername in service
	_ = ctx
	_ = c.Params("username")
	_ = errors.Is(gorm.ErrRecordNotFound, nil)
	return c.SendStatus(501)
}

func (sh *SellerHandler) GetRandomSellers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	// TODO: Implement GetRandomSellers in service
	_ = ctx
	_ = c.Params("count")
	return c.SendStatus(501)
}

func (sh *SellerHandler) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 5*time.Second)
	defer cancel()

	userInfo, ok := c.UserContext().Value("current_user").(*types.JWTClaims)
	if !ok {
		return fiber.NewError(http.StatusBadRequest, "invalid data. Please re-signin")
	}

	// TODO: Implement Create in service
	_ = ctx
	_ = userInfo
	return c.SendStatus(501)
}

func (sh *SellerHandler) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 5*time.Second)
	defer cancel()

	userInfo, ok := c.UserContext().Value("current_user").(*types.JWTClaims)
	if !ok {
		return fiber.NewError(http.StatusBadRequest, "invalid data. Please re-signin")
	}

	// TODO: Implement Update in service
	_ = ctx
	_ = userInfo
	_ = errors.Is(gorm.ErrRecordNotFound, nil)
	return c.SendStatus(501)
}
