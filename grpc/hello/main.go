package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/HsiaoCz/worker/grpc/hello/pb"
	"google.golang.org/grpc"
)

type Hello struct {
	Username string
	Msg      string
	Content  string
}

type HelloService struct {
	pb.UnimplementedHelloServer
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h *HelloService) HandleHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	hello := Hello{
		Username: in.GetUsername(),
		Msg:      in.GetMsg(),
		Content:  in.GetContent(),
	}
	fmt.Println(hello)
	return &pb.HelloResponse{Username: "lisi", Msg: "Hi,"}, nil
}

var (
	network = "tcp"
	addr    = "127.0.0.1:9001"
)

func main() {
	conn, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, NewHelloService())
	log.Fatal(server.Serve(conn))
}
