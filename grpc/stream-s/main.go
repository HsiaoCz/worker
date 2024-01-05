package main

import (
	"context"

	"github.com/HsiaoCz/worker/grpc/stream-s/pb"
)

type StreamService struct{
	pb.UnimplementedGreeterServer
}

func NewStreamService()*StreamService{
	return &StreamService{}
}

func (s *StreamService)Login(ctx context.Context,in *pb.LoginRequest)(*pb.LoginResponse,error){
	return &pb.LoginResponse{},nil
}

func (s *StreamService)Hello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloResponse,error){
	return &pb.HelloResponse{},nil
}

func main(){
	
}
