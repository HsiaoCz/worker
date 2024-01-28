package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLogin struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
	Email    string `binding:"required,email" json:"email"`
}

type UserSignup struct {
	Username   string `binding:"required" json:"username"`
	Password   string `binding:"required" json:"password"`
	RePassword string `binding:"required,eqfiled=Password" json:"re_password"`
	Email      string `binding:"required,email" json:"email"`
}

func HandleUserSignup(c *gin.Context) {
	userS := new(UserSignup)
	if err := c.BindJSON(userS); err != nil {
		zap.L().Info("user signup err:", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
			"Code":    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Signup success",
		"Code":    http.StatusOK,
	})
}

func HandleUserLogin(c *gin.Context) {
	userL := new(UserLogin)
	if err := c.BindJSON(userL); err != nil {
		zap.L().Info("user login err:", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
			"Code":    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "login success",
		"Code":    http.StatusOK,
	})
}
