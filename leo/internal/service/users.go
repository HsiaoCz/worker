package service

import (
	"context"

	"github.com/HsiaoCz/worker/leo/internal/pbo"
)

type UserService struct {
	pbo.UnimplementedLeoServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Signup(ctx context.Context, in *pbo.SignupRequest) (*pbo.SignupResponse, error) {
	return &pbo.SignupResponse{}, nil
}

func (u *UserService) Login(ctx context.Context, in *pbo.LoginRequest) (*pbo.LoginResponse, error) {
	return &pbo.LoginResponse{}, nil
}
