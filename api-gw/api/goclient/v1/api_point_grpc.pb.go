// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api_point.proto

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
	PointService_CreateOrUpdatePointDef_FullMethodName = "/api.point.v1.PointService/CreateOrUpdatePointDef"
	PointService_GetAllPointDef_FullMethodName         = "/api.point.v1.PointService/GetAllPointDef"
)

// PointServiceClient is the client API for PointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointServiceClient interface {
	CreateOrUpdatePointDef(ctx context.Context, in *CreateOrUpdatePointDefRequest, opts ...grpc.CallOption) (*CreateOrUpdatePointDefResponse, error)
	GetAllPointDef(ctx context.Context, in *GetAllPointDefRequest, opts ...grpc.CallOption) (*GetAllPointDefResponse, error)
}

type pointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPointServiceClient(cc grpc.ClientConnInterface) PointServiceClient {
	return &pointServiceClient{cc}
}

func (c *pointServiceClient) CreateOrUpdatePointDef(ctx context.Context, in *CreateOrUpdatePointDefRequest, opts ...grpc.CallOption) (*CreateOrUpdatePointDefResponse, error) {
	out := new(CreateOrUpdatePointDefResponse)
	err := c.cc.Invoke(ctx, PointService_CreateOrUpdatePointDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetAllPointDef(ctx context.Context, in *GetAllPointDefRequest, opts ...grpc.CallOption) (*GetAllPointDefResponse, error) {
	out := new(GetAllPointDefResponse)
	err := c.cc.Invoke(ctx, PointService_GetAllPointDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointServiceServer is the server API for PointService service.
// All implementations must embed UnimplementedPointServiceServer
// for forward compatibility
type PointServiceServer interface {
	CreateOrUpdatePointDef(context.Context, *CreateOrUpdatePointDefRequest) (*CreateOrUpdatePointDefResponse, error)
	GetAllPointDef(context.Context, *GetAllPointDefRequest) (*GetAllPointDefResponse, error)
	mustEmbedUnimplementedPointServiceServer()
}

// UnimplementedPointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPointServiceServer struct {
}

func (UnimplementedPointServiceServer) CreateOrUpdatePointDef(context.Context, *CreateOrUpdatePointDefRequest) (*CreateOrUpdatePointDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdatePointDef not implemented")
}
func (UnimplementedPointServiceServer) GetAllPointDef(context.Context, *GetAllPointDefRequest) (*GetAllPointDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPointDef not implemented")
}
func (UnimplementedPointServiceServer) mustEmbedUnimplementedPointServiceServer() {}

// UnsafePointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointServiceServer will
// result in compilation errors.
type UnsafePointServiceServer interface {
	mustEmbedUnimplementedPointServiceServer()
}

func RegisterPointServiceServer(s grpc.ServiceRegistrar, srv PointServiceServer) {
	s.RegisterService(&PointService_ServiceDesc, srv)
}

func _PointService_CreateOrUpdatePointDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdatePointDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).CreateOrUpdatePointDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PointService_CreateOrUpdatePointDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).CreateOrUpdatePointDef(ctx, req.(*CreateOrUpdatePointDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetAllPointDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPointDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetAllPointDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PointService_GetAllPointDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetAllPointDef(ctx, req.(*GetAllPointDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PointService_ServiceDesc is the grpc.ServiceDesc for PointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.point.v1.PointService",
	HandlerType: (*PointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdatePointDef",
			Handler:    _PointService_CreateOrUpdatePointDef_Handler,
		},
		{
			MethodName: "GetAllPointDef",
			Handler:    _PointService_GetAllPointDef_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_point.proto",
}
