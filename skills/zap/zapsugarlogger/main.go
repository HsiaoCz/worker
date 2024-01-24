package main

import (
	"net/http"

	"go.uber.org/zap"
)

var sugaredLogegr *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugaredLogegr.Sync()
	SimpleHttpGet("www.google.com")
	SimpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ := zap.NewProduction()
	sugaredLogegr = logger.Sugar()
}

func SimpleHttpGet(url string) {
	sugaredLogegr.Debug("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugaredLogegr.Errorf("Error fetching URL %s: Error=%s", url, err)
	} else {
		sugaredLogegr.Info("Success statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
