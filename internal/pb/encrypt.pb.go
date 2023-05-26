// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: encrypt.proto

package pb

import (
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

type EncryptContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EncryptContentRequest) Reset() {
	*x = EncryptContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encrypt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptContentRequest) ProtoMessage() {}

func (x *EncryptContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encrypt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptContentRequest.ProtoReflect.Descriptor instead.
func (*EncryptContentRequest) Descriptor() ([]byte, []int) {
	return file_encrypt_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptContentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type EncryptContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EncryptContentResponse) Reset() {
	*x = EncryptContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encrypt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptContentResponse) ProtoMessage() {}

func (x *EncryptContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encrypt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptContentResponse.ProtoReflect.Descriptor instead.
func (*EncryptContentResponse) Descriptor() ([]byte, []int) {
	return file_encrypt_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptContentResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type EncryptFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *EncryptFileRequest) Reset() {
	*x = EncryptFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encrypt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptFileRequest) ProtoMessage() {}

func (x *EncryptFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encrypt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptFileRequest.ProtoReflect.Descriptor instead.
func (*EncryptFileRequest) Descriptor() ([]byte, []int) {
	return file_encrypt_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type EncryptFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EncryptFileResponse) Reset() {
	*x = EncryptFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encrypt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptFileResponse) ProtoMessage() {}

func (x *EncryptFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encrypt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptFileResponse.ProtoReflect.Descriptor instead.
func (*EncryptFileResponse) Descriptor() ([]byte, []int) {
	return file_encrypt_proto_rawDescGZIP(), []int{3}
}

var File_encrypt_proto protoreflect.FileDescriptor

var file_encrypt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x15, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a,
	0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x86,
	0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7a, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2d, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x78, 0x0a,
	0x0b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12,
	0x26, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2d, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x2f, 0x7b, 0x70, 0x61, 0x74, 0x68, 0x7d, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encrypt_proto_rawDescOnce sync.Once
	file_encrypt_proto_rawDescData = file_encrypt_proto_rawDesc
)

func file_encrypt_proto_rawDescGZIP() []byte {
	file_encrypt_proto_rawDescOnce.Do(func() {
		file_encrypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_encrypt_proto_rawDescData)
	})
	return file_encrypt_proto_rawDescData
}

var file_encrypt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_encrypt_proto_goTypes = []interface{}{
	(*EncryptContentRequest)(nil),  // 0: encrypt.EncryptContentRequest
	(*EncryptContentResponse)(nil), // 1: encrypt.EncryptContentResponse
	(*EncryptFileRequest)(nil),     // 2: encrypt.EncryptFileRequest
	(*EncryptFileResponse)(nil),    // 3: encrypt.EncryptFileResponse
}
var file_encrypt_proto_depIdxs = []int32{
	0, // 0: encrypt.EncryptService.EncryptContent:input_type -> encrypt.EncryptContentRequest
	2, // 1: encrypt.EncryptService.EncryptFile:input_type -> encrypt.EncryptFileRequest
	1, // 2: encrypt.EncryptService.EncryptContent:output_type -> encrypt.EncryptContentResponse
	3, // 3: encrypt.EncryptService.EncryptFile:output_type -> encrypt.EncryptFileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_encrypt_proto_init() }
func file_encrypt_proto_init() {
	if File_encrypt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encrypt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptContentRequest); i {
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
		file_encrypt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptContentResponse); i {
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
		file_encrypt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptFileRequest); i {
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
		file_encrypt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptFileResponse); i {
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
			RawDescriptor: file_encrypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encrypt_proto_goTypes,
		DependencyIndexes: file_encrypt_proto_depIdxs,
		MessageInfos:      file_encrypt_proto_msgTypes,
	}.Build()
	File_encrypt_proto = out.File
	file_encrypt_proto_rawDesc = nil
	file_encrypt_proto_goTypes = nil
	file_encrypt_proto_depIdxs = nil
}
