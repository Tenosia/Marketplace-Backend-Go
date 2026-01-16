package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// sendHttpReqToAnotherService sends an HTTP request to another microservice
// Returns status code, response body, and error
func sendHttpReqToAnotherService(c *fiber.Ctx, url string) (int, []byte, error) {
	// Create context with timeout for the request
	ctx, cancel := context.WithTimeout(c.UserContext(), 30*time.Second)
	defer cancel()

	// Acquire agent and ensure it's released
	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	// Set timeout for the request
	a.Timeout(30 * time.Second)

	req := a.Request()
	req.Header.SetMethod(c.Method())
	req.SetRequestURI(url)

	// Add gateway token if available
	if gatewayToken, ok := c.UserContext().Value("gatewayToken").(string); ok && gatewayToken != "" {
		req.Header.Add("gatewayToken", gatewayToken)
	}

	// Extract token from cookie or Authorization header
	tokenStr := c.Cookies("token", "")
	if tokenStr == "" {
		authHeader := c.Get("Authorization", "")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) > 1 {
				tokenStr = parts[1]
			}
		}
	}
	if tokenStr != "" {
		a.Cookie("token", tokenStr)
	}

	// Set request body and headers if present
	if len(c.Body()) > 0 {
		contentType := c.Get("Content-Type", "application/json")
		a.ContentType(contentType)
		req.Header.Add("Accept", c.Get("Accept", "application/json"))
		a.Body(c.Body())
	}

	// Parse the request
	if err := a.Parse(); err != nil {
		log.Printf("Failed to parse HTTP request to %s: %v", url, err)
		return 0, nil, fmt.Errorf("failed to parse request: %w", err)
	}

	// Execute the request with context
	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		log.Printf("HTTP request to %s failed: %v", url, errs)
		return statusCode, body, fmt.Errorf("request failed: %v", errs)
	}

	// Check if context was cancelled
	if ctx.Err() != nil {
		return 0, nil, fmt.Errorf("request timeout: %w", ctx.Err())
	}

	return statusCode, body, nil
}

// handleServiceResponse is a helper function to handle responses from microservices
func handleServiceResponse(c *fiber.Ctx, statusCode int, body []byte, err error, serviceName string, operation string) error {
	if err != nil {
		log.Printf("%s - %s error: %v", serviceName, operation, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if statusCode >= 400 {
		return c.Status(statusCode).Send(body)
	}

	return c.Status(statusCode).Send(body)
}
