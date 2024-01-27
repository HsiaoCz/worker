package main

import (
	"log/slog"

	"github.com/HsiaoCz/worker/skills/wxmini/router"
)

var (
	mode = "debug"
	addr = "127.0.0.1:9003"
)

func main() {
	if err := router.Route(mode, addr); err != nil {
		slog.Error("listen addr err:", err)
		return
	}
}
