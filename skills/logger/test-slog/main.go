package main

import (
	"log/slog"
	"net"
	"os"
)

func main() {

}

// slog info

func SwapErrorString() {
	slog.Info("this is first msg", "greeting", "hello,slog")
}

// slog json
func JSONError() {
	h := slog.NewTextHandler(os.Stderr, nil)
	l := slog.New(h)
	l.Info("greeting", "name", "tony")
	l.Error("oops", "err", net.ErrClosed, "status", 500)

	h1 := slog.NewJSONHandler(os.Stderr, nil)
	l1 := slog.New(h1)
	l1.Info("oops", "err", net.ErrClosed, "status", 500)
}

// log level
func LogLevel() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))
	slog.Info("open file for reading", "name", "foo.txt", "path", "/home/tonybai/demo/foo.txt")
	slog.Error("open file error", "err", os.ErrNotExist, "status", 2)
}
