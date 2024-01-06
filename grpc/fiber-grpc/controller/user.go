package controller

import "github.com/gofiber/fiber/v2"

type UserSignup struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Email      string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandleUserSignup(c *fiber.Ctx) error {
	return nil
}

func HandleUserLogin(c *fiber.Ctx) error {
	return nil
}
