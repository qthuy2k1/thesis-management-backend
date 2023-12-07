// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: schedule.proto

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
	ScheduleService_GetSchedules_FullMethodName           = "/schedule.v1.ScheduleService/GetSchedules"
	ScheduleService_CreateSchedule_FullMethodName         = "/schedule.v1.ScheduleService/CreateSchedule"
	ScheduleService_GetNotifications_FullMethodName       = "/schedule.v1.ScheduleService/GetNotifications"
	ScheduleService_CreateNotification_FullMethodName     = "/schedule.v1.ScheduleService/CreateNotification"
	ScheduleService_CreateOrUpdatePointDef_FullMethodName = "/schedule.v1.ScheduleService/CreateOrUpdatePointDef"
	ScheduleService_GetAllPointDefs_FullMethodName        = "/schedule.v1.ScheduleService/GetAllPointDefs"
	ScheduleService_UpdatePointDef_FullMethodName         = "/schedule.v1.ScheduleService/UpdatePointDef"
	ScheduleService_DeletePointDef_FullMethodName         = "/schedule.v1.ScheduleService/DeletePointDef"
)

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	GetSchedules(ctx context.Context, in *GetSchedulesRequest, opts ...grpc.CallOption) (*GetSchedulesResponse, error)
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	CreateOrUpdatePointDef(ctx context.Context, in *CreateOrUpdatePointDefRequest, opts ...grpc.CallOption) (*CreateOrUpdatePointDefResponse, error)
	GetAllPointDefs(ctx context.Context, in *GetAllPointDefsRequest, opts ...grpc.CallOption) (*GetAllPointDefsResponse, error)
	UpdatePointDef(ctx context.Context, in *UpdatePointDefRequest, opts ...grpc.CallOption) (*UpdatePointDefResponse, error)
	DeletePointDef(ctx context.Context, in *DeletePointDefRequest, opts ...grpc.CallOption) (*DeletePointDefResponse, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) GetSchedules(ctx context.Context, in *GetSchedulesRequest, opts ...grpc.CallOption) (*GetSchedulesResponse, error) {
	out := new(GetSchedulesResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetSchedules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleResponse, error) {
	out := new(CreateScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_CreateSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetNotifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, ScheduleService_CreateNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) CreateOrUpdatePointDef(ctx context.Context, in *CreateOrUpdatePointDefRequest, opts ...grpc.CallOption) (*CreateOrUpdatePointDefResponse, error) {
	out := new(CreateOrUpdatePointDefResponse)
	err := c.cc.Invoke(ctx, ScheduleService_CreateOrUpdatePointDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetAllPointDefs(ctx context.Context, in *GetAllPointDefsRequest, opts ...grpc.CallOption) (*GetAllPointDefsResponse, error) {
	out := new(GetAllPointDefsResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetAllPointDefs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) UpdatePointDef(ctx context.Context, in *UpdatePointDefRequest, opts ...grpc.CallOption) (*UpdatePointDefResponse, error) {
	out := new(UpdatePointDefResponse)
	err := c.cc.Invoke(ctx, ScheduleService_UpdatePointDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) DeletePointDef(ctx context.Context, in *DeletePointDefRequest, opts ...grpc.CallOption) (*DeletePointDefResponse, error) {
	out := new(DeletePointDefResponse)
	err := c.cc.Invoke(ctx, ScheduleService_DeletePointDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility
type ScheduleServiceServer interface {
	GetSchedules(context.Context, *GetSchedulesRequest) (*GetSchedulesResponse, error)
	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	CreateOrUpdatePointDef(context.Context, *CreateOrUpdatePointDefRequest) (*CreateOrUpdatePointDefResponse, error)
	GetAllPointDefs(context.Context, *GetAllPointDefsRequest) (*GetAllPointDefsResponse, error)
	UpdatePointDef(context.Context, *UpdatePointDefRequest) (*UpdatePointDefResponse, error)
	DeletePointDef(context.Context, *DeletePointDefRequest) (*DeletePointDefResponse, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (UnimplementedScheduleServiceServer) GetSchedules(context.Context, *GetSchedulesRequest) (*GetSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedules not implemented")
}
func (UnimplementedScheduleServiceServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedScheduleServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedScheduleServiceServer) CreateOrUpdatePointDef(context.Context, *CreateOrUpdatePointDefRequest) (*CreateOrUpdatePointDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdatePointDef not implemented")
}
func (UnimplementedScheduleServiceServer) GetAllPointDefs(context.Context, *GetAllPointDefsRequest) (*GetAllPointDefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPointDefs not implemented")
}
func (UnimplementedScheduleServiceServer) UpdatePointDef(context.Context, *UpdatePointDefRequest) (*UpdatePointDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePointDef not implemented")
}
func (UnimplementedScheduleServiceServer) DeletePointDef(context.Context, *DeletePointDefRequest) (*DeletePointDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePointDef not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_GetSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetSchedules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSchedules(ctx, req.(*GetSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_CreateSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_CreateOrUpdatePointDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdatePointDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateOrUpdatePointDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_CreateOrUpdatePointDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateOrUpdatePointDef(ctx, req.(*CreateOrUpdatePointDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetAllPointDefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPointDefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetAllPointDefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetAllPointDefs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetAllPointDefs(ctx, req.(*GetAllPointDefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_UpdatePointDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePointDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).UpdatePointDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_UpdatePointDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).UpdatePointDef(ctx, req.(*UpdatePointDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_DeletePointDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePointDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).DeletePointDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_DeletePointDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).DeletePointDef(ctx, req.(*DeletePointDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule.v1.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchedules",
			Handler:    _ScheduleService_GetSchedules_Handler,
		},
		{
			MethodName: "CreateSchedule",
			Handler:    _ScheduleService_CreateSchedule_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _ScheduleService_GetNotifications_Handler,
		},
		{
			MethodName: "CreateNotification",
			Handler:    _ScheduleService_CreateNotification_Handler,
		},
		{
			MethodName: "CreateOrUpdatePointDef",
			Handler:    _ScheduleService_CreateOrUpdatePointDef_Handler,
		},
		{
			MethodName: "GetAllPointDefs",
			Handler:    _ScheduleService_GetAllPointDefs_Handler,
		},
		{
			MethodName: "UpdatePointDef",
			Handler:    _ScheduleService_UpdatePointDef_Handler,
		},
		{
			MethodName: "DeletePointDef",
			Handler:    _ScheduleService_DeletePointDef_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule.proto",
}
