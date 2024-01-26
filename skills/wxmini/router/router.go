package router

import "github.com/gin-gonic/gin"

func Route(addr string) error {
	r := gin.Default()
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
