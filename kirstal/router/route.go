package router

import (
	"github.com/HsiaoCz/worker/kirstal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(r *fiber.App) {
	app := r.Group("/app")
	{
		v1 := app.Group("/v1")
		{
			auther := v1.Group("/auther")
			{
				auther.Post("/signup", handlers.HandleUserSignup)
				auther.Post("/login", handlers.HandleUserLogin)
			}
		}
	}
}
