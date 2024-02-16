package main

import (
	"log/slog"

	"github.com/HsiaoCz/worker/moo/internal/server"
)

var (
	addr = "127.0.0.1:9003"
)

func main() {
	if err := server.RegisterGrpc(addr); err != nil {
		slog.Error("grpc start err:", err)
		return
	}
}
