// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: api/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginIn) Reset() {
	*x = LoginIn{}
	mi := &file_api_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginIn) ProtoMessage() {}

func (x *LoginIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginIn.ProtoReflect.Descriptor instead.
func (*LoginIn) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginIn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginOut) Reset() {
	*x = LoginOut{}
	mi := &file_api_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginOut) ProtoMessage() {}

func (x *LoginOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginOut.ProtoReflect.Descriptor instead.
func (*LoginOut) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginOut) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignupIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignupIn) Reset() {
	*x = SignupIn{}
	mi := &file_api_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupIn) ProtoMessage() {}

func (x *SignupIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupIn.ProtoReflect.Descriptor instead.
func (*SignupIn) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignupIn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignupIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignupOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignupOut) Reset() {
	*x = SignupOut{}
	mi := &file_api_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignupOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupOut) ProtoMessage() {}

func (x *SignupOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupOut.ProtoReflect.Descriptor instead.
func (*SignupOut) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignupOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ChangePasswordIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OldPassword   string                 `protobuf:"bytes,1,opt,name=OldPassword,proto3" json:"OldPassword,omitempty"`
	NewPassword   string                 `protobuf:"bytes,2,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordIn) Reset() {
	*x = ChangePasswordIn{}
	mi := &file_api_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordIn) ProtoMessage() {}

func (x *ChangePasswordIn) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordIn.ProtoReflect.Descriptor instead.
func (*ChangePasswordIn) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordIn) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordIn) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ChangePasswordOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordOut) Reset() {
	*x = ChangePasswordOut{}
	mi := &file_api_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordOut) ProtoMessage() {}

func (x *ChangePasswordOut) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordOut.ProtoReflect.Descriptor instead.
func (*ChangePasswordOut) Descriptor() ([]byte, []int) {
	return file_api_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_auth_proto protoreflect.FileDescriptor

var file_api_auth_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x56, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x08, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x1a, 0x09, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x12, 0x09, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49, 0x6e, 0x1a, 0x0a, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x11, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x1a,
	0x12, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_auth_proto_rawDescOnce sync.Once
	file_api_auth_proto_rawDescData []byte
)

func file_api_auth_proto_rawDescGZIP() []byte {
	file_api_auth_proto_rawDescOnce.Do(func() {
		file_api_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_auth_proto_rawDesc), len(file_api_auth_proto_rawDesc)))
	})
	return file_api_auth_proto_rawDescData
}

var file_api_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_auth_proto_goTypes = []any{
	(*LoginIn)(nil),           // 0: LoginIn
	(*LoginOut)(nil),          // 1: LoginOut
	(*SignupIn)(nil),          // 2: SignupIn
	(*SignupOut)(nil),         // 3: SignupOut
	(*ChangePasswordIn)(nil),  // 4: ChangePasswordIn
	(*ChangePasswordOut)(nil), // 5: ChangePasswordOut
}
var file_api_auth_proto_depIdxs = []int32{
	0, // 0: AuthService.Login:input_type -> LoginIn
	2, // 1: AuthService.Signup:input_type -> SignupIn
	4, // 2: AuthService.ChangePassword:input_type -> ChangePasswordIn
	1, // 3: AuthService.Login:output_type -> LoginOut
	3, // 4: AuthService.Signup:output_type -> SignupOut
	5, // 5: AuthService.ChangePassword:output_type -> ChangePasswordOut
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_auth_proto_init() }
func file_api_auth_proto_init() {
	if File_api_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_auth_proto_rawDesc), len(file_api_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_auth_proto_goTypes,
		DependencyIndexes: file_api_auth_proto_depIdxs,
		MessageInfos:      file_api_auth_proto_msgTypes,
	}.Build()
	File_api_auth_proto = out.File
	file_api_auth_proto_goTypes = nil
	file_api_auth_proto_depIdxs = nil
}
