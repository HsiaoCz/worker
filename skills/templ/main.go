package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/worker/skills/templ/view"
)

func main() {
	http.HandleFunc("/name", HandlePostName)
	http.ListenAndServe("192.168.206.1:9001", nil)
}

func HandlePostName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("read request body err:", err)
		return
	}
	var message map[string]any
	if err := json.Unmarshal(body, &message); err != nil {
		slog.Error("json unmarshal err:", err)
		return
	}
	name := message["name"]
	username, ok := name.(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	view.Hello(username).Render(r.Context(), w)
}
