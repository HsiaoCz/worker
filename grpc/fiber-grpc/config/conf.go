package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(FiberGrpc)

type FiberGrpc struct {
	App   AppConf   `mapstructure:"app"`
	Mysql MysqlConf `mapstructure:"mysql"`
	Log   LogConf   `mapstructure:"log"`
}

type AppConf struct {
	Addr      string `mapstructure:"addr"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`
}
type MysqlConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
}
type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func ParseConf() (err error) {
	viper.SetConfigName("fiber-grpc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("viper read config err:", zap.Error(err))
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal config err:", zap.Error(err))
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config changed....")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("reunmarshal config err:", zap.Error(err))
			return
		}
	})
	return err
}
