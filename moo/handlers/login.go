package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Code":    http.StatusBadRequest,
			"Message": err.Error(),
		})
	}
	return nil
}
