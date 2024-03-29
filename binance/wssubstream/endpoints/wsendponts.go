package endpoints

var (
	BookTicker          = "wss://stream.binance.com:9443/ws/bnbusdt@bookTicker"
	OrderBook           = "wss://stream.binance.com:9443/ws/bnbbtc@depth"
	CountSetSymbol      = "wss://stream.binance.com:9443/ws/bnbbtc@aggTrade"
	EveryDealTrade      = "wss://stream.binance.com:9443/ws/bnbbtc@trade"
	Klines              = "wss://stream.binance.com:9443/ws/bnbbtc@kline_1"
	MinSymbolTicker     = "wss://stream.binance.com:9443/ws/bnbbtc@miniTicker"
	SymbolTicker        = "wss://stream.binance.com:9443/ws/bnbbtc@ticker"
	AllSymbolTicker     = "wss://stream.binance.com:9443/ws/!ticker@arr"
	AvgPrice            = "wss://stream.binance.com:9443/ws/btcusdt@avgPrice"
	SymbolDepthLevel    = "wss://stream.binance.com:9443/ws/btcusdt@depth20"
	SymbolAddDepth      = "wss://stream.binance.com:9443/ws/btcusdt@depth"
	SymbolTickerWindows = "wss://stream.binance.com:9443/ws/btcusdt@ticker_1h"
	AllMarketTicker     = "wss://stream.binance.com:9443/ws/!ticker_1h@arr"
)
