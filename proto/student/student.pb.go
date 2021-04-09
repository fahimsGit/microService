// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/student/student.proto

package student

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Roll    string `protobuf:"bytes,3,opt,name=Roll,proto3" json:"Roll,omitempty"`
	Session string `protobuf:"bytes,4,opt,name=Session,proto3" json:"Session,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Student) GetRoll() string {
	if x != nil {
		return x.Roll
	}
	return ""
}

func (x *Student) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RequestCreateStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student *Student `protobuf:"bytes,1,opt,name=Student,proto3" json:"Student,omitempty"`
}

func (x *RequestCreateStudent) Reset() {
	*x = RequestCreateStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreateStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreateStudent) ProtoMessage() {}

func (x *RequestCreateStudent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreateStudent.ProtoReflect.Descriptor instead.
func (*RequestCreateStudent) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{2}
}

func (x *RequestCreateStudent) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type ResponseCreateStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Id     string  `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseCreateStudent) Reset() {
	*x = ResponseCreateStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCreateStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCreateStudent) ProtoMessage() {}

func (x *ResponseCreateStudent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCreateStudent.ProtoReflect.Descriptor instead.
func (*ResponseCreateStudent) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseCreateStudent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseCreateStudent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseCreateStudent) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ResponseGetAllStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
	Status   *Status    `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseGetAllStudent) Reset() {
	*x = ResponseGetAllStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetAllStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetAllStudent) ProtoMessage() {}

func (x *ResponseGetAllStudent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetAllStudent.ProtoReflect.Descriptor instead.
func (*ResponseGetAllStudent) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseGetAllStudent) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

func (x *ResponseGetAllStudent) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RequestCreateCourseEnrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	CourseId  string `protobuf:"bytes,2,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
}

func (x *RequestCreateCourseEnrollment) Reset() {
	*x = RequestCreateCourseEnrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreateCourseEnrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreateCourseEnrollment) ProtoMessage() {}

func (x *RequestCreateCourseEnrollment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreateCourseEnrollment.ProtoReflect.Descriptor instead.
func (*RequestCreateCourseEnrollment) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{5}
}

func (x *RequestCreateCourseEnrollment) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *RequestCreateCourseEnrollment) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type ResponseCreateCourseEnrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CourseId   string  `protobuf:"bytes,3,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
	CourseName string  `protobuf:"bytes,4,opt,name=CourseName,proto3" json:"CourseName,omitempty"`
	Status     *Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseCreateCourseEnrollment) Reset() {
	*x = ResponseCreateCourseEnrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCreateCourseEnrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCreateCourseEnrollment) ProtoMessage() {}

func (x *ResponseCreateCourseEnrollment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCreateCourseEnrollment.ProtoReflect.Descriptor instead.
func (*ResponseCreateCourseEnrollment) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseCreateCourseEnrollment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseCreateCourseEnrollment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseCreateCourseEnrollment) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *ResponseCreateCourseEnrollment) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *ResponseCreateCourseEnrollment) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ResponseGetAllEnrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CourseId   []string `protobuf:"bytes,3,rep,name=CourseId,proto3" json:"CourseId,omitempty"`
	CourseName []string `protobuf:"bytes,4,rep,name=CourseName,proto3" json:"CourseName,omitempty"`
	Status     *Status  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseGetAllEnrollment) Reset() {
	*x = ResponseGetAllEnrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGetAllEnrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGetAllEnrollment) ProtoMessage() {}

func (x *ResponseGetAllEnrollment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGetAllEnrollment.ProtoReflect.Descriptor instead.
func (*ResponseGetAllEnrollment) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseGetAllEnrollment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseGetAllEnrollment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseGetAllEnrollment) GetCourseId() []string {
	if x != nil {
		return x.CourseId
	}
	return nil
}

func (x *ResponseGetAllEnrollment) GetCourseName() []string {
	if x != nil {
		return x.CourseName
	}
	return nil
}

func (x *ResponseGetAllEnrollment) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RequestGetAllEnrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RequestGetAllEnrollment) Reset() {
	*x = RequestGetAllEnrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_student_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetAllEnrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetAllEnrollment) ProtoMessage() {}

func (x *RequestGetAllEnrollment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_student_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetAllEnrollment.ProtoReflect.Descriptor instead.
func (*RequestGetAllEnrollment) Descriptor() ([]byte, []int) {
	return file_proto_student_student_proto_rawDescGZIP(), []int{8}
}

func (x *RequestGetAllEnrollment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_student_student_proto protoreflect.FileDescriptor

var file_proto_student_student_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x38, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x64,
	0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x6e, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x59, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22,
	0xa9, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x18,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x29, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x32, 0x94, 0x02, 0x0a,
	0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x69, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_student_student_proto_rawDescOnce sync.Once
	file_proto_student_student_proto_rawDescData = file_proto_student_student_proto_rawDesc
)

func file_proto_student_student_proto_rawDescGZIP() []byte {
	file_proto_student_student_proto_rawDescOnce.Do(func() {
		file_proto_student_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_student_student_proto_rawDescData)
	})
	return file_proto_student_student_proto_rawDescData
}

var file_proto_student_student_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_student_student_proto_goTypes = []interface{}{
	(*Student)(nil),                        // 0: student.Student
	(*Status)(nil),                         // 1: student.Status
	(*RequestCreateStudent)(nil),           // 2: student.RequestCreateStudent
	(*ResponseCreateStudent)(nil),          // 3: student.ResponseCreateStudent
	(*ResponseGetAllStudent)(nil),          // 4: student.ResponseGetAllStudent
	(*RequestCreateCourseEnrollment)(nil),  // 5: student.RequestCreateCourseEnrollment
	(*ResponseCreateCourseEnrollment)(nil), // 6: student.ResponseCreateCourseEnrollment
	(*ResponseGetAllEnrollment)(nil),       // 7: student.ResponseGetAllEnrollment
	(*RequestGetAllEnrollment)(nil),        // 8: student.RequestGetAllEnrollment
	(*emptypb.Empty)(nil),                  // 9: google.protobuf.Empty
}
var file_proto_student_student_proto_depIdxs = []int32{
	0, // 0: student.RequestCreateStudent.Student:type_name -> student.Student
	1, // 1: student.ResponseCreateStudent.status:type_name -> student.Status
	0, // 2: student.ResponseGetAllStudent.students:type_name -> student.Student
	1, // 3: student.ResponseGetAllStudent.status:type_name -> student.Status
	1, // 4: student.ResponseCreateCourseEnrollment.status:type_name -> student.Status
	1, // 5: student.ResponseGetAllEnrollment.status:type_name -> student.Status
	2, // 6: student.StudentService.CreateStudent:input_type -> student.RequestCreateStudent
	9, // 7: student.StudentService.GetAllStudent:input_type -> google.protobuf.Empty
	5, // 8: student.StudentService.CreateCourseEnrollment:input_type -> student.RequestCreateCourseEnrollment
	3, // 9: student.StudentService.CreateStudent:output_type -> student.ResponseCreateStudent
	4, // 10: student.StudentService.GetAllStudent:output_type -> student.ResponseGetAllStudent
	6, // 11: student.StudentService.CreateCourseEnrollment:output_type -> student.ResponseCreateCourseEnrollment
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_student_student_proto_init() }
func file_proto_student_student_proto_init() {
	if File_proto_student_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_student_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_proto_student_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_student_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreateStudent); i {
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
		file_proto_student_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCreateStudent); i {
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
		file_proto_student_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetAllStudent); i {
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
		file_proto_student_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreateCourseEnrollment); i {
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
		file_proto_student_student_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCreateCourseEnrollment); i {
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
		file_proto_student_student_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGetAllEnrollment); i {
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
		file_proto_student_student_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetAllEnrollment); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_student_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_student_student_proto_goTypes,
		DependencyIndexes: file_proto_student_student_proto_depIdxs,
		MessageInfos:      file_proto_student_student_proto_msgTypes,
	}.Build()
	File_proto_student_student_proto = out.File
	file_proto_student_student_proto_rawDesc = nil
	file_proto_student_student_proto_goTypes = nil
	file_proto_student_student_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *RequestCreateStudent, opts ...grpc.CallOption) (*ResponseCreateStudent, error)
	GetAllStudent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseGetAllStudent, error)
	CreateCourseEnrollment(ctx context.Context, in *RequestCreateCourseEnrollment, opts ...grpc.CallOption) (*ResponseCreateCourseEnrollment, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *RequestCreateStudent, opts ...grpc.CallOption) (*ResponseCreateStudent, error) {
	out := new(ResponseCreateStudent)
	err := c.cc.Invoke(ctx, "/student.StudentService/CreateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetAllStudent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseGetAllStudent, error) {
	out := new(ResponseGetAllStudent)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetAllStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) CreateCourseEnrollment(ctx context.Context, in *RequestCreateCourseEnrollment, opts ...grpc.CallOption) (*ResponseCreateCourseEnrollment, error) {
	out := new(ResponseCreateCourseEnrollment)
	err := c.cc.Invoke(ctx, "/student.StudentService/CreateCourseEnrollment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	CreateStudent(context.Context, *RequestCreateStudent) (*ResponseCreateStudent, error)
	GetAllStudent(context.Context, *emptypb.Empty) (*ResponseGetAllStudent, error)
	CreateCourseEnrollment(context.Context, *RequestCreateCourseEnrollment) (*ResponseCreateCourseEnrollment, error)
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) CreateStudent(context.Context, *RequestCreateStudent) (*ResponseCreateStudent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (*UnimplementedStudentServiceServer) GetAllStudent(context.Context, *emptypb.Empty) (*ResponseGetAllStudent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStudent not implemented")
}
func (*UnimplementedStudentServiceServer) CreateCourseEnrollment(context.Context, *RequestCreateCourseEnrollment) (*ResponseCreateCourseEnrollment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourseEnrollment not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateStudent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/CreateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*RequestCreateStudent))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetAllStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetAllStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetAllStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetAllStudent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_CreateCourseEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateCourseEnrollment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateCourseEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/CreateCourseEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateCourseEnrollment(ctx, req.(*RequestCreateCourseEnrollment))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "GetAllStudent",
			Handler:    _StudentService_GetAllStudent_Handler,
		},
		{
			MethodName: "CreateCourseEnrollment",
			Handler:    _StudentService_CreateCourseEnrollment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/student/student.proto",
}
