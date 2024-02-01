package handlers

import (
	"net/http"

	"github.com/HsiaoCz/worker/west/model"
	"github.com/HsiaoCz/worker/west/utils"
	"github.com/HsiaoCz/worker/west/view/userv"
	"github.com/gin-gonic/gin"
)

type UserCase struct {
}

func (u UserCase) HandlerUserShow(c *gin.Context) {
	user := model.User{
		Username: "Jhon",
		Email:    "acid@g.com",
		Content:  "What's up SLG",
	}
	r := utils.New(c.Request.Context(), http.StatusOK, userv.Show(user))
	c.Render(http.StatusOK, r)
}

func (u UserCase) HandlerUserShowEmail(c *gin.Context) {
	user := model.User{
		Username: "Jhon",
		Email:    "acid@g.com",
		Content:  "What's up SLG",
	}
	r := utils.New(c.Request.Context(), http.StatusOK, userv.ShowUserInfo(user))
	c.Render(http.StatusOK, r)
}

func (u UserCase) HandlerUserShowContent(c *gin.Context) {
	user := model.User{
		Username: "Jhon",
		Email:    "acid@g.com",
		Content:  "What's up SLG",
	}
	r := utils.New(c.Request.Context(), http.StatusOK, userv.ShowUserContent(user))
	c.Render(http.StatusOK, r)
}
