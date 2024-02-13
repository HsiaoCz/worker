package main

import (
	"github.com/HsiaoCz/worker/someb/endpoints"
	handlemessage "github.com/HsiaoCz/worker/someb/handleMessage"
	"github.com/HsiaoCz/worker/someb/wss"
)

func main() {
	wss.WSGetBookTicker(endpoints.BookTicker, map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{
			"bnbusdt@bookTicker",
		},
		"id": "1",
	}, handlemessage.HandleSymbolTicker)
}
