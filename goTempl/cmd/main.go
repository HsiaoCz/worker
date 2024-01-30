package main

import (
	"github.com/HsiaoCz/worker/goTempl/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Start("192.168.206.1:9001")
}
