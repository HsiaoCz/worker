package handler

import (
	"net/http"

	"github.com/HsiaoCz/worker/stdtempl/data"
	"github.com/HsiaoCz/worker/stdtempl/view/userve"
)

type Usercase struct{}

func (u Usercase) Show(w http.ResponseWriter, r *http.Request) {
	user := data.User{
		Username: "zhangs",
		Content:  "what's up slg",
		Email:    "acid@g.com",
	}
	userve.Show(user).Render(r.Context(), w)
}
