package main

import (
	"log/slog"

	"github.com/HsiaoCz/worker/skills/wxmini/conf"
	"github.com/HsiaoCz/worker/skills/wxmini/logger"
	"github.com/HsiaoCz/worker/skills/wxmini/router"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置文件
	if err := conf.Init(); err != nil {
		zap.L().Error("init conf failed,err:%v\n", zap.Error(err))
		return
	} else {
		zap.L().Info("init conf successed")
	}
	// 初始化日志
	if err := logger.Init(); err != nil {
		zap.L().Error("init logger failed,err:%v\n", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	if err := router.Route(conf.Conf.AppConfig.Mode, conf.Conf.AppConfig.Port); err != nil {
		slog.Error("listen addr err:", err)
		return
	}
}
