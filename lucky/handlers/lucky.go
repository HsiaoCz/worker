package handlers

import (
	"net/http"

	"github.com/HsiaoCz/worker/lucky/models"
	"github.com/gofiber/fiber/v2"
)

type UserCse struct{}

func (u UserCse) HandlePost(c *fiber.Ctx) error {
	post := models.Post{}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
			"Code":    http.StatusBadRequest,
		})
	}
	if post.RequestText != "query luck value" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "please input the currect request text",
			"Code":    http.StatusBadRequest,
		})
	}

	return nil
}
