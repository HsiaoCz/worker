package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var wsendopint = "wss://fstream.binance.com/stream?streams=btcusdt@markPrice/btcusdt@depth"

var wg sync.WaitGroup

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsendopint, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect successed:", wsendopint)
	var result map[string]any
	wg.Add(1)
	go func() {
		defer wg.Done()
		var flag int
		for {
			if flag == 19 {
				return
			}
			if err := conn.ReadJSON(&result); err != nil {
				log.Println("read json failed:", err)
				continue
			}
			fmt.Println(result)
		}
	}()
	wg.Wait()
}
