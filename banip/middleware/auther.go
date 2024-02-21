package middleware

import "github.com/gofiber/fiber/v2"

func AutherJWT(c *fiber.Ctx) error {
	return c.Next()
}
