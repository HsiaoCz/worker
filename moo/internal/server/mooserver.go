package server

import (
	"net"

	"github.com/HsiaoCz/worker/moo/internal/pb"
	"github.com/HsiaoCz/worker/moo/internal/service"
	"google.golang.org/grpc"
)

func RegisterGrpc(addr string) error {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterMooServer(srv, service.NewMooService())
	return srv.Serve(conn)
}
