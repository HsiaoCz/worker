package handler

import (
	"context"
	"net/http"

	"github.com/HsiaoCz/worker/templ/hello"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	comp := hello.Hello(name)
	comp.Render(context.Background(), w)
}
