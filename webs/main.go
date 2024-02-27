package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(p))
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start the server on localhost port 8000 and log any errors
	println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
