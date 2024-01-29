package main

import (
	"net/http"

	"github.com/HsiaoCz/worker/templ/handler"
)

func main() {
	http.HandleFunc("/user", handler.HandleHello)
	http.ListenAndServe("127.0.0.1:9001", nil)
}
