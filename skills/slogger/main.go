package main

import (
	"io"
	"log/slog"
	"os"
	"time"
)

var logger *slog.Logger

func main() {
	InitLogger("./slog.log")
	logger.Error("something is very very very very", "auther", "HsiaoCz")
}

func InitLogger(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		slog.Error("open file err:", err)
		return
	}
	logger = slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, file), &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))
}
