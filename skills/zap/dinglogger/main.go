package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugarLogger *zap.SugaredLogger

// 定制logger
// zap.New() 来传递所有配置
// func New(core zapcore.Core, options ...Option) *Logger
// zapcore.Core需要三个配置
// Encoder:编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()
// WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去
// Log Level：哪种级别的日志将被写入
func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func InitLogger() {
	// encoder := getEncoder()
	// // test.log 记录全量日志
	// logF, _ := os.Create("./test.log")
	// c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// // test.err.log 记录ERROR级别的日志
	// errF, _ := os.Create("./test.err.log")
	// c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zapcore.ErrorLevel)
	// core := zapcore.NewTee(c1, c2)
	// logger = zap.New(core, zap.AddCaller())
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志文件的位置
		MaxSize:    10,           // 在进行切割之前 日志文件的最大大小 MB
		MaxBackups: 5,            // 保留旧文件的最大个数
		MaxAge:     30,           //保留旧文件的最大个数
		Compress:   false,        // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
