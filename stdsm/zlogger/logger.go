package zlogger

import (
	"os"

	"github.com/HsiaoCz/worker/stdsm/dstu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(zcf dstu.ZapLoggerConfig, loggerLevel string, mode string) error {
	writerSyncer := getLogWriter(zcf)
	encode := getEncode()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(loggerLevel)); err != nil {
		return err
	}
	var core zapcore.Core
	if mode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encode, writerSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encode, writerSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
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

func getLogWriter(zcf dstu.ZapLoggerConfig) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   zcf.Filename,
		MaxSize:    zcf.MaxSize,
		MaxBackups: zcf.MaxBackup,
		MaxAge:     zcf.MaxAge,
	}
	return zapcore.AddSync(lumberjackLogger)
}
