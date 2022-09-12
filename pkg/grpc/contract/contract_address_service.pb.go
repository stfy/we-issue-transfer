// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: contract/contract_address_service.proto

package contract

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	grpc "wego/pkg/grpc"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *AddressesResponse) Reset() {
	*x = AddressesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_address_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressesResponse) ProtoMessage() {}

func (x *AddressesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_address_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressesResponse.ProtoReflect.Descriptor instead.
func (*AddressesResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_address_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddressesResponse) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type AddressDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Limit   *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *AddressDataRequest) Reset() {
	*x = AddressDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_address_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressDataRequest) ProtoMessage() {}

func (x *AddressDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_address_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressDataRequest.ProtoReflect.Descriptor instead.
func (*AddressDataRequest) Descriptor() ([]byte, []int) {
	return file_contract_contract_address_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddressDataRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressDataRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *AddressDataRequest) GetOffset() *wrapperspb.Int32Value {
	if x != nil {
		return x.Offset
	}
	return nil
}

type AddressDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*grpc.DataEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *AddressDataResponse) Reset() {
	*x = AddressDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_address_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressDataResponse) ProtoMessage() {}

func (x *AddressDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_address_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressDataResponse.ProtoReflect.Descriptor instead.
func (*AddressDataResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_address_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddressDataResponse) GetEntries() []*grpc.DataEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AssetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AssetId *wrapperspb.BytesValue `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *AssetBalanceRequest) Reset() {
	*x = AssetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_address_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetBalanceRequest) ProtoMessage() {}

func (x *AssetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_address_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetBalanceRequest.ProtoReflect.Descriptor instead.
func (*AssetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_contract_contract_address_service_proto_rawDescGZIP(), []int{3}
}

func (x *AssetBalanceRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *AssetBalanceRequest) GetAssetId() *wrapperspb.BytesValue {
	if x != nil {
		return x.AssetId
	}
	return nil
}

type AssetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId  *wrapperspb.BytesValue `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount   int64                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Decimals int32                  `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *AssetBalanceResponse) Reset() {
	*x = AssetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_address_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetBalanceResponse) ProtoMessage() {}

func (x *AssetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_address_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetBalanceResponse.ProtoReflect.Descriptor instead.
func (*AssetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_contract_contract_address_service_proto_rawDescGZIP(), []int{4}
}

func (x *AssetBalanceResponse) GetAssetId() *wrapperspb.BytesValue {
	if x != nil {
		return x.AssetId
	}
	return nil
}

func (x *AssetBalanceResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AssetBalanceResponse) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

var File_contract_contract_address_service_proto protoreflect.FileDescriptor

var file_contract_contract_address_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x76, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x96, 0x01, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x33, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4b, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x67, 0x0a, 0x13, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x14,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73,
	0x32, 0x99, 0x02, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x77, 0x61,
	0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x24, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5b, 0x0a, 0x2d,
	0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x01, 0x5a,
	0x16, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0xaa, 0x02, 0x0f, 0x57, 0x61, 0x76, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contract_contract_address_service_proto_rawDescOnce sync.Once
	file_contract_contract_address_service_proto_rawDescData = file_contract_contract_address_service_proto_rawDesc
)

func file_contract_contract_address_service_proto_rawDescGZIP() []byte {
	file_contract_contract_address_service_proto_rawDescOnce.Do(func() {
		file_contract_contract_address_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_contract_address_service_proto_rawDescData)
	})
	return file_contract_contract_address_service_proto_rawDescData
}

var file_contract_contract_address_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contract_contract_address_service_proto_goTypes = []interface{}{
	(*AddressesResponse)(nil),     // 0: wavesenterprise.AddressesResponse
	(*AddressDataRequest)(nil),    // 1: wavesenterprise.AddressDataRequest
	(*AddressDataResponse)(nil),   // 2: wavesenterprise.AddressDataResponse
	(*AssetBalanceRequest)(nil),   // 3: wavesenterprise.AssetBalanceRequest
	(*AssetBalanceResponse)(nil),  // 4: wavesenterprise.AssetBalanceResponse
	(*wrapperspb.Int32Value)(nil), // 5: google.protobuf.Int32Value
	(*grpc.DataEntry)(nil),        // 6: wavesenterprise.DataEntry
	(*wrapperspb.BytesValue)(nil), // 7: google.protobuf.BytesValue
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_contract_contract_address_service_proto_depIdxs = []int32{
	5, // 0: wavesenterprise.AddressDataRequest.limit:type_name -> google.protobuf.Int32Value
	5, // 1: wavesenterprise.AddressDataRequest.offset:type_name -> google.protobuf.Int32Value
	6, // 2: wavesenterprise.AddressDataResponse.entries:type_name -> wavesenterprise.DataEntry
	7, // 3: wavesenterprise.AssetBalanceRequest.asset_id:type_name -> google.protobuf.BytesValue
	7, // 4: wavesenterprise.AssetBalanceResponse.asset_id:type_name -> google.protobuf.BytesValue
	8, // 5: wavesenterprise.AddressService.GetAddresses:input_type -> google.protobuf.Empty
	1, // 6: wavesenterprise.AddressService.GetAddressData:input_type -> wavesenterprise.AddressDataRequest
	3, // 7: wavesenterprise.AddressService.GetAssetBalance:input_type -> wavesenterprise.AssetBalanceRequest
	0, // 8: wavesenterprise.AddressService.GetAddresses:output_type -> wavesenterprise.AddressesResponse
	2, // 9: wavesenterprise.AddressService.GetAddressData:output_type -> wavesenterprise.AddressDataResponse
	4, // 10: wavesenterprise.AddressService.GetAssetBalance:output_type -> wavesenterprise.AssetBalanceResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contract_contract_address_service_proto_init() }
func file_contract_contract_address_service_proto_init() {
	if File_contract_contract_address_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_contract_address_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressesResponse); i {
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
		file_contract_contract_address_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressDataRequest); i {
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
		file_contract_contract_address_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressDataResponse); i {
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
		file_contract_contract_address_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetBalanceRequest); i {
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
		file_contract_contract_address_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetBalanceResponse); i {
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
			RawDescriptor: file_contract_contract_address_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_contract_address_service_proto_goTypes,
		DependencyIndexes: file_contract_contract_address_service_proto_depIdxs,
		MessageInfos:      file_contract_contract_address_service_proto_msgTypes,
	}.Build()
	File_contract_contract_address_service_proto = out.File
	file_contract_contract_address_service_proto_rawDesc = nil
	file_contract_contract_address_service_proto_goTypes = nil
	file_contract_contract_address_service_proto_depIdxs = nil
}
