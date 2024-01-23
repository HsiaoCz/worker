package main

import (
	"context"
	"log/slog"
	"net"
	"os"
)

func main() {
	DaynamicSetLoggerLevel()
}

// 性能建议
// 使用Logger.With避免重复格式化公共属性字段，这使得处理程序可以缓存格式化结果。
// 将昂贵的计算推迟到日志输出时再进行，例如传递指针而不是格式化后的字符串。这可以避免在禁用的日志行上进行不必要的工作。
// 对于昂贵的值，可以实现LogValuer接口，这样在输出时可以进行lazy加载计算
func DaynamicSetLoggerLevel() {
	var lvl slog.LevelVar
	lvl.Set(slog.LevelDebug)
	opts := slog.HandlerOptions{
		Level: &lvl,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))

	// 输出日志
	slog.Info("greeting", "name", "tom")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops", slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	// 重新设置日志级别 LevelError
	lvl.Set(slog.LevelError)
	slog.Info("greeting", "name", "tom")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops", slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
