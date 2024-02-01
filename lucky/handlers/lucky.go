package handlers

import "github.com/gofiber/fiber/v2"

type UserCse struct{}

func (u UserCse) HandlePost(c *fiber.Ctx) error {
	return nil
}
