// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api_thesis_commitee.proto

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
	CommiteeService_CreateCommitee_FullMethodName = "/api.commitee.v1.CommiteeService/CreateCommitee"
	CommiteeService_GetCommitee_FullMethodName    = "/api.commitee.v1.CommiteeService/GetCommitee"
	CommiteeService_UpdateCommitee_FullMethodName = "/api.commitee.v1.CommiteeService/UpdateCommitee"
	CommiteeService_DeleteCommitee_FullMethodName = "/api.commitee.v1.CommiteeService/DeleteCommitee"
	CommiteeService_GetCommitees_FullMethodName   = "/api.commitee.v1.CommiteeService/GetCommitees"
)

// CommiteeServiceClient is the client API for CommiteeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommiteeServiceClient interface {
	CreateCommitee(ctx context.Context, in *CreateCommiteeRequest, opts ...grpc.CallOption) (*CreateCommiteeResponse, error)
	GetCommitee(ctx context.Context, in *GetCommiteeRequest, opts ...grpc.CallOption) (*GetCommiteeResponse, error)
	UpdateCommitee(ctx context.Context, in *UpdateCommiteeRequest, opts ...grpc.CallOption) (*UpdateCommiteeResponse, error)
	DeleteCommitee(ctx context.Context, in *DeleteCommiteeRequest, opts ...grpc.CallOption) (*DeleteCommiteeResponse, error)
	GetCommitees(ctx context.Context, in *GetCommiteesRequest, opts ...grpc.CallOption) (*GetCommiteesResponse, error)
}

type commiteeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommiteeServiceClient(cc grpc.ClientConnInterface) CommiteeServiceClient {
	return &commiteeServiceClient{cc}
}

func (c *commiteeServiceClient) CreateCommitee(ctx context.Context, in *CreateCommiteeRequest, opts ...grpc.CallOption) (*CreateCommiteeResponse, error) {
	out := new(CreateCommiteeResponse)
	err := c.cc.Invoke(ctx, CommiteeService_CreateCommitee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commiteeServiceClient) GetCommitee(ctx context.Context, in *GetCommiteeRequest, opts ...grpc.CallOption) (*GetCommiteeResponse, error) {
	out := new(GetCommiteeResponse)
	err := c.cc.Invoke(ctx, CommiteeService_GetCommitee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commiteeServiceClient) UpdateCommitee(ctx context.Context, in *UpdateCommiteeRequest, opts ...grpc.CallOption) (*UpdateCommiteeResponse, error) {
	out := new(UpdateCommiteeResponse)
	err := c.cc.Invoke(ctx, CommiteeService_UpdateCommitee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commiteeServiceClient) DeleteCommitee(ctx context.Context, in *DeleteCommiteeRequest, opts ...grpc.CallOption) (*DeleteCommiteeResponse, error) {
	out := new(DeleteCommiteeResponse)
	err := c.cc.Invoke(ctx, CommiteeService_DeleteCommitee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commiteeServiceClient) GetCommitees(ctx context.Context, in *GetCommiteesRequest, opts ...grpc.CallOption) (*GetCommiteesResponse, error) {
	out := new(GetCommiteesResponse)
	err := c.cc.Invoke(ctx, CommiteeService_GetCommitees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommiteeServiceServer is the server API for CommiteeService service.
// All implementations must embed UnimplementedCommiteeServiceServer
// for forward compatibility
type CommiteeServiceServer interface {
	CreateCommitee(context.Context, *CreateCommiteeRequest) (*CreateCommiteeResponse, error)
	GetCommitee(context.Context, *GetCommiteeRequest) (*GetCommiteeResponse, error)
	UpdateCommitee(context.Context, *UpdateCommiteeRequest) (*UpdateCommiteeResponse, error)
	DeleteCommitee(context.Context, *DeleteCommiteeRequest) (*DeleteCommiteeResponse, error)
	GetCommitees(context.Context, *GetCommiteesRequest) (*GetCommiteesResponse, error)
	mustEmbedUnimplementedCommiteeServiceServer()
}

// UnimplementedCommiteeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommiteeServiceServer struct {
}

func (UnimplementedCommiteeServiceServer) CreateCommitee(context.Context, *CreateCommiteeRequest) (*CreateCommiteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommitee not implemented")
}
func (UnimplementedCommiteeServiceServer) GetCommitee(context.Context, *GetCommiteeRequest) (*GetCommiteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitee not implemented")
}
func (UnimplementedCommiteeServiceServer) UpdateCommitee(context.Context, *UpdateCommiteeRequest) (*UpdateCommiteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommitee not implemented")
}
func (UnimplementedCommiteeServiceServer) DeleteCommitee(context.Context, *DeleteCommiteeRequest) (*DeleteCommiteeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommitee not implemented")
}
func (UnimplementedCommiteeServiceServer) GetCommitees(context.Context, *GetCommiteesRequest) (*GetCommiteesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitees not implemented")
}
func (UnimplementedCommiteeServiceServer) mustEmbedUnimplementedCommiteeServiceServer() {}

// UnsafeCommiteeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommiteeServiceServer will
// result in compilation errors.
type UnsafeCommiteeServiceServer interface {
	mustEmbedUnimplementedCommiteeServiceServer()
}

func RegisterCommiteeServiceServer(s grpc.ServiceRegistrar, srv CommiteeServiceServer) {
	s.RegisterService(&CommiteeService_ServiceDesc, srv)
}

func _CommiteeService_CreateCommitee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommiteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommiteeServiceServer).CreateCommitee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommiteeService_CreateCommitee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommiteeServiceServer).CreateCommitee(ctx, req.(*CreateCommiteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommiteeService_GetCommitee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommiteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommiteeServiceServer).GetCommitee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommiteeService_GetCommitee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommiteeServiceServer).GetCommitee(ctx, req.(*GetCommiteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommiteeService_UpdateCommitee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommiteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommiteeServiceServer).UpdateCommitee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommiteeService_UpdateCommitee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommiteeServiceServer).UpdateCommitee(ctx, req.(*UpdateCommiteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommiteeService_DeleteCommitee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommiteeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommiteeServiceServer).DeleteCommitee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommiteeService_DeleteCommitee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommiteeServiceServer).DeleteCommitee(ctx, req.(*DeleteCommiteeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommiteeService_GetCommitees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommiteesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommiteeServiceServer).GetCommitees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommiteeService_GetCommitees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommiteeServiceServer).GetCommitees(ctx, req.(*GetCommiteesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommiteeService_ServiceDesc is the grpc.ServiceDesc for CommiteeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommiteeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.commitee.v1.CommiteeService",
	HandlerType: (*CommiteeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommitee",
			Handler:    _CommiteeService_CreateCommitee_Handler,
		},
		{
			MethodName: "GetCommitee",
			Handler:    _CommiteeService_GetCommitee_Handler,
		},
		{
			MethodName: "UpdateCommitee",
			Handler:    _CommiteeService_UpdateCommitee_Handler,
		},
		{
			MethodName: "DeleteCommitee",
			Handler:    _CommiteeService_DeleteCommitee_Handler,
		},
		{
			MethodName: "GetCommitees",
			Handler:    _CommiteeService_GetCommitees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_thesis_commitee.proto",
}