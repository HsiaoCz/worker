package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

// 日志级别
// INFO Debug Warn Error
func main() {
	// 不携带上下文的Info
	slog.Info("Hello", "Count", 3)
	slog.InfoContext(context.Background(), "slog msg with context", "name", "tom")
	slog.Log(context.Background(), slog.LevelInfo, "Hello", "Count", 3)
	slog.LogAttrs(context.Background(), slog.LevelInfo, "敏感数据", slog.Any("password", Password("123456789")))
}

func FormatLogger() {
	// 更改日志格式
	textLogger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	textLogger.InfoContext(context.Background(), "TextHandler", "name", "tom")
	// JSON格式的日志
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	jsonLogger.InfoContext(context.Background(), "hello", "Count", 3)
}

// 默认的logger slog.Default

func SetLogger() {
	// 创建一个JsonHandler 处理slog 对象
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// 将这个slog 对象设置为默认的实例
	slog.SetDefault(jsonLogger)
	slog.InfoContext(context.Background(), "JsonHandler", "name", "tom")
}

// 日志分组
func GroupLogger() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, nil)).WithGroup("information")
	jsonLogger.InfoContext(context.Background(), "json-log", slog.String("name", "tom"), slog.Int("phone", 1234567))

	textLogger := slog.New(slog.NewTextHandler(os.Stderr, nil)).WithGroup("information")
	textLogger.InfoContext(context.Background(), "text-log", slog.String("name", "tom"))
}

// 高效输出日志
// 使用其他方法如Info输出日志时，内部会将键值对转成Attr类型，
// 而使用LogAttrs方法，可以直接指定Attr类型，减少转换的过程，因此会更高效
func HighEffect() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	jsonLogger.LogAttrs(context.Background(), slog.LevelInfo, "高效输出日志", slog.String("name", "tom"), slog.Int("phone", 123456789))
}

// slog 提供了LogValue 接口 可以在输出的时候进行延迟加载计算

type Password string

func (p Password) LogValue() slog.Value {
	return slog.StringValue("READCTED_PASSWORD")
}

// 设置一些属性
// slog 自带了一些属性
/*
// log/slog/handler.go

// Keys for "built-in" attributes.
const (
    // TimeKey is the key used by the built-in handlers for the time
    // when the log method is called. The associated Value is a [time.Time].
    TimeKey = "time"
    // LevelKey is the key used by the built-in handlers for the level
    // of the log call. The associated value is a [Level].
    LevelKey = "level"
    // MessageKey is the key used by the built-in handlers for the
    // message of the log call. The associated value is a string.
    MessageKey = "msg"
    // SourceKey is the key used by the built-in handlers for the source file
    // and line of the log call. The associated value is a string.
    SourceKey = "source"
)
*/

func SetConfig() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	// 返回新的logger实例
	// 这里的with可以设置 logger 实例
	logger := jsonLogger.With("requestID", "12")
	logger.LogAttrs(context.Background(), slog.LevelInfo, "json-log", slog.String("k1", "v1"))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "json-log", slog.String("k2", "v2"))
}

// 配置日志处理器
func ConfigLogController() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,            // 日志输出的位置信息
		Level:     slog.LevelError, // 设置最低日志等级
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey { // 格式化key 为"time"的属性值
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}))
	jsonLogger.InfoContext(context.Background(), "json-log", slog.String("name", "tom"))
	jsonLogger.InfoContext(context.Background(), "json-log", slog.String("name", "tom"))
}
