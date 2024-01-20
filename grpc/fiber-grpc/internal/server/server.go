package server

import (
	"log"
	"net"

	"github.com/HsiaoCz/worker/grpc/fiber-grpc/internal/service"
	v1 "github.com/HsiaoCz/worker/grpc/fiber-grpc/internal/v1"
	"google.golang.org/grpc"
)

func RegisterGrpc(network string, addr string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		log.Println("dial err:", err)
		return err
	}
	srv := grpc.NewServer()
	v1.RegisterFiberServiceServer(srv, service.NewFiberService())
	return srv.Serve(conn)
}
