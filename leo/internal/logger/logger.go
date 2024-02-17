package logger

import (
	"log/slog"
	"os"
	"time"
)

var logger *slog.Logger

func InitLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey { // 格式化key 为"time"的属性值
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))
	slog.SetDefault(logger)
}
