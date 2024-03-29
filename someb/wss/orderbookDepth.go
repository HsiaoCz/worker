package wss

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/someb/safe"
	"github.com/gorilla/websocket"
)

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
