// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: submission.proto

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
	SubmissionService_CreateSubmission_FullMethodName            = "/submission.v1.SubmissionService/CreateSubmission"
	SubmissionService_GetSubmission_FullMethodName               = "/submission.v1.SubmissionService/GetSubmission"
	SubmissionService_UpdateSubmission_FullMethodName            = "/submission.v1.SubmissionService/UpdateSubmission"
	SubmissionService_DeleteSubmission_FullMethodName            = "/submission.v1.SubmissionService/DeleteSubmission"
	SubmissionService_GetAllSubmissionsOfExercise_FullMethodName = "/submission.v1.SubmissionService/GetAllSubmissionsOfExercise"
	SubmissionService_GetSubmissionOfUser_FullMethodName         = "/submission.v1.SubmissionService/GetSubmissionOfUser"
	SubmissionService_GetSubmissionFromUser_FullMethodName       = "/submission.v1.SubmissionService/GetSubmissionFromUser"
)

// SubmissionServiceClient is the client API for SubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionServiceClient interface {
	CreateSubmission(ctx context.Context, in *CreateSubmissionRequest, opts ...grpc.CallOption) (*CreateSubmissionResponse, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error)
	UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*UpdateSubmissionResponse, error)
	DeleteSubmission(ctx context.Context, in *DeleteSubmissionRequest, opts ...grpc.CallOption) (*DeleteSubmissionResponse, error)
	GetAllSubmissionsOfExercise(ctx context.Context, in *GetAllSubmissionsOfExerciseRequest, opts ...grpc.CallOption) (*GetAllSubmissionsOfExerciseResponse, error)
	GetSubmissionOfUser(ctx context.Context, in *GetSubmissionOfUserRequest, opts ...grpc.CallOption) (*GetSubmissionOfUserResponse, error)
	GetSubmissionFromUser(ctx context.Context, in *GetSubmissionFromUserRequest, opts ...grpc.CallOption) (*GetSubmissionFromUserResponse, error)
}

type submissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionServiceClient(cc grpc.ClientConnInterface) SubmissionServiceClient {
	return &submissionServiceClient{cc}
}

func (c *submissionServiceClient) CreateSubmission(ctx context.Context, in *CreateSubmissionRequest, opts ...grpc.CallOption) (*CreateSubmissionResponse, error) {
	out := new(CreateSubmissionResponse)
	err := c.cc.Invoke(ctx, SubmissionService_CreateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error) {
	out := new(GetSubmissionResponse)
	err := c.cc.Invoke(ctx, SubmissionService_GetSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*UpdateSubmissionResponse, error) {
	out := new(UpdateSubmissionResponse)
	err := c.cc.Invoke(ctx, SubmissionService_UpdateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) DeleteSubmission(ctx context.Context, in *DeleteSubmissionRequest, opts ...grpc.CallOption) (*DeleteSubmissionResponse, error) {
	out := new(DeleteSubmissionResponse)
	err := c.cc.Invoke(ctx, SubmissionService_DeleteSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) GetAllSubmissionsOfExercise(ctx context.Context, in *GetAllSubmissionsOfExerciseRequest, opts ...grpc.CallOption) (*GetAllSubmissionsOfExerciseResponse, error) {
	out := new(GetAllSubmissionsOfExerciseResponse)
	err := c.cc.Invoke(ctx, SubmissionService_GetAllSubmissionsOfExercise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) GetSubmissionOfUser(ctx context.Context, in *GetSubmissionOfUserRequest, opts ...grpc.CallOption) (*GetSubmissionOfUserResponse, error) {
	out := new(GetSubmissionOfUserResponse)
	err := c.cc.Invoke(ctx, SubmissionService_GetSubmissionOfUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionServiceClient) GetSubmissionFromUser(ctx context.Context, in *GetSubmissionFromUserRequest, opts ...grpc.CallOption) (*GetSubmissionFromUserResponse, error) {
	out := new(GetSubmissionFromUserResponse)
	err := c.cc.Invoke(ctx, SubmissionService_GetSubmissionFromUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServiceServer is the server API for SubmissionService service.
// All implementations must embed UnimplementedSubmissionServiceServer
// for forward compatibility
type SubmissionServiceServer interface {
	CreateSubmission(context.Context, *CreateSubmissionRequest) (*CreateSubmissionResponse, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error)
	UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*UpdateSubmissionResponse, error)
	DeleteSubmission(context.Context, *DeleteSubmissionRequest) (*DeleteSubmissionResponse, error)
	GetAllSubmissionsOfExercise(context.Context, *GetAllSubmissionsOfExerciseRequest) (*GetAllSubmissionsOfExerciseResponse, error)
	GetSubmissionOfUser(context.Context, *GetSubmissionOfUserRequest) (*GetSubmissionOfUserResponse, error)
	GetSubmissionFromUser(context.Context, *GetSubmissionFromUserRequest) (*GetSubmissionFromUserResponse, error)
	mustEmbedUnimplementedSubmissionServiceServer()
}

// UnimplementedSubmissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionServiceServer struct {
}

func (UnimplementedSubmissionServiceServer) CreateSubmission(context.Context, *CreateSubmissionRequest) (*CreateSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*UpdateSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) DeleteSubmission(context.Context, *DeleteSubmissionRequest) (*DeleteSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubmission not implemented")
}
func (UnimplementedSubmissionServiceServer) GetAllSubmissionsOfExercise(context.Context, *GetAllSubmissionsOfExerciseRequest) (*GetAllSubmissionsOfExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSubmissionsOfExercise not implemented")
}
func (UnimplementedSubmissionServiceServer) GetSubmissionOfUser(context.Context, *GetSubmissionOfUserRequest) (*GetSubmissionOfUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmissionOfUser not implemented")
}
func (UnimplementedSubmissionServiceServer) GetSubmissionFromUser(context.Context, *GetSubmissionFromUserRequest) (*GetSubmissionFromUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmissionFromUser not implemented")
}
func (UnimplementedSubmissionServiceServer) mustEmbedUnimplementedSubmissionServiceServer() {}

// UnsafeSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServiceServer will
// result in compilation errors.
type UnsafeSubmissionServiceServer interface {
	mustEmbedUnimplementedSubmissionServiceServer()
}

func RegisterSubmissionServiceServer(s grpc.ServiceRegistrar, srv SubmissionServiceServer) {
	s.RegisterService(&SubmissionService_ServiceDesc, srv)
}

func _SubmissionService_CreateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).CreateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_CreateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).CreateSubmission(ctx, req.(*CreateSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_GetSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_UpdateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).UpdateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_UpdateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).UpdateSubmission(ctx, req.(*UpdateSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_DeleteSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).DeleteSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_DeleteSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).DeleteSubmission(ctx, req.(*DeleteSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_GetAllSubmissionsOfExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSubmissionsOfExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).GetAllSubmissionsOfExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_GetAllSubmissionsOfExercise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).GetAllSubmissionsOfExercise(ctx, req.(*GetAllSubmissionsOfExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_GetSubmissionOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).GetSubmissionOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_GetSubmissionOfUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).GetSubmissionOfUser(ctx, req.(*GetSubmissionOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmissionService_GetSubmissionFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionFromUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServiceServer).GetSubmissionFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmissionService_GetSubmissionFromUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServiceServer).GetSubmissionFromUser(ctx, req.(*GetSubmissionFromUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmissionService_ServiceDesc is the grpc.ServiceDesc for SubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submission.v1.SubmissionService",
	HandlerType: (*SubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubmission",
			Handler:    _SubmissionService_CreateSubmission_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _SubmissionService_GetSubmission_Handler,
		},
		{
			MethodName: "UpdateSubmission",
			Handler:    _SubmissionService_UpdateSubmission_Handler,
		},
		{
			MethodName: "DeleteSubmission",
			Handler:    _SubmissionService_DeleteSubmission_Handler,
		},
		{
			MethodName: "GetAllSubmissionsOfExercise",
			Handler:    _SubmissionService_GetAllSubmissionsOfExercise_Handler,
		},
		{
			MethodName: "GetSubmissionOfUser",
			Handler:    _SubmissionService_GetSubmissionOfUser_Handler,
		},
		{
			MethodName: "GetSubmissionFromUser",
			Handler:    _SubmissionService_GetSubmissionFromUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "submission.proto",
}