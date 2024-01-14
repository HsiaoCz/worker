package main

import (
	"github.com/HsiaoCz/worker/binance/wssubstream/endpoints"
	handlemessage "github.com/HsiaoCz/worker/binance/wssubstream/handleMessage"
	"github.com/HsiaoCz/worker/binance/wssubstream/wss"
)

func main() {
	// wss.GetOrderBookDepth(endpoints.OrderBook, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"btcusdt@depth",
	// 	},
	// }, handlemessage.HandleOrderBookDepth)

	wss.WSGetBookTicker(endpoints.BookTicker, map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{
			"bnbusdt@bookTicker",
		},
		"id": "1",
	}, handlemessage.HandleBookTicker)
	// wss.GetCountSetSymbol(endpoints.CountSetSymbol, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbbtc@aggtrade",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleCountSet)
}
