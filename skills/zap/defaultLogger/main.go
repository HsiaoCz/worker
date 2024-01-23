package main

import (
	"log"
	"log/slog"
	"os"
)

func SetDefalutLogger() {
	logFileLocation, err := os.OpenFile("./logger.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		slog.Error("open file err:", err)
		return
	}
	log.SetOutput(logFileLocation)
}

