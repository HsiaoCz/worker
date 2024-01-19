package wss

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/binance/wssubstream/data"
	"github.com/HsiaoCz/worker/binance/wssubstream/safe"
	"github.com/gorilla/websocket"
)

func WSGetBookTicker(endpoint string, sendMessage map[string]any, handleMessage func(data *data.BookTicker)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("send message err:", err)
		return
	}
	var data = new(data.BookTicker)
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(data); err != nil {
				log.Println("Read message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

func GetOrderBookDepth(endpoint string, sendMessage map[string]any, handleMessage func(message map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("send message err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("recv message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

// 归集交易流
func GetCountSetSymbol(endpoint string, sendMessage map[string]any, handleMessage func(message map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("send message err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("recv message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

// 逐笔交易
func GetDealTrade(endpoint string, sendMessage map[string]any, handleMessage func(message map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("Send message err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read json err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

// K线
func GetKlines(endpoint string, sendMessage map[string]any, handleMessage func(message map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("Send message err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read json err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

func GetMiniSymbolTicker(endpoint string, sendMessage map[string]any, handleMessage func(data map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("Dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("Send message err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read json err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

func GetSymbolTicker(endpoint string, sendMessage map[string]any, handleMessage func(data map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("write json err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

// all ticker
func GetAllTicker(endpoint string, sendMessage map[string]any, handleMessage func(data []any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("write json err:", err)
		return
	}
	var data []any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}

// avg price
func GetAvgPrice(endpoint string, sendMessage map[string]any, handleMessage func(data map[string]any)) {
	dial := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dial.Dial(endpoint, nil)
	if err != nil {
		log.Println("dial err:", err)
		return
	}
	if err := conn.WriteJSON(sendMessage); err != nil {
		log.Println("write json err:", err)
		return
	}
	var data map[string]any
	safe.Wg.Add(1)
	go func() {
		defer safe.Wg.Done()
		var flag int
		for {
			if flag == 19 {
				break
			}
			if err := conn.ReadJSON(&data); err != nil {
				log.Println("read message err:", err)
				flag++
				continue
			}
			handleMessage(data)
		}
	}()
	safe.Wg.Wait()
}
