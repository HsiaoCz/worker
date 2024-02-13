package wss

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/someb/data"
	"github.com/HsiaoCz/worker/someb/safe"
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
