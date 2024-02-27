package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	addr string
	user *UserHandler
}

func NewHandler(addr string, user *UserHandler) *Handler {
	return &Handler{
		addr: addr,
		user: user,
	}
}

func (h *Handler) StartServer() error {
	r := mux.NewRouter()
	user := r.PathPrefix("/api/v1/user").Subrouter()
	user.HandleFunc("/login", h.user.HandleLogin)
	return http.ListenAndServe(h.addr, r)
}
