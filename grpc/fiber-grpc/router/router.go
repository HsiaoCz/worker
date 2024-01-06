package router

import (
	"github.com/HsiaoCz/worker/grpc/fiber-grpc/controller"
	"github.com/gofiber/fiber/v2"
)

func Router(addr string) error {
	r := fiber.New()
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.Post("/signup", controller.HandleUserSignup)
				user.Post("/login", controller.HandleUserLogin)
			}
		}
	}
	return r.Listen(addr)
}
