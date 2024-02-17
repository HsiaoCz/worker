package main

import (
	"os"
	"time"

	"github.com/HsiaoCz/worker/leo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	addr = "127.0.0.1:9001"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))
	app.Post("/user/signup", handlers.HandleUserSignup)
	app.Post("/user/login", handlers.HandleUserLogin)
	app.Listen(addr)
}
