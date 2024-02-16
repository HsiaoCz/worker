package service

import (
	"context"

	"github.com/HsiaoCz/worker/moo/internal/pb"
)

type MooService struct {
	pb.UnimplementedMooServer
}

func NewMooService() *MooService {
	return &MooService{}
}

func (m *MooService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetUsername() != "张三" || in.GetPassword() != "12306" {
		return &pb.LoginResponse{Code: 10001, Msg: "登录失败,请检查用户名或密码!"}, nil
	}
	return &pb.LoginResponse{Code: 10000, Msg: "登录成功!"}, nil
}
