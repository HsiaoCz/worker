package main

import (
	"log/slog"
	"net/http"

	"github.com/HsiaoCz/worker/stdsm/controllers"
	"github.com/HsiaoCz/worker/stdsm/dstu"
	"github.com/HsiaoCz/worker/stdsm/zlogger"
	"go.uber.org/zap"
)

var addr = "192.168.206.1:9009"

func main() {
	// if err := slogger.InitLogger("./stdsm.log"); err != nil {
	// 	slog.Error("init logger err:", err)
	// 	return
	// }
	if err := zlogger.Init(dstu.ZapLoggerConfig{
		Filename:  "stdsm.log",
		MaxSize:   200,
		MaxBackup: 7,
		MaxAge:    30,
	}, "debug", "dev"); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	zap.L().Info("the http server is running", zap.String("Addr", addr))
	http.HandleFunc("/user", controllers.HandleUserShow)
	http.ListenAndServe(addr, nil)
}
