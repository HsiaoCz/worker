package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// websocket client

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := wsjson.Write(ctx, c, "Hello WebSocket Server"); err != nil {
		log.Fatal(err)
	}

	var v interface{}
	if err := wsjson.Read(ctx, c, &v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("接收到服务端响应:%v\n", v)
	c.Close(websocket.StatusNormalClosure, "")
}
