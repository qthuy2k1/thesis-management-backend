// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: api_point.proto

package v1

import (
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Class    string  `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Major    *string `protobuf:"bytes,3,opt,name=major,proto3,oneof" json:"major,omitempty"`
	Phone    *string `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	PhotoSrc string  `protobuf:"bytes,5,opt,name=photoSrc,proto3" json:"photoSrc,omitempty"`
	Role     string  `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	Name     string  `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Email    string  `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserScheduleResponse) Reset() {
	*x = UserScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScheduleResponse) ProtoMessage() {}

func (x *UserScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScheduleResponse.ProtoReflect.Descriptor instead.
func (*UserScheduleResponse) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{0}
}

func (x *UserScheduleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserScheduleResponse) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *UserScheduleResponse) GetMajor() string {
	if x != nil && x.Major != nil {
		return *x.Major
	}
	return ""
}

func (x *UserScheduleResponse) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UserScheduleResponse) GetPhotoSrc() string {
	if x != nil {
		return x.PhotoSrc
	}
	return ""
}

func (x *UserScheduleResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserScheduleResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserScheduleResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// ===========================
// POINT
type AssessItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lecturer *UserScheduleResponse `protobuf:"bytes,2,opt,name=lecturer,proto3" json:"lecturer,omitempty"`
	Point    int64                 `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
	Comment  string                `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *AssessItemResponse) Reset() {
	*x = AssessItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssessItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssessItemResponse) ProtoMessage() {}

func (x *AssessItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssessItemResponse.ProtoReflect.Descriptor instead.
func (*AssessItemResponse) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{1}
}

func (x *AssessItemResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssessItemResponse) GetLecturer() *UserScheduleResponse {
	if x != nil {
		return x.Lecturer
	}
	return nil
}

func (x *AssessItemResponse) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *AssessItemResponse) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type AssessItemInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LecturerID string `protobuf:"bytes,2,opt,name=lecturerID,proto3" json:"lecturerID,omitempty"`
	Point      int64  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
	Comment    string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *AssessItemInput) Reset() {
	*x = AssessItemInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssessItemInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssessItemInput) ProtoMessage() {}

func (x *AssessItemInput) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssessItemInput.ProtoReflect.Descriptor instead.
func (*AssessItemInput) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{2}
}

func (x *AssessItemInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssessItemInput) GetLecturerID() string {
	if x != nil {
		return x.LecturerID
	}
	return ""
}

func (x *AssessItemInput) GetPoint() int64 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *AssessItemInput) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type PointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Student  *UserScheduleResponse `protobuf:"bytes,2,opt,name=student,proto3" json:"student,omitempty"`
	Assesses []*AssessItemResponse `protobuf:"bytes,3,rep,name=assesses,proto3" json:"assesses,omitempty"`
}

func (x *PointResponse) Reset() {
	*x = PointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointResponse) ProtoMessage() {}

func (x *PointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointResponse.ProtoReflect.Descriptor instead.
func (*PointResponse) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{3}
}

func (x *PointResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PointResponse) GetStudent() *UserScheduleResponse {
	if x != nil {
		return x.Student
	}
	return nil
}

func (x *PointResponse) GetAssesses() []*AssessItemResponse {
	if x != nil {
		return x.Assesses
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentID string             `protobuf:"bytes,2,opt,name=studentID,proto3" json:"studentID,omitempty"`
	Assesses  []*AssessItemInput `protobuf:"bytes,3,rep,name=assesses,proto3" json:"assesses,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{4}
}

func (x *Point) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Point) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *Point) GetAssesses() []*AssessItemInput {
	if x != nil {
		return x.Assesses
	}
	return nil
}

type CreateOrUpdatePointDefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *Point `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *CreateOrUpdatePointDefRequest) Reset() {
	*x = CreateOrUpdatePointDefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdatePointDefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdatePointDefRequest) ProtoMessage() {}

func (x *CreateOrUpdatePointDefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdatePointDefRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdatePointDefRequest) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOrUpdatePointDefRequest) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type CreateOrUpdatePointDefResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point   *PointResponse `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateOrUpdatePointDefResponse) Reset() {
	*x = CreateOrUpdatePointDefResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdatePointDefResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdatePointDefResponse) ProtoMessage() {}

func (x *CreateOrUpdatePointDefResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdatePointDefResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdatePointDefResponse) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOrUpdatePointDefResponse) GetPoint() *PointResponse {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *CreateOrUpdatePointDefResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllPointDefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllPointDefRequest) Reset() {
	*x = GetAllPointDefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPointDefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPointDefRequest) ProtoMessage() {}

func (x *GetAllPointDefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPointDefRequest.ProtoReflect.Descriptor instead.
func (*GetAllPointDefRequest) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{7}
}

type GetAllPointDefResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*PointResponse `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *GetAllPointDefResponse) Reset() {
	*x = GetAllPointDefResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_point_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPointDefResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPointDefResponse) ProtoMessage() {}

func (x *GetAllPointDefResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_point_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPointDefResponse.ProtoReflect.Descriptor instead.
func (*GetAllPointDefResponse) Descriptor() ([]byte, []int) {
	return file_api_point_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllPointDefResponse) GetPoints() []*PointResponse {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_api_point_proto protoreflect.FileDescriptor

var file_api_point_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0,
	0x01, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x19, 0x0a,
	0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x72, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x53, 0x72, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x0d,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a,
	0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x05, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x39, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x1d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x6d, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x32, 0x92,
	0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x8d, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x72, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x66, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x44, 0x65, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x74, 0x68, 0x75, 0x79, 0x32, 0x6b, 0x31, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x69,
	0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_point_proto_rawDescOnce sync.Once
	file_api_point_proto_rawDescData = file_api_point_proto_rawDesc
)

func file_api_point_proto_rawDescGZIP() []byte {
	file_api_point_proto_rawDescOnce.Do(func() {
		file_api_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_point_proto_rawDescData)
	})
	return file_api_point_proto_rawDescData
}

var file_api_point_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_point_proto_goTypes = []interface{}{
	(*UserScheduleResponse)(nil),           // 0: api.point.v1.UserScheduleResponse
	(*AssessItemResponse)(nil),             // 1: api.point.v1.AssessItemResponse
	(*AssessItemInput)(nil),                // 2: api.point.v1.AssessItemInput
	(*PointResponse)(nil),                  // 3: api.point.v1.PointResponse
	(*Point)(nil),                          // 4: api.point.v1.Point
	(*CreateOrUpdatePointDefRequest)(nil),  // 5: api.point.v1.CreateOrUpdatePointDefRequest
	(*CreateOrUpdatePointDefResponse)(nil), // 6: api.point.v1.CreateOrUpdatePointDefResponse
	(*GetAllPointDefRequest)(nil),          // 7: api.point.v1.GetAllPointDefRequest
	(*GetAllPointDefResponse)(nil),         // 8: api.point.v1.GetAllPointDefResponse
}
var file_api_point_proto_depIdxs = []int32{
	0, // 0: api.point.v1.AssessItemResponse.lecturer:type_name -> api.point.v1.UserScheduleResponse
	0, // 1: api.point.v1.PointResponse.student:type_name -> api.point.v1.UserScheduleResponse
	1, // 2: api.point.v1.PointResponse.assesses:type_name -> api.point.v1.AssessItemResponse
	2, // 3: api.point.v1.Point.assesses:type_name -> api.point.v1.AssessItemInput
	4, // 4: api.point.v1.CreateOrUpdatePointDefRequest.point:type_name -> api.point.v1.Point
	3, // 5: api.point.v1.CreateOrUpdatePointDefResponse.point:type_name -> api.point.v1.PointResponse
	3, // 6: api.point.v1.GetAllPointDefResponse.points:type_name -> api.point.v1.PointResponse
	5, // 7: api.point.v1.PointService.CreateOrUpdatePointDef:input_type -> api.point.v1.CreateOrUpdatePointDefRequest
	7, // 8: api.point.v1.PointService.GetAllPointDef:input_type -> api.point.v1.GetAllPointDefRequest
	6, // 9: api.point.v1.PointService.CreateOrUpdatePointDef:output_type -> api.point.v1.CreateOrUpdatePointDefResponse
	8, // 10: api.point.v1.PointService.GetAllPointDef:output_type -> api.point.v1.GetAllPointDefResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_point_proto_init() }
func file_api_point_proto_init() {
	if File_api_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScheduleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssessItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssessItemInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdatePointDefRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdatePointDefResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPointDefRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_point_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPointDefResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_point_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_point_proto_goTypes,
		DependencyIndexes: file_api_point_proto_depIdxs,
		MessageInfos:      file_api_point_proto_msgTypes,
	}.Build()
	File_api_point_proto = out.File
	file_api_point_proto_rawDesc = nil
	file_api_point_proto_goTypes = nil
	file_api_point_proto_depIdxs = nil
}
