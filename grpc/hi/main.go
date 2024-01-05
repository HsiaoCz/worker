package main

import (
	"context"
	"log"
	"time"

	"github.com/HsiaoCz/worker/grpc/hi/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "127.0.0.1:9001"
)

func SayHello(cc pb.HelloClient) error {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cc.HandleHello(ctx, &pb.HelloRequest{Username: "zhangsan", Msg: "hello my man", Content: "what's up"})
	if err != nil {
		return err
	}
	log.Println(r.GetUsername(), r.GetMsg())
	return nil
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if err := SayHello(pb.NewHelloClient(conn)); err != nil {
		log.Fatal(err)
	}
}
