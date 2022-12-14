// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: transaction/transaction_public_service.proto

package transaction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	managed "wego/pkg/grpc/managed"
	util "wego/pkg/grpc/util"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UtxSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size        int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	SizeInBytes int64 `protobuf:"varint,2,opt,name=size_in_bytes,json=sizeInBytes,proto3" json:"size_in_bytes,omitempty"`
}

func (x *UtxSize) Reset() {
	*x = UtxSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_transaction_public_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtxSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtxSize) ProtoMessage() {}

func (x *UtxSize) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_transaction_public_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtxSize.ProtoReflect.Descriptor instead.
func (*UtxSize) Descriptor() ([]byte, []int) {
	return file_transaction_transaction_public_service_proto_rawDescGZIP(), []int{0}
}

func (x *UtxSize) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UtxSize) GetSizeInBytes() int64 {
	if x != nil {
		return x.SizeInBytes
	}
	return 0
}

var File_transaction_transaction_public_service_proto protoreflect.FileDescriptor

var file_transaction_transaction_public_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x75, 0x74,
	0x69, 0x6c, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x07, 0x55, 0x74, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x32, 0xd1, 0x03, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x77,
	0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x12, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x73,
	0x12, 0x29, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x1c, 0x2e, 0x77, 0x61,
	0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x55, 0x74, 0x78,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x77,
	0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x74, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x30, 0x01, 0x12, 0x64, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x27, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x61, 0x76, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x1a, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x27, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x5f, 0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x19, 0x77, 0x65,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x0f, 0x57, 0x61, 0x76, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_transaction_transaction_public_service_proto_rawDescOnce sync.Once
	file_transaction_transaction_public_service_proto_rawDescData = file_transaction_transaction_public_service_proto_rawDesc
)

func file_transaction_transaction_public_service_proto_rawDescGZIP() []byte {
	file_transaction_transaction_public_service_proto_rawDescOnce.Do(func() {
		file_transaction_transaction_public_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_transaction_public_service_proto_rawDescData)
	})
	return file_transaction_transaction_public_service_proto_rawDescData
}

var file_transaction_transaction_public_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_transaction_public_service_proto_goTypes = []interface{}{
	(*UtxSize)(nil),                       // 0: wavesenterprise.grpc.UtxSize
	(*managed.Transaction)(nil),           // 1: wavesenterprise.Transaction
	(*util.TransactionWithCertChain)(nil), // 2: wavesenterprise.TransactionWithCertChain
	(*emptypb.Empty)(nil),                 // 3: google.protobuf.Empty
	(*util.TransactionInfoRequest)(nil),   // 4: wavesenterprise.TransactionInfoRequest
	(*util.TransactionInfoResponse)(nil),  // 5: wavesenterprise.TransactionInfoResponse
}
var file_transaction_transaction_public_service_proto_depIdxs = []int32{
	1, // 0: wavesenterprise.grpc.TransactionPublicService.Broadcast:input_type -> wavesenterprise.Transaction
	2, // 1: wavesenterprise.grpc.TransactionPublicService.BroadcastWithCerts:input_type -> wavesenterprise.TransactionWithCertChain
	3, // 2: wavesenterprise.grpc.TransactionPublicService.UtxInfo:input_type -> google.protobuf.Empty
	4, // 3: wavesenterprise.grpc.TransactionPublicService.TransactionInfo:input_type -> wavesenterprise.TransactionInfoRequest
	4, // 4: wavesenterprise.grpc.TransactionPublicService.UnconfirmedTransactionInfo:input_type -> wavesenterprise.TransactionInfoRequest
	1, // 5: wavesenterprise.grpc.TransactionPublicService.Broadcast:output_type -> wavesenterprise.Transaction
	1, // 6: wavesenterprise.grpc.TransactionPublicService.BroadcastWithCerts:output_type -> wavesenterprise.Transaction
	0, // 7: wavesenterprise.grpc.TransactionPublicService.UtxInfo:output_type -> wavesenterprise.grpc.UtxSize
	5, // 8: wavesenterprise.grpc.TransactionPublicService.TransactionInfo:output_type -> wavesenterprise.TransactionInfoResponse
	1, // 9: wavesenterprise.grpc.TransactionPublicService.UnconfirmedTransactionInfo:output_type -> wavesenterprise.Transaction
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transaction_transaction_public_service_proto_init() }
func file_transaction_transaction_public_service_proto_init() {
	if File_transaction_transaction_public_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_transaction_public_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtxSize); i {
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
			RawDescriptor: file_transaction_transaction_public_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_transaction_public_service_proto_goTypes,
		DependencyIndexes: file_transaction_transaction_public_service_proto_depIdxs,
		MessageInfos:      file_transaction_transaction_public_service_proto_msgTypes,
	}.Build()
	File_transaction_transaction_public_service_proto = out.File
	file_transaction_transaction_public_service_proto_rawDesc = nil
	file_transaction_transaction_public_service_proto_goTypes = nil
	file_transaction_transaction_public_service_proto_depIdxs = nil
}
