// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: managed/reissue_transaction.proto

package managed

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

type ReissueTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChainId         int32    `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	SenderPublicKey []byte   `protobuf:"bytes,3,opt,name=sender_public_key,json=senderPublicKey,proto3" json:"sender_public_key,omitempty"`
	AssetId         []byte   `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Quantity        int64    `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Reissuable      bool     `protobuf:"varint,6,opt,name=reissuable,proto3" json:"reissuable,omitempty"`
	Fee             int64    `protobuf:"varint,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Timestamp       int64    `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Proofs          [][]byte `protobuf:"bytes,9,rep,name=proofs,proto3" json:"proofs,omitempty"`
	SenderAddress   []byte   `protobuf:"bytes,10,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
}

func (x *ReissueTransaction) Reset() {
	*x = ReissueTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managed_reissue_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReissueTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReissueTransaction) ProtoMessage() {}

func (x *ReissueTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_managed_reissue_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReissueTransaction.ProtoReflect.Descriptor instead.
func (*ReissueTransaction) Descriptor() ([]byte, []int) {
	return file_managed_reissue_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ReissueTransaction) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ReissueTransaction) GetChainId() int32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *ReissueTransaction) GetSenderPublicKey() []byte {
	if x != nil {
		return x.SenderPublicKey
	}
	return nil
}

func (x *ReissueTransaction) GetAssetId() []byte {
	if x != nil {
		return x.AssetId
	}
	return nil
}

func (x *ReissueTransaction) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ReissueTransaction) GetReissuable() bool {
	if x != nil {
		return x.Reissuable
	}
	return false
}

func (x *ReissueTransaction) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *ReissueTransaction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ReissueTransaction) GetProofs() [][]byte {
	if x != nil {
		return x.Proofs
	}
	return nil
}

func (x *ReissueTransaction) GetSenderAddress() []byte {
	if x != nil {
		return x.SenderAddress
	}
	return nil
}

var File_managed_reissue_transaction_proto protoreflect.FileDescriptor

var file_managed_reissue_transaction_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x22, 0xb1, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x69,
	0x73, 0x73, 0x75, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72,
	0x65, 0x69, 0x73, 0x73, 0x75, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x5c, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x50, 0x01, 0x5a, 0x15, 0x77,
	0x65, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0xaa, 0x02, 0x0f, 0x57, 0x61, 0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managed_reissue_transaction_proto_rawDescOnce sync.Once
	file_managed_reissue_transaction_proto_rawDescData = file_managed_reissue_transaction_proto_rawDesc
)

func file_managed_reissue_transaction_proto_rawDescGZIP() []byte {
	file_managed_reissue_transaction_proto_rawDescOnce.Do(func() {
		file_managed_reissue_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_managed_reissue_transaction_proto_rawDescData)
	})
	return file_managed_reissue_transaction_proto_rawDescData
}

var file_managed_reissue_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_managed_reissue_transaction_proto_goTypes = []interface{}{
	(*ReissueTransaction)(nil), // 0: wavesenterprise.ReissueTransaction
}
var file_managed_reissue_transaction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managed_reissue_transaction_proto_init() }
func file_managed_reissue_transaction_proto_init() {
	if File_managed_reissue_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managed_reissue_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReissueTransaction); i {
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
			RawDescriptor: file_managed_reissue_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managed_reissue_transaction_proto_goTypes,
		DependencyIndexes: file_managed_reissue_transaction_proto_depIdxs,
		MessageInfos:      file_managed_reissue_transaction_proto_msgTypes,
	}.Build()
	File_managed_reissue_transaction_proto = out.File
	file_managed_reissue_transaction_proto_rawDesc = nil
	file_managed_reissue_transaction_proto_goTypes = nil
	file_managed_reissue_transaction_proto_depIdxs = nil
}
