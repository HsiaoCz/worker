package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

// 定制logger
// 将日志写入文件 而不是写入到控制台
// 使用zap.New()来传递所有的配置
// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel

// Encoder:编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()
// WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去
// Log Level：哪种级别的日志将被写入

func InitLogger() {
	writeSyncer := getWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getWriter() zapcore.WriteSyncer {
	file, _ := os.Open("./test.log")
	return zapcore.AddSync(file)
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	SimpleHttpGet("www.google.com")
	SimpleHttpGet("http://www.google.com")
}

func SimpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
