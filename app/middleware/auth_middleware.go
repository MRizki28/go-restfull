package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	authentication := true
	_ = authentication

	if !authentication {
		return c.Status(401).SendString("Unauthorized")
	}
	return c.Next()
}
