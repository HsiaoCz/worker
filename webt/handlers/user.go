package handlers

import (
	"net/http"

	"github.com/HsiaoCz/worker/webt/store"
)

type UserHandler struct {
	store store.Store
}

func NewUserHandler(store store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {}
