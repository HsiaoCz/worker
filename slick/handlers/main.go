package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type H map[string]any

func HandleHello(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("read request body err:", err)
		return
	}
	fmt.Println(string(body))
	ResponseJSON(w, http.StatusOK, H{
		"Message": "Hello",
		"data":    string(body),
	})
}

func main() {
	http.HandleFunc("/hello", HandleHello)
	http.ListenAndServe("192.168.206.1:9001", nil)
}

func ResponseJSON(w http.ResponseWriter, code int, v any) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
