package handlers

import "github.com/gofiber/fiber/v2"

type UserCase struct {
}

func (u UserCase) HandleLogin(c *fiber.Ctx) error {
	return nil
}
