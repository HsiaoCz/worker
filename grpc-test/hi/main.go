package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/HsiaoCz/worker/grpc-test/hi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "127.0.0.1:9001"
)

func HanleHello(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()
	resp, err := c.HandleHello(ctx, &pb.HelloRequest{
		Username: "zhangsan",
		Msg:      "hello my man",
		Content:  "what's up",
	})
	if err != nil {
		slog.Error("handle hello err:", err)
		return
	}
	fmt.Println("response msg:", resp.GetMsg())
	fmt.Println("response user:", resp.GetUsername())
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("grpc dial err:", err)
		return
	}
	defer conn.Close()
	HanleHello(pb.NewHelloClient(conn))
}
