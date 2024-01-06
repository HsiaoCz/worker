package main

import (
	"log"

	"github.com/HsiaoCz/worker/grpc/fiber-grpc/router"
)

var (
	addr = "127.0.0.1:9001"
)

func main() {
	if err := router.Router(addr); err != nil {
		log.Fatal(err)
	}
}
