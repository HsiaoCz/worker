package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/HsiaoCz/worker/grpc/stream-c/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 一元grpc
func Login(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.Login(ctx, &pb.LoginRequest{Username: "zhangsan", Password: "1233abc", Repassword: "1233abc"})
	if err != nil {
		slog.Error("Login reply err:", err)
		return
	}
	fmt.Println(resp.GetMsg())
}

// 服务端流式grpc
func Hello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Hello(ctx, &pb.HelloRequest{Content: "hello,my name is zhangsan"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(res.GetSomething())
	}
}

// 客户端流式grpc
func HandleHi(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.HandleHi(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if err = stream.Send(&pb.HIRequest{Conent: GenHI()}); err != nil {
			log.Fatal(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetSomething())
}

// 双向流式grpc
func HandleChat(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.HandleChat(ctx)
	if err != nil {
		log.Fatal(err)
	}
	watic := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			if err == io.EOF {
				close(watic)
				return
			}
			fmt.Println("Anwser:", res.GetAnwser())
			fmt.Println("PID:", res.GetPid())
		}
	}()
	names := []string{"zhasss", "bob", "andi", "hall", "will"}
	for _, name := range names {
		if err = stream.Send(&pb.ChatRequest{Name: name, Chatcontent: GenHI()}); err != nil {
			log.Fatal(err)
		}
	}
	stream.CloseSend()
	<-watic
}

func GenHI() string {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(randm.Intn(1000000))
}

var (
	addr = "127.0.0.1:9001"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	Login(c)
	Hello(c)
	HandleChat(c)
	HandleHi(c)
}
