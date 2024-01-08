package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// websocket server

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP,Hello"))
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(websocket.StatusInternalError, "内部错误")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		if err := wsjson.Read(ctx, conn, &v); err != nil {
			log.Fatal(err)
		}
		log.Printf("接收到客户端:%v\n", v)

		if err := wsjson.Write(ctx, conn, "Hello WebSocket Client"); err != nil {
			log.Fatal(err)
		}

		conn.Close(websocket.StatusAbnormalClosure, "")
	})
	log.Fatal(http.ListenAndServe(":3006", nil))
}
