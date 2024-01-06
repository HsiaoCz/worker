package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/HsiaoCz/worker/grpc/stream-s/pb"
	"google.golang.org/grpc"
)

type StreamService struct {
	pb.UnimplementedGreeterServer
}

func NewStreamService() *StreamService {
	return &StreamService{}
}

func (s *StreamService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetPassword() != in.GetRepassword() {
		return &pb.LoginResponse{Msg: "请确认用户密码"}, nil
	}
	return &pb.LoginResponse{Msg: "登录成功"}, nil
}

func (s *StreamService) Hello(in *pb.HelloRequest, server pb.Greeter_HelloServer) error {
	fmt.Println(in.Content)
	words := []string{"你好", "hello", "halo", "sawadika", "bonjor"}
	for _, word := range words {
		data := &pb.HelloResponse{Something: word}
		if err := server.Send(data); err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) HandleHi(stream pb.Greeter_HandleHiServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			return stream.SendAndClose(&pb.HIResponse{
				Something: "你好,my man",
			})
		}
		fmt.Println(res.GetConent())
	}
}

func (s *StreamService) HandleChat(stream pb.Greeter_HandleChatServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			return stream.Send(&pb.ChatResponse{
				Anwser: "你好:" + res.GetName(),
				Pid:    GenPID(),
			})
		}
		fmt.Println()
	}
}

func GenPID() string {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(randm.Intn(1000000))
}

var (
	network = "tcp"
	addr    = "127.0.0.1:9001"
)

func main() {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, NewStreamService())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
