package slogger

import (
	"log/slog"
	"os"
	"time"
)

var Logger *slog.Logger

func InitLogger(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	Logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,            // 输出日志的位置信息
		Level:     slog.LevelError, // 最低的日志等级
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))
	return nil
}
