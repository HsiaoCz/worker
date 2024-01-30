package handler

import (
	"github.com/HsiaoCz/worker/goTempl/model"
	"github.com/HsiaoCz/worker/goTempl/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "a@gg.com",
	}
	return render(c, user.Show(u))
}
