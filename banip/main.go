package main

import (
	"github.com/HsiaoCz/worker/banip/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	user := app.Group("/user")
	{
		user.Post("/register", api.HandleUserRegister)
		user.Post("/login", api.HandleUserLogin)
		user.Post("/content", api.HandleUserCreateContent)
	}
	app.Listen("127.0.0.1:9001")
}
