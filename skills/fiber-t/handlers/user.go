package handlers

import (
	"github.com/HsiaoCz/worker/skills/fiber-t/storage"
	"github.com/gofiber/fiber/v2"
)

type UserCase struct {
	db storage.DBs
}

func (u UserCase) HandleLogin(c *fiber.Ctx) error {
	return nil
}
