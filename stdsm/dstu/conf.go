package dstu

type ZapLoggerConfig struct {
	Filename  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
}
