package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var symbolTicker = "wss://stream.binance.com:9443/ws/bnbusdt@bookTicker"
var wg sync.WaitGroup

type BookTicker struct {
	UpdateId      int64  `json:"u"` // update id
	Symbol        string `json:"s"` // 交易对
	BuyBestPrice  string `json:"b"` // 买单最优挂单价格
	BuyBestCount  string `json:"B"` // 买单最优挂单数量
	SealBestPrice string `json:"a"` // 卖单最优挂单价格
	SealBestCount string `json:"A"` // 卖单最优挂单数据
}

func (b *BookTicker) String() string {
	bs, _ := json.Marshal(b)
	return string(bs)
}

func GetSymbolTicker(endpoint string, sendMessage map[string]any, handleMessage func(data *BookTicker)) {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 15 * time.Second,
	}
	conn, _, err := dialer.Dial(endpoint, nil)
	if err != nil {
		slog.Error("websocket dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		slog.Error("write json err:", err)
		return
	}
	var data = new(BookTicker)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(data); err != nil {
				slog.Error("read message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	wg.Wait()
}

func HandleMessage(data *BookTicker) {
	fmt.Println(data)
}

func main() {
	GetSymbolTicker(symbolTicker, map[string]any{
		"method": "SUBSCRIBE",
		"params": []string{
			"bnbusdt@bookTicker",
		},
		"id": "1",
	}, HandleMessage)
}
