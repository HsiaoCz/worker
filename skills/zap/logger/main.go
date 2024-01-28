package logger

import (
	"os"

	"github.com/HsiaoCz/worker/skills/wxmini/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() error {
	cfg := conf.Conf
	writerSyncer := getLogWriter(
		cfg.LogConfig.Filename,
		cfg.LogConfig.MaxSize,
		cfg.LogConfig.MaxBackup,
		cfg.LogConfig.MaxAge,
	)
	encoder := getEncode()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(cfg.LogConfig.Level)); err != nil {
		return err
	}
	var core zapcore.Core
	if cfg.AppConfig.Mode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writerSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
	// 替换zap库里面的全局logger
	// 后续在其他包中只需要使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return nil
}

func getEncode() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
