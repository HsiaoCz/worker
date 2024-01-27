package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignup struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Email      string `json:"email"`
}

func HandleUserSignup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Signup success",
		"Code":    http.StatusOK,
	})
}

func HandleUserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "login success",
		"Code":    http.StatusOK,
	})
}
