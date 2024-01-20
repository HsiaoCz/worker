package service

import (
	"context"

	v1 "github.com/HsiaoCz/worker/grpc/fiber-grpc/internal/v1"
)

type FiberService struct {
	v1.UnimplementedFiberServiceServer
}

func NewFiberService() *FiberService {
	return &FiberService{}
}

func (f *FiberService) Login(context context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
}
func (f *FiberService) Signup(context context.Context, in *v1.SignupRequest) (*v1.SignupResponse, error) {
}
