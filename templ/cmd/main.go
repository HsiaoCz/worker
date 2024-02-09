package main

import (
	"net/http"

	"github.com/HsiaoCz/worker/templ/handler"
	"github.com/HsiaoCz/worker/templ/logger"
)

var addr = "192.168.206.1:9002"

func main() {
	logger.InitLogger()
	http.HandleFunc("/user", handler.HandleHello)
	http.HandleFunc("/button", handler.HandleButton)
	http.HandleFunc("/images", handler.HandleImages)
	logger.Logger.Info("the server is runnig", "port", addr)
	http.ListenAndServe(addr, nil)
}
