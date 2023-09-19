// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: redis.proto

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
	RedisService_SetUser_FullMethodName = "/api.redis.v1.RedisService/SetUser"
	RedisService_GetUser_FullMethodName = "/api.redis.v1.RedisService/GetUser"
)

// RedisServiceClient is the client API for RedisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisServiceClient interface {
	SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*SetUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type redisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisServiceClient(cc grpc.ClientConnInterface) RedisServiceClient {
	return &redisServiceClient{cc}
}

func (c *redisServiceClient) SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*SetUserResponse, error) {
	out := new(SetUserResponse)
	err := c.cc.Invoke(ctx, RedisService_SetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, RedisService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedisServiceServer is the server API for RedisService service.
// All implementations must embed UnimplementedRedisServiceServer
// for forward compatibility
type RedisServiceServer interface {
	SetUser(context.Context, *SetUserRequest) (*SetUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedRedisServiceServer()
}

// UnimplementedRedisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRedisServiceServer struct {
}

func (UnimplementedRedisServiceServer) SetUser(context.Context, *SetUserRequest) (*SetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (UnimplementedRedisServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedRedisServiceServer) mustEmbedUnimplementedRedisServiceServer() {}

// UnsafeRedisServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisServiceServer will
// result in compilation errors.
type UnsafeRedisServiceServer interface {
	mustEmbedUnimplementedRedisServiceServer()
}

func RegisterRedisServiceServer(s grpc.ServiceRegistrar, srv RedisServiceServer) {
	s.RegisterService(&RedisService_ServiceDesc, srv)
}

func _RedisService_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisServiceServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisService_SetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisServiceServer).SetUser(ctx, req.(*SetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedisService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedisService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedisService_ServiceDesc is the grpc.ServiceDesc for RedisService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.redis.v1.RedisService",
	HandlerType: (*RedisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUser",
			Handler:    _RedisService_SetUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _RedisService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redis.proto",
}
