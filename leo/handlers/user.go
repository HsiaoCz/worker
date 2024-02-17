package handlers

import (
	"net/http"
	"strings"

	"github.com/HsiaoCz/worker/leo/post"
	"github.com/gofiber/fiber/v2"
)

func HandleUserSignup(c *fiber.Ctx) error {
	users := new(post.UserSignup)
	if err := c.BodyParser(users); err != nil {
		return err
	}
	if users.Password != users.RePassword {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Code":    http.StatusOK,
			"Message": "请检查用户名或密码",
		})
	}
	if !strings.Contains(users.Email, "@") {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Code":    http.StatusOK,
			"Message": "请检查邮箱是否符合规范",
		})
	}
    
	return nil
}

func HandleUserLogin(c *fiber.Ctx) error {
	userl := new(post.UserLogin)
	if err := c.BodyParser(userl); err != nil {
		return err
	}

	return nil
}
