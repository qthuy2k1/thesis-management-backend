// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: classroom.proto

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
	ClassroomService_CreateClassroom_FullMethodName      = "/classroom.v1.ClassroomService/CreateClassroom"
	ClassroomService_GetClassroom_FullMethodName         = "/classroom.v1.ClassroomService/GetClassroom"
	ClassroomService_UpdateClassroom_FullMethodName      = "/classroom.v1.ClassroomService/UpdateClassroom"
	ClassroomService_DeleteClassroom_FullMethodName      = "/classroom.v1.ClassroomService/DeleteClassroom"
	ClassroomService_GetClassrooms_FullMethodName        = "/classroom.v1.ClassroomService/GetClassrooms"
	ClassroomService_CheckClassroomExists_FullMethodName = "/classroom.v1.ClassroomService/CheckClassroomExists"
	ClassroomService_GetLecturerClassroom_FullMethodName = "/classroom.v1.ClassroomService/GetLecturerClassroom"
)

// ClassroomServiceClient is the client API for ClassroomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassroomServiceClient interface {
	CreateClassroom(ctx context.Context, in *CreateClassroomRequest, opts ...grpc.CallOption) (*CreateClassroomResponse, error)
	GetClassroom(ctx context.Context, in *GetClassroomRequest, opts ...grpc.CallOption) (*GetClassroomResponse, error)
	UpdateClassroom(ctx context.Context, in *UpdateClassroomRequest, opts ...grpc.CallOption) (*UpdateClassroomResponse, error)
	DeleteClassroom(ctx context.Context, in *DeleteClassroomRequest, opts ...grpc.CallOption) (*DeleteClassroomResponse, error)
	GetClassrooms(ctx context.Context, in *GetClassroomsRequest, opts ...grpc.CallOption) (*GetClassroomsResponse, error)
	CheckClassroomExists(ctx context.Context, in *CheckClassroomExistsRequest, opts ...grpc.CallOption) (*CheckClassroomExistsResponse, error)
	GetLecturerClassroom(ctx context.Context, in *GetLecturerClassroomRequest, opts ...grpc.CallOption) (*GetLecturerClassroomResponse, error)
}

type classroomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassroomServiceClient(cc grpc.ClientConnInterface) ClassroomServiceClient {
	return &classroomServiceClient{cc}
}

func (c *classroomServiceClient) CreateClassroom(ctx context.Context, in *CreateClassroomRequest, opts ...grpc.CallOption) (*CreateClassroomResponse, error) {
	out := new(CreateClassroomResponse)
	err := c.cc.Invoke(ctx, ClassroomService_CreateClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) GetClassroom(ctx context.Context, in *GetClassroomRequest, opts ...grpc.CallOption) (*GetClassroomResponse, error) {
	out := new(GetClassroomResponse)
	err := c.cc.Invoke(ctx, ClassroomService_GetClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) UpdateClassroom(ctx context.Context, in *UpdateClassroomRequest, opts ...grpc.CallOption) (*UpdateClassroomResponse, error) {
	out := new(UpdateClassroomResponse)
	err := c.cc.Invoke(ctx, ClassroomService_UpdateClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) DeleteClassroom(ctx context.Context, in *DeleteClassroomRequest, opts ...grpc.CallOption) (*DeleteClassroomResponse, error) {
	out := new(DeleteClassroomResponse)
	err := c.cc.Invoke(ctx, ClassroomService_DeleteClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) GetClassrooms(ctx context.Context, in *GetClassroomsRequest, opts ...grpc.CallOption) (*GetClassroomsResponse, error) {
	out := new(GetClassroomsResponse)
	err := c.cc.Invoke(ctx, ClassroomService_GetClassrooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) CheckClassroomExists(ctx context.Context, in *CheckClassroomExistsRequest, opts ...grpc.CallOption) (*CheckClassroomExistsResponse, error) {
	out := new(CheckClassroomExistsResponse)
	err := c.cc.Invoke(ctx, ClassroomService_CheckClassroomExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classroomServiceClient) GetLecturerClassroom(ctx context.Context, in *GetLecturerClassroomRequest, opts ...grpc.CallOption) (*GetLecturerClassroomResponse, error) {
	out := new(GetLecturerClassroomResponse)
	err := c.cc.Invoke(ctx, ClassroomService_GetLecturerClassroom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassroomServiceServer is the server API for ClassroomService service.
// All implementations must embed UnimplementedClassroomServiceServer
// for forward compatibility
type ClassroomServiceServer interface {
	CreateClassroom(context.Context, *CreateClassroomRequest) (*CreateClassroomResponse, error)
	GetClassroom(context.Context, *GetClassroomRequest) (*GetClassroomResponse, error)
	UpdateClassroom(context.Context, *UpdateClassroomRequest) (*UpdateClassroomResponse, error)
	DeleteClassroom(context.Context, *DeleteClassroomRequest) (*DeleteClassroomResponse, error)
	GetClassrooms(context.Context, *GetClassroomsRequest) (*GetClassroomsResponse, error)
	CheckClassroomExists(context.Context, *CheckClassroomExistsRequest) (*CheckClassroomExistsResponse, error)
	GetLecturerClassroom(context.Context, *GetLecturerClassroomRequest) (*GetLecturerClassroomResponse, error)
	mustEmbedUnimplementedClassroomServiceServer()
}

// UnimplementedClassroomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClassroomServiceServer struct {
}

func (UnimplementedClassroomServiceServer) CreateClassroom(context.Context, *CreateClassroomRequest) (*CreateClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClassroom not implemented")
}
func (UnimplementedClassroomServiceServer) GetClassroom(context.Context, *GetClassroomRequest) (*GetClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassroom not implemented")
}
func (UnimplementedClassroomServiceServer) UpdateClassroom(context.Context, *UpdateClassroomRequest) (*UpdateClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClassroom not implemented")
}
func (UnimplementedClassroomServiceServer) DeleteClassroom(context.Context, *DeleteClassroomRequest) (*DeleteClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClassroom not implemented")
}
func (UnimplementedClassroomServiceServer) GetClassrooms(context.Context, *GetClassroomsRequest) (*GetClassroomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassrooms not implemented")
}
func (UnimplementedClassroomServiceServer) CheckClassroomExists(context.Context, *CheckClassroomExistsRequest) (*CheckClassroomExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckClassroomExists not implemented")
}
func (UnimplementedClassroomServiceServer) GetLecturerClassroom(context.Context, *GetLecturerClassroomRequest) (*GetLecturerClassroomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLecturerClassroom not implemented")
}
func (UnimplementedClassroomServiceServer) mustEmbedUnimplementedClassroomServiceServer() {}

// UnsafeClassroomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassroomServiceServer will
// result in compilation errors.
type UnsafeClassroomServiceServer interface {
	mustEmbedUnimplementedClassroomServiceServer()
}

func RegisterClassroomServiceServer(s grpc.ServiceRegistrar, srv ClassroomServiceServer) {
	s.RegisterService(&ClassroomService_ServiceDesc, srv)
}

func _ClassroomService_CreateClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).CreateClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_CreateClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).CreateClassroom(ctx, req.(*CreateClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_GetClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).GetClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_GetClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).GetClassroom(ctx, req.(*GetClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_UpdateClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).UpdateClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_UpdateClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).UpdateClassroom(ctx, req.(*UpdateClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_DeleteClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).DeleteClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_DeleteClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).DeleteClassroom(ctx, req.(*DeleteClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_GetClassrooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassroomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).GetClassrooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_GetClassrooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).GetClassrooms(ctx, req.(*GetClassroomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_CheckClassroomExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckClassroomExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).CheckClassroomExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_CheckClassroomExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).CheckClassroomExists(ctx, req.(*CheckClassroomExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassroomService_GetLecturerClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLecturerClassroomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassroomServiceServer).GetLecturerClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassroomService_GetLecturerClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassroomServiceServer).GetLecturerClassroom(ctx, req.(*GetLecturerClassroomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassroomService_ServiceDesc is the grpc.ServiceDesc for ClassroomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassroomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "classroom.v1.ClassroomService",
	HandlerType: (*ClassroomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClassroom",
			Handler:    _ClassroomService_CreateClassroom_Handler,
		},
		{
			MethodName: "GetClassroom",
			Handler:    _ClassroomService_GetClassroom_Handler,
		},
		{
			MethodName: "UpdateClassroom",
			Handler:    _ClassroomService_UpdateClassroom_Handler,
		},
		{
			MethodName: "DeleteClassroom",
			Handler:    _ClassroomService_DeleteClassroom_Handler,
		},
		{
			MethodName: "GetClassrooms",
			Handler:    _ClassroomService_GetClassrooms_Handler,
		},
		{
			MethodName: "CheckClassroomExists",
			Handler:    _ClassroomService_CheckClassroomExists_Handler,
		},
		{
			MethodName: "GetLecturerClassroom",
			Handler:    _ClassroomService_GetLecturerClassroom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "classroom.proto",
}
