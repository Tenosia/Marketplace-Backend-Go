package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/marketplace-go-backend/services/1-gateway/handler"
	"github.com/marketplace-go-backend/services/1-gateway/types"
	"github.com/marketplace-go-backend/services/1-gateway/util"
	"github.com/gofiber/fiber/v2"
)

var (
	BASE_PATH = "/api/v1/gateway"
)

func generateGatewayToken(c *fiber.Ctx) error {
	gatewaySecret := os.Getenv("GATEWAY_TOKEN")
	gatewayToken, err := util.GenerateJWT(gatewaySecret)
	if err != nil {
		fmt.Printf("generating gateway token error:\n%+v", err)
		return fiber.NewError(http.StatusInternalServerError, "Unexpected error happened.")
	}

	ctx := context.WithValue(c.UserContext(), "gatewayToken", gatewayToken)
	c.SetUserContext(ctx)
	return c.Next()
}

func authOnly(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	if tokenStr == "" {
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(strings.Split(authHeader, " ")) == 0 {
			return fiber.NewError(http.StatusUnauthorized, "sign in first")
		}
		tokenStr = strings.Split(authHeader, " ")[1]
	}

	token, err := util.VerifyingJWT(os.Getenv("JWT_SECRET"), tokenStr)
	if err != nil {
		fmt.Printf("authOnly error:\n%+v", err)
		return fiber.NewError(http.StatusUnauthorized, "sign in first")
	}

	claims, ok := token.Claims.(*types.JWTClaims)
	if !ok {
		fmt.Println("token is not matched with claims:", token.Claims)
		return fiber.NewError(http.StatusUnauthorized, "sign in first")
	}

	c.SetUserContext(context.WithValue(c.UserContext(), "current_user", claims))
	return c.Next()
}

func MainRouter(app *fiber.App) {
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("API Gateway Service is health and OK!")
	})

	AUTH_URL := os.Getenv("AUTH_URL")
	USER_URL := os.Getenv("USER_URL")
	PRODUCT_URL := os.Getenv("PRODUCT_URL")
	ORDER_URL := os.Getenv("ORDER_URL")
	PAYMENT_URL := os.Getenv("PAYMENT_URL")
	REVIEW_URL := os.Getenv("REVIEW_URL")
	api := app.Group(BASE_PATH)
	api.Use(generateGatewayToken)

	authRouter(AUTH_URL, api.Group("/auths"))
	userRouter(USER_URL, api.Group("/users"))
	productRouter(PRODUCT_URL, api.Group("/products"))
	orderRouter(ORDER_URL, api.Group("/orders"))
	paymentRouter(PAYMENT_URL, api.Group("/payments"))
	reviewRouter(REVIEW_URL, api.Group("/reviews"))

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Resource is not found")
	})
}

func authRouter(base_url string, r fiber.Router) {
	ah := handler.NewAuthHandler(base_url)
	r.Get("/health-check", ah.HealthCheck)

	r.Get("/google/:action", ah.AuthWithGoogle)
	r.Get("/signup/google-callback", ah.SignUpWithGoogle)
	r.Post("/signup", ah.SignUp).Name("signup")
	r.Get("/signin/google-callback", ah.SignInWithGoogle)
	r.Post("/signin", ah.SignIn).Name("signin")
	r.Patch("/forgot-password/:email", ah.SendForgotPasswordURL)
	r.Patch("/reset-password/:token", ah.ResetPassword)

	r.Use(authOnly)
	r.Get("/user-info", ah.GetUserInfo)
	r.Get("/refresh-token/:username", ah.RefreshToken)
	r.Post("/send-verification-email", ah.SendVerifyEmailURL)
	r.Patch("/verify-email/:token", ah.VerifyEmail)
	r.Patch("/change-password", ah.ChangePassword)
}

func userRouter(base_url string, r fiber.Router) {
	uh := handler.NewUserHandler(base_url)
	r.Get("/health-check", uh.HealthCheck)

	r.Use(authOnly)
	r.Get("/buyers/my-info", uh.GetMyBuyerInfo)
	r.Get("/buyers/:username", uh.FindBuyerByUsername)
	r.Put("/buyers", uh.UpdateBuyer)

	r.Get("/sellers/my-info", uh.GetMySellerInfo)
	r.Get("/sellers/id/:id", uh.FindSellerByID)
	r.Get("/sellers/username/:username", uh.FindSellerByUsername)
	r.Get("/sellers/random/:count", uh.GetRandomSellers)
	r.Post("/sellers", uh.Create)
	r.Put("/sellers", uh.UpdateSeller)
}

func productRouter(base_url string, r fiber.Router) {
	ph := handler.NewProductHandler(base_url)
	r.Get("/health-check", ph.HealthCheck)

	r.Get("/popular", ph.GetPopularProducts).Name("home")
	r.Get("/id/:id", ph.FindProductByID)
	r.Get("/category/:category/:page/:size", ph.FindProductsByCategory)
	r.Get("/similar/:productId/:page/:size", ph.FindSimilarProducts)
	r.Get("/search/:page/:size", ph.ProductQuerySearch)

	r.Use(authOnly)
	r.Get("/sellers/active/:page/:size", ph.FindSellerActiveProducts)
	r.Get("/sellers/inactive/:page/:size", ph.FindSellerInactiveProducts)
	r.Post("", ph.CreateProduct)
	r.Put("/:sellerId/:productId", ph.UpdateProduct)
	r.Patch("/update-status/:sellerId/:productId", ph.ActivateProductStatus)
	r.Delete("/:sellerId/:productId", ph.DeactivateProductStatus)
}

func orderRouter(base_url string, r fiber.Router) {
	oh := handler.NewOrderHandler(base_url)
	r.Get("/health-check", oh.HealthCheck)

	r.Use(authOnly)
	r.Get("/:id", oh.FindOrderByID)
	r.Get("/buyer/my-orders", oh.FindOrdersByBuyerID)
	r.Get("/seller/my-orders", oh.FindOrdersBySellerID)
	r.Get("/buyer/my-orders-notifications", oh.FindMyOrdersNotifications)
	r.Post("", oh.CreateOrder)
	r.Patch("/:orderId/status", oh.UpdateOrderStatus)
	r.Post("/:orderId/cancel", oh.CancelOrder)
}

func paymentRouter(base_url string, r fiber.Router) {
	ph := handler.NewPaymentHandler(base_url)
	r.Get("/health-check", ph.HealthCheck)

	r.Use(authOnly)
	r.Post("/process", ph.ProcessPayment)
	r.Get("/:paymentId", ph.FindPaymentByID)
	r.Post("/stripe/webhook", ph.HandleStripeWebhook)
}

func reviewRouter(base_url string, r fiber.Router) {
	rh := handler.NewReviewHandler(base_url)
	r.Get("/health-check", rh.HealthCheck)

	r.Use(authOnly)
	r.Get("/seller/:sellerId", rh.FindSellerReviews)
	r.Get("/product/:productId", rh.FindProductReviews)
	r.Post("", rh.Add)
	r.Patch("/:reviewId", rh.Update)
	r.Delete("/:reviewId", rh.Remove)
}
