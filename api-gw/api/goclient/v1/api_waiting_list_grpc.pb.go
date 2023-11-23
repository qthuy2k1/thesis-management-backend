// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api_waiting_list.proto

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
	WaitingListService_CreateWaitingList_FullMethodName                 = "/api.waitingList.v1.WaitingListService/CreateWaitingList"
	WaitingListService_GetWaitingList_FullMethodName                    = "/api.waitingList.v1.WaitingListService/GetWaitingList"
	WaitingListService_UpdateWaitingList_FullMethodName                 = "/api.waitingList.v1.WaitingListService/UpdateWaitingList"
	WaitingListService_DeleteWaitingList_FullMethodName                 = "/api.waitingList.v1.WaitingListService/DeleteWaitingList"
	WaitingListService_GetWaitingListsOfClassroom_FullMethodName        = "/api.waitingList.v1.WaitingListService/GetWaitingListsOfClassroom"
	WaitingListService_GetWaitingLists_FullMethodName                   = "/api.waitingList.v1.WaitingListService/GetWaitingLists"
	WaitingListService_CheckUserInWaitingListOfClassroom_FullMethodName = "/api.waitingList.v1.WaitingListService/CheckUserInWaitingListOfClassroom"
)

// WaitingListServiceClient is the client API for WaitingListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WaitingListServiceClient interface {
	CreateWaitingList(ctx context.Context, in *CreateWaitingListRequest, opts ...grpc.CallOption) (*CreateWaitingListResponse, error)
	GetWaitingList(ctx context.Context, in *GetWaitingListRequest, opts ...grpc.CallOption) (*GetWaitingListResponse, error)
	UpdateWaitingList(ctx context.Context, in *UpdateWaitingListRequest, opts ...grpc.CallOption) (*UpdateWaitingListResponse, error)
	DeleteWaitingList(ctx context.Context, in *DeleteWaitingListRequest, opts ...grpc.CallOption) (*DeleteWaitingListResponse, error)
	GetWaitingListsOfClassroom(ctx context.Context, in *GetWaitingListsOfClassroomRequest, opts ...grpc.CallOption) (*GetWaitingListsOfClassroomResponse, error)
	GetWaitingLists(ctx context.Context, in *GetWaitingListsRequest, opts ...grpc.CallOption) (*GetWaitingListsResponse, error)
	CheckUserInWaitingListOfClassroom(ctx context.Context, in *CheckUserInWaitingListClassroomRequest, opts ...grpc.CallOption) (*CheckUserInWaitingListClassroomResponse, error)
}

type waitingListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWaitingListServiceClient(cc grpc.ClientConnInterface) WaitingListServiceClient {
	return &waitingListServiceClient{cc}
}

func (c *waitingListServiceClient) CreateWaitingList(ctx context.Context, in *CreateWaitingListRequest, opts ...grpc.CallOption) (*CreateWaitingListResponse, error) {
	out := new(CreateWaitingListResponse)
	err := c.cc.Invoke(ctx, WaitingListService_CreateWaitingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) GetWaitingList(ctx context.Context, in *GetWaitingListRequest, opts ...grpc.CallOption) (*GetWaitingListResponse, error) {
	out := new(GetWaitingListResponse)
	err := c.cc.Invoke(ctx, WaitingListService_GetWaitingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) UpdateWaitingList(ctx context.Context, in *UpdateWaitingListRequest, opts ...grpc.CallOption) (*UpdateWaitingListResponse, error) {
	out := new(UpdateWaitingListResponse)
	err := c.cc.Invoke(ctx, WaitingListService_UpdateWaitingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) DeleteWaitingList(ctx context.Context, in *DeleteWaitingListRequest, opts ...grpc.CallOption) (*DeleteWaitingListResponse, error) {
	out := new(DeleteWaitingListResponse)
	err := c.cc.Invoke(ctx, WaitingListService_DeleteWaitingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) GetWaitingListsOfClassroom(ctx context.Context, in *GetWaitingListsOfClassroomRequest, opts ...grpc.CallOption) (*GetWaitingListsOfClassroomResponse, error) {
	out := new(GetWaitingListsOfClassroomResponse)
	err := c.cc.Invoke(ctx, WaitingListService_GetWaitingListsOfClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) GetWaitingLists(ctx context.Context, in *GetWaitingListsRequest, opts ...grpc.CallOption) (*GetWaitingListsResponse, error) {
	out := new(GetWaitingListsResponse)
	err := c.cc.Invoke(ctx, WaitingListService_GetWaitingLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waitingListServiceClient) CheckUserInWaitingListOfClassroom(ctx context.Context, in *CheckUserInWaitingListClassroomRequest, opts ...grpc.CallOption) (*CheckUserInWaitingListClassroomResponse, error) {
	out := new(CheckUserInWaitingListClassroomResponse)
	err := c.cc.Invoke(ctx, WaitingListService_CheckUserInWaitingListOfClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaitingListServiceServer is the server API for WaitingListService service.
// All implementations must embed UnimplementedWaitingListServiceServer
// for forward compatibility
type WaitingListServiceServer interface {
	CreateWaitingList(context.Context, *CreateWaitingListRequest) (*CreateWaitingListResponse, error)
	GetWaitingList(context.Context, *GetWaitingListRequest) (*GetWaitingListResponse, error)
	UpdateWaitingList(context.Context, *UpdateWaitingListRequest) (*UpdateWaitingListResponse, error)
	DeleteWaitingList(context.Context, *DeleteWaitingListRequest) (*DeleteWaitingListResponse, error)
	GetWaitingListsOfClassroom(context.Context, *GetWaitingListsOfClassroomRequest) (*GetWaitingListsOfClassroomResponse, error)
	GetWaitingLists(context.Context, *GetWaitingListsRequest) (*GetWaitingListsResponse, error)
	CheckUserInWaitingListOfClassroom(context.Context, *CheckUserInWaitingListClassroomRequest) (*CheckUserInWaitingListClassroomResponse, error)
	mustEmbedUnimplementedWaitingListServiceServer()
}

// UnimplementedWaitingListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWaitingListServiceServer struct {
}

func (UnimplementedWaitingListServiceServer) CreateWaitingList(context.Context, *CreateWaitingListRequest) (*CreateWaitingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWaitingList not implemented")
}
func (UnimplementedWaitingListServiceServer) GetWaitingList(context.Context, *GetWaitingListRequest) (*GetWaitingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingList not implemented")
}
func (UnimplementedWaitingListServiceServer) UpdateWaitingList(context.Context, *UpdateWaitingListRequest) (*UpdateWaitingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWaitingList not implemented")
}
func (UnimplementedWaitingListServiceServer) DeleteWaitingList(context.Context, *DeleteWaitingListRequest) (*DeleteWaitingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWaitingList not implemented")
}
func (UnimplementedWaitingListServiceServer) GetWaitingListsOfClassroom(context.Context, *GetWaitingListsOfClassroomRequest) (*GetWaitingListsOfClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingListsOfClassroom not implemented")
}
func (UnimplementedWaitingListServiceServer) GetWaitingLists(context.Context, *GetWaitingListsRequest) (*GetWaitingListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaitingLists not implemented")
}
func (UnimplementedWaitingListServiceServer) CheckUserInWaitingListOfClassroom(context.Context, *CheckUserInWaitingListClassroomRequest) (*CheckUserInWaitingListClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserInWaitingListOfClassroom not implemented")
}
func (UnimplementedWaitingListServiceServer) mustEmbedUnimplementedWaitingListServiceServer() {}

// UnsafeWaitingListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WaitingListServiceServer will
// result in compilation errors.
type UnsafeWaitingListServiceServer interface {
	mustEmbedUnimplementedWaitingListServiceServer()
}

func RegisterWaitingListServiceServer(s grpc.ServiceRegistrar, srv WaitingListServiceServer) {
	s.RegisterService(&WaitingListService_ServiceDesc, srv)
}

func _WaitingListService_CreateWaitingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWaitingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).CreateWaitingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_CreateWaitingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).CreateWaitingList(ctx, req.(*CreateWaitingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_GetWaitingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).GetWaitingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_GetWaitingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).GetWaitingList(ctx, req.(*GetWaitingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_UpdateWaitingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWaitingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).UpdateWaitingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_UpdateWaitingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).UpdateWaitingList(ctx, req.(*UpdateWaitingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_DeleteWaitingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWaitingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).DeleteWaitingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_DeleteWaitingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).DeleteWaitingList(ctx, req.(*DeleteWaitingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_GetWaitingListsOfClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitingListsOfClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).GetWaitingListsOfClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_GetWaitingListsOfClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).GetWaitingListsOfClassroom(ctx, req.(*GetWaitingListsOfClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_GetWaitingLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaitingListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).GetWaitingLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_GetWaitingLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).GetWaitingLists(ctx, req.(*GetWaitingListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaitingListService_CheckUserInWaitingListOfClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserInWaitingListClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitingListServiceServer).CheckUserInWaitingListOfClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaitingListService_CheckUserInWaitingListOfClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitingListServiceServer).CheckUserInWaitingListOfClassroom(ctx, req.(*CheckUserInWaitingListClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WaitingListService_ServiceDesc is the grpc.ServiceDesc for WaitingListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WaitingListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.waitingList.v1.WaitingListService",
	HandlerType: (*WaitingListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWaitingList",
			Handler:    _WaitingListService_CreateWaitingList_Handler,
		},
		{
			MethodName: "GetWaitingList",
			Handler:    _WaitingListService_GetWaitingList_Handler,
		},
		{
			MethodName: "UpdateWaitingList",
			Handler:    _WaitingListService_UpdateWaitingList_Handler,
		},
		{
			MethodName: "DeleteWaitingList",
			Handler:    _WaitingListService_DeleteWaitingList_Handler,
		},
		{
			MethodName: "GetWaitingListsOfClassroom",
			Handler:    _WaitingListService_GetWaitingListsOfClassroom_Handler,
		},
		{
			MethodName: "GetWaitingLists",
			Handler:    _WaitingListService_GetWaitingLists_Handler,
		},
		{
			MethodName: "CheckUserInWaitingListOfClassroom",
			Handler:    _WaitingListService_CheckUserInWaitingListOfClassroom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_waiting_list.proto",
}
