// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pbo/leo.proto

package pbo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Leo_Signup_FullMethodName = "/pbo.Leo/Signup"
	Leo_Login_FullMethodName  = "/pbo.Leo/Login"
)

// LeoClient is the client API for Leo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeoClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginReqest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type leoClient struct {
	cc grpc.ClientConnInterface
}

func NewLeoClient(cc grpc.ClientConnInterface) LeoClient {
	return &leoClient{cc}
}

func (c *leoClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, Leo_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leoClient) Login(ctx context.Context, in *LoginReqest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Leo_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeoServer is the server API for Leo service.
// All implementations must embed UnimplementedLeoServer
// for forward compatibility
type LeoServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginReqest) (*LoginResponse, error)
	mustEmbedUnimplementedLeoServer()
}

// UnimplementedLeoServer must be embedded to have forward compatible implementations.
type UnimplementedLeoServer struct {
}

func (UnimplementedLeoServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedLeoServer) Login(context.Context, *LoginReqest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLeoServer) mustEmbedUnimplementedLeoServer() {}

// UnsafeLeoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeoServer will
// result in compilation errors.
type UnsafeLeoServer interface {
	mustEmbedUnimplementedLeoServer()
}

func RegisterLeoServer(s grpc.ServiceRegistrar, srv LeoServer) {
	s.RegisterService(&Leo_ServiceDesc, srv)
}

func _Leo_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeoServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leo_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeoServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leo_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leo_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeoServer).Login(ctx, req.(*LoginReqest))
	}
	return interceptor(ctx, in, info, handler)
}

// Leo_ServiceDesc is the grpc.ServiceDesc for Leo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbo.Leo",
	HandlerType: (*LeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Leo_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Leo_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbo/leo.proto",
}
