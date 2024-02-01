package main

import (
	"os"
	"time"

	"github.com/HsiaoCz/worker/lucky/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	use := handlers.UserCse{}
	app.Get("/user", use.HandlePost)
	app.Listen("192.168.206.1:5738")
}
