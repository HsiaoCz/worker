package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Post("/user/login")
	app.Listen("127.0.0.1:9001")
}
