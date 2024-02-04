package main

import (
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/worker/stdsm/controllers"
	"github.com/HsiaoCz/worker/stdsm/slogger"
)

var addr = "192.168.206.1:9009"

func main() {
	if err := slogger.InitLogger("./stdsm.log"); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	slogger.Logger.Info("the http server is running", "Prot:", addr)
	http.HandleFunc("/user", controllers.HandleUserShow)
	http.ListenAndServe(addr, nil)
}
