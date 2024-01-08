// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: v1/fiber.proto

package v1

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
	FiberService_Signup_FullMethodName = "/v1.FiberService/Signup"
	FiberService_Login_FullMethodName  = "/v1.FiberService/Login"
)

// FiberServiceClient is the client API for FiberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FiberServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type fiberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFiberServiceClient(cc grpc.ClientConnInterface) FiberServiceClient {
	return &fiberServiceClient{cc}
}

func (c *fiberServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, FiberService_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fiberServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, FiberService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiberServiceServer is the server API for FiberService service.
// All implementations must embed UnimplementedFiberServiceServer
// for forward compatibility
type FiberServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedFiberServiceServer()
}

// UnimplementedFiberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFiberServiceServer struct {
}

func (UnimplementedFiberServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedFiberServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedFiberServiceServer) mustEmbedUnimplementedFiberServiceServer() {}

// UnsafeFiberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FiberServiceServer will
// result in compilation errors.
type UnsafeFiberServiceServer interface {
	mustEmbedUnimplementedFiberServiceServer()
}

func RegisterFiberServiceServer(s grpc.ServiceRegistrar, srv FiberServiceServer) {
	s.RegisterService(&FiberService_ServiceDesc, srv)
}

func _FiberService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiberServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FiberService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiberServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FiberService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiberServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FiberService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiberServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FiberService_ServiceDesc is the grpc.ServiceDesc for FiberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FiberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FiberService",
	HandlerType: (*FiberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _FiberService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _FiberService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/fiber.proto",
}