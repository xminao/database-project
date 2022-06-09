// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: university.proto

package university

import (
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

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId       string `protobuf:"bytes,1,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	StudentName     string `protobuf:"bytes,2,opt,name=StudentName,proto3" json:"StudentName,omitempty"`
	College         string `protobuf:"bytes,3,opt,name=College,proto3" json:"College,omitempty"`
	Year            string `protobuf:"bytes,4,opt,name=Year,proto3" json:"Year,omitempty"`
	Major           string `protobuf:"bytes,5,opt,name=Major,proto3" json:"Major,omitempty"`
	Class           string `protobuf:"bytes,6,opt,name=Class,proto3" json:"Class,omitempty"`
	IsBinding       bool   `protobuf:"varint,7,opt,name=IsBinding,proto3" json:"IsBinding,omitempty"`
	BindingUsername string `protobuf:"bytes,8,opt,name=BindingUsername,proto3" json:"BindingUsername,omitempty"`
	CreatedAt       int64  `protobuf:"varint,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_university_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_university_proto_msgTypes[0]
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
	return file_university_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Student) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *Student) GetCollege() string {
	if x != nil {
		return x.College
	}
	return ""
}

func (x *Student) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Student) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *Student) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Student) GetIsBinding() bool {
	if x != nil {
		return x.IsBinding
	}
	return false
}

func (x *Student) GetBindingUsername() string {
	if x != nil {
		return x.BindingUsername
	}
	return ""
}

func (x *Student) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetStudentInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
}

func (x *GetStudentInfoReq) Reset() {
	*x = GetStudentInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_university_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentInfoReq) ProtoMessage() {}

func (x *GetStudentInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_university_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentInfoReq.ProtoReflect.Descriptor instead.
func (*GetStudentInfoReq) Descriptor() ([]byte, []int) {
	return file_university_proto_rawDescGZIP(), []int{1}
}

func (x *GetStudentInfoReq) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type GetStudentInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentInfo *Student `protobuf:"bytes,1,opt,name=studentInfo,proto3" json:"studentInfo,omitempty"`
}

func (x *GetStudentInfoResp) Reset() {
	*x = GetStudentInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_university_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentInfoResp) ProtoMessage() {}

func (x *GetStudentInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_university_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentInfoResp.ProtoReflect.Descriptor instead.
func (*GetStudentInfoResp) Descriptor() ([]byte, []int) {
	return file_university_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentInfoResp) GetStudentInfo() *Student {
	if x != nil {
		return x.StudentInfo
	}
	return nil
}

var File_university_proto protoreflect.FileDescriptor

var file_university_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x22, 0x89,
	0x02, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x60, 0x0a, 0x0d, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x52, 0x70, 0x63, 0x12, 0x4f, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e,
	0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x75,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_university_proto_rawDescOnce sync.Once
	file_university_proto_rawDescData = file_university_proto_rawDesc
)

func file_university_proto_rawDescGZIP() []byte {
	file_university_proto_rawDescOnce.Do(func() {
		file_university_proto_rawDescData = protoimpl.X.CompressGZIP(file_university_proto_rawDescData)
	})
	return file_university_proto_rawDescData
}

var file_university_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_university_proto_goTypes = []interface{}{
	(*Student)(nil),            // 0: university.Student
	(*GetStudentInfoReq)(nil),  // 1: university.GetStudentInfoReq
	(*GetStudentInfoResp)(nil), // 2: university.GetStudentInfoResp
}
var file_university_proto_depIdxs = []int32{
	0, // 0: university.GetStudentInfoResp.studentInfo:type_name -> university.Student
	1, // 1: university.UniversityRpc.GetStudentInfo:input_type -> university.GetStudentInfoReq
	2, // 2: university.UniversityRpc.GetStudentInfo:output_type -> university.GetStudentInfoResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_university_proto_init() }
func file_university_proto_init() {
	if File_university_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_university_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_university_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentInfoReq); i {
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
		file_university_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentInfoResp); i {
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
			RawDescriptor: file_university_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_university_proto_goTypes,
		DependencyIndexes: file_university_proto_depIdxs,
		MessageInfos:      file_university_proto_msgTypes,
	}.Build()
	File_university_proto = out.File
	file_university_proto_rawDesc = nil
	file_university_proto_goTypes = nil
	file_university_proto_depIdxs = nil
}
