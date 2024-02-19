package middleware

import "github.com/gofiber/fiber/v2"

func AutherMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
