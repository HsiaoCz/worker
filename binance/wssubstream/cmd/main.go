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

	// wss.WSGetBookTicker(endpoints.BookTicker, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbusdt@bookTicker",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleBookTicker)
	// wss.GetCountSetSymbol(endpoints.CountSetSymbol, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbbtc@aggtrade",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleCountSet)
	// wss.GetDealTrade(endpoints.EveryDealTrade, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbbtc@trade",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleDealTrade)
	// wss.GetKlines(endpoints.Klines, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbbtc@kline_1",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleKlines)

	// wss.GetMiniSymbolTicker(endpoints.MinSymbolTicker, map[string]any{
	// 	"method": "SUBSCRIBE",
	// 	"params": []string{
	// 		"bnbbtc@miniTicker",
	// 	},
	// 	"id": "1",
	// }, handlemessage.HandleMiniSymbolTicker)

	wss.GetSymbolTicker(endpoints.SymbolTicker, map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{
			"bnbbtc@ticker",
		},
		"id": "1",
	}, handlemessage.HandleSymbolTicker)
}
