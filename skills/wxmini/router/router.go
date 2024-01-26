package router

import (
	"github.com/HsiaoCz/worker/skills/wxmini/logger"
	"github.com/gin-gonic/gin"
)

func Route(mode string, addr string) error {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/signup")
			user.POST("/login")
		}
	}
	return r.Run(addr)
}
