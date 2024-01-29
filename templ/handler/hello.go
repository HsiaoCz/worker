package handler

import (
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/worker/templ/hello"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	comp := hello.Hello(name)
	comp.Render(context.Background(), w)
}

func HandleButton(w http.ResponseWriter, r *http.Request) {
	text, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("read request body err:", err)
		return
	}

	button := hello.Button(string(text))
	button.Render(context.Background(), w)
}

func HandleImages(w http.ResponseWriter, r *http.Request) {
	comp := hello.Comp()
	comp.Render(context.Background(), w)
}
