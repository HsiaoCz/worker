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
	return render(c, userview.Show(user))
}

func (u UserCase) HandleUserShowEmial(c *fiber.Ctx) error {
	user := models.User{
		Username: "bob",
		Password: "12333",
		Email:    "12333@g.com",
		Money:    "12345.99",
		Deal:     "120",
	}
	c.Set("Content-Type", "text/html")
	return userview.ShowUserInfo(user).Render(c.Context(), c.Response().BodyWriter())
}
