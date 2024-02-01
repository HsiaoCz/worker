package main

import (
	"github.com/HsiaoCz/worker/west/handlers"
	"github.com/HsiaoCz/worker/west/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)
	app.HTMLRender = utils.Default
	userCase := handlers.UserCase{}
	app.GET("/user", userCase.HandlerUserShow)
	app.GET("/user/email", userCase.HandlerUserShowEmail)
	app.GET("/user/content", userCase.HandlerUserShowContent)
	app.Run("192.168.206.1:9007")
}
