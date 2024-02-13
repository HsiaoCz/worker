package slogger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func InitLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}))
	slog.SetDefault(logger)
}
