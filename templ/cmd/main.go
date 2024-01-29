package main

import (
	"net/http"

	"github.com/HsiaoCz/worker/templ/handler"
)

func main() {
	http.HandleFunc("/user", handler.HandleHello)
	http.HandleFunc("/button", handler.HandleButton)
	http.HandleFunc("/images", handler.HandleImages)
	http.ListenAndServe("192.168.206.1:9002", nil)
}
