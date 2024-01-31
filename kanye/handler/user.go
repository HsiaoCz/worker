package handler

import (
	"github.com/HsiaoCz/worker/kanye/models"
	"github.com/HsiaoCz/worker/kanye/view/userview"
	"github.com/gofiber/fiber/v2"
)

type UserCase struct {
}

func (u UserCase) HandlerUserShow(c *fiber.Ctx) error {
	user := models.User{
		Username: "bob",
		Password: "12333",
		Email:    "12333@g.com",
		Money:    "12345.99",
		Deal:     "120",
	}
	return userview.Show(user).Render(c.Context(), c)
}
