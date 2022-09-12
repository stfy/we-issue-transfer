// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: managed/api_version.proto

package managed

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrentContractApiVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CurrentContractApiVersion) Reset() {
	*x = CurrentContractApiVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managed_api_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentContractApiVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentContractApiVersion) ProtoMessage() {}

func (x *CurrentContractApiVersion) ProtoReflect() protoreflect.Message {
	mi := &file_managed_api_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentContractApiVersion.ProtoReflect.Descriptor instead.
func (*CurrentContractApiVersion) Descriptor() ([]byte, []int) {
	return file_managed_api_version_proto_rawDescGZIP(), []int{0}
}

var file_managed_api_version_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51234,
		Name:          "wavesenterprise.contract_api_version",
		Tag:           "bytes,51234,opt,name=contract_api_version",
		Filename:      "managed/api_version.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string contract_api_version = 51234;
	E_ContractApiVersion = &file_managed_api_version_proto_extTypes[0]
)

var File_managed_api_version_proto protoreflect.FileDescriptor

var file_managed_api_version_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e,
	0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x31, 0x92, 0x82, 0x19,
	0x2d, 0x31, 0x2e, 0x35, 0x2d, 0x61, 0x6d, 0x61, 0x7a, 0x75, 0x72, 0x5f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x3a, 0x53,
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x17, 0x5a, 0x15, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managed_api_version_proto_rawDescOnce sync.Once
	file_managed_api_version_proto_rawDescData = file_managed_api_version_proto_rawDesc
)

func file_managed_api_version_proto_rawDescGZIP() []byte {
	file_managed_api_version_proto_rawDescOnce.Do(func() {
		file_managed_api_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_managed_api_version_proto_rawDescData)
	})
	return file_managed_api_version_proto_rawDescData
}

var file_managed_api_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_managed_api_version_proto_goTypes = []interface{}{
	(*CurrentContractApiVersion)(nil),   // 0: wavesenterprise.CurrentContractApiVersion
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_managed_api_version_proto_depIdxs = []int32{
	1, // 0: wavesenterprise.contract_api_version:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managed_api_version_proto_init() }
func file_managed_api_version_proto_init() {
	if File_managed_api_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managed_api_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentContractApiVersion); i {
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
			RawDescriptor: file_managed_api_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_managed_api_version_proto_goTypes,
		DependencyIndexes: file_managed_api_version_proto_depIdxs,
		MessageInfos:      file_managed_api_version_proto_msgTypes,
		ExtensionInfos:    file_managed_api_version_proto_extTypes,
	}.Build()
	File_managed_api_version_proto = out.File
	file_managed_api_version_proto_rawDesc = nil
	file_managed_api_version_proto_goTypes = nil
	file_managed_api_version_proto_depIdxs = nil
}