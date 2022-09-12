// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: util/util_node_info_service.proto

package util

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type NodeConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                     string               `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CryptoType                  grpc.CryptoType      `protobuf:"varint,2,opt,name=crypto_type,json=cryptoType,proto3,enum=wavesenterprise.CryptoType" json:"crypto_type,omitempty"`
	ChainId                     int32                `protobuf:"varint,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Consensus                   grpc.ConsensusType   `protobuf:"varint,4,opt,name=consensus,proto3,enum=wavesenterprise.ConsensusType" json:"consensus,omitempty"`
	MinimumFee                  map[string]int64     `protobuf:"bytes,5,rep,name=minimum_fee,json=minimumFee,proto3" json:"minimum_fee,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	AdditionalFee               map[string]int64     `protobuf:"bytes,6,rep,name=additional_fee,json=additionalFee,proto3" json:"additional_fee,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MaxTransactionsInMicroBlock int32                `protobuf:"varint,7,opt,name=max_transactions_in_micro_block,json=maxTransactionsInMicroBlock,proto3" json:"max_transactions_in_micro_block,omitempty"`
	MinMicroBlockAge            *durationpb.Duration `protobuf:"bytes,8,opt,name=min_micro_block_age,json=minMicroBlockAge,proto3" json:"min_micro_block_age,omitempty"`
	MicroBlockInterval          *durationpb.Duration `protobuf:"bytes,9,opt,name=micro_block_interval,json=microBlockInterval,proto3" json:"micro_block_interval,omitempty"`
	PkiMode                     grpc.PkiMode         `protobuf:"varint,10,opt,name=pki_mode,json=pkiMode,proto3,enum=wavesenterprise.PkiMode" json:"pki_mode,omitempty"`
	RequiredOids                []string             `protobuf:"bytes,11,rep,name=required_oids,json=requiredOids,proto3" json:"required_oids,omitempty"`
	CrlCheckEnabled             bool                 `protobuf:"varint,12,opt,name=crl_check_enabled,json=crlCheckEnabled,proto3" json:"crl_check_enabled,omitempty"`
	// Types that are assignable to BlockTiming:
	//	*NodeConfigResponse_PoaRoundInfo
	//	*NodeConfigResponse_PosRoundInfo
	BlockTiming isNodeConfigResponse_BlockTiming `protobuf_oneof:"block_timing"`
}

func (x *NodeConfigResponse) Reset() {
	*x = NodeConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_util_node_info_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConfigResponse) ProtoMessage() {}

func (x *NodeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_util_util_node_info_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConfigResponse.ProtoReflect.Descriptor instead.
func (*NodeConfigResponse) Descriptor() ([]byte, []int) {
	return file_util_util_node_info_service_proto_rawDescGZIP(), []int{0}
}

func (x *NodeConfigResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeConfigResponse) GetCryptoType() grpc.CryptoType {
	if x != nil {
		return x.CryptoType
	}
	return grpc.CryptoType(0)
}

func (x *NodeConfigResponse) GetChainId() int32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *NodeConfigResponse) GetConsensus() grpc.ConsensusType {
	if x != nil {
		return x.Consensus
	}
	return grpc.ConsensusType(0)
}

func (x *NodeConfigResponse) GetMinimumFee() map[string]int64 {
	if x != nil {
		return x.MinimumFee
	}
	return nil
}

func (x *NodeConfigResponse) GetAdditionalFee() map[string]int64 {
	if x != nil {
		return x.AdditionalFee
	}
	return nil
}

func (x *NodeConfigResponse) GetMaxTransactionsInMicroBlock() int32 {
	if x != nil {
		return x.MaxTransactionsInMicroBlock
	}
	return 0
}

func (x *NodeConfigResponse) GetMinMicroBlockAge() *durationpb.Duration {
	if x != nil {
		return x.MinMicroBlockAge
	}
	return nil
}

func (x *NodeConfigResponse) GetMicroBlockInterval() *durationpb.Duration {
	if x != nil {
		return x.MicroBlockInterval
	}
	return nil
}

func (x *NodeConfigResponse) GetPkiMode() grpc.PkiMode {
	if x != nil {
		return x.PkiMode
	}
	return grpc.PkiMode(0)
}

func (x *NodeConfigResponse) GetRequiredOids() []string {
	if x != nil {
		return x.RequiredOids
	}
	return nil
}

func (x *NodeConfigResponse) GetCrlCheckEnabled() bool {
	if x != nil {
		return x.CrlCheckEnabled
	}
	return false
}

func (m *NodeConfigResponse) GetBlockTiming() isNodeConfigResponse_BlockTiming {
	if m != nil {
		return m.BlockTiming
	}
	return nil
}

func (x *NodeConfigResponse) GetPoaRoundInfo() *PoaRoundInfo {
	if x, ok := x.GetBlockTiming().(*NodeConfigResponse_PoaRoundInfo); ok {
		return x.PoaRoundInfo
	}
	return nil
}

func (x *NodeConfigResponse) GetPosRoundInfo() *PosRoundInfo {
	if x, ok := x.GetBlockTiming().(*NodeConfigResponse_PosRoundInfo); ok {
		return x.PosRoundInfo
	}
	return nil
}

type isNodeConfigResponse_BlockTiming interface {
	isNodeConfigResponse_BlockTiming()
}

type NodeConfigResponse_PoaRoundInfo struct {
	PoaRoundInfo *PoaRoundInfo `protobuf:"bytes,101,opt,name=poa_round_info,json=poaRoundInfo,proto3,oneof"`
}

type NodeConfigResponse_PosRoundInfo struct {
	PosRoundInfo *PosRoundInfo `protobuf:"bytes,102,opt,name=pos_round_info,json=posRoundInfo,proto3,oneof"`
}

func (*NodeConfigResponse_PoaRoundInfo) isNodeConfigResponse_BlockTiming() {}

func (*NodeConfigResponse_PosRoundInfo) isNodeConfigResponse_BlockTiming() {}

type PoaRoundInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundDuration *durationpb.Duration `protobuf:"bytes,1,opt,name=round_duration,json=roundDuration,proto3" json:"round_duration,omitempty"`
	SyncDuration  *durationpb.Duration `protobuf:"bytes,2,opt,name=sync_duration,json=syncDuration,proto3" json:"sync_duration,omitempty"`
}

func (x *PoaRoundInfo) Reset() {
	*x = PoaRoundInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_util_node_info_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoaRoundInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoaRoundInfo) ProtoMessage() {}

func (x *PoaRoundInfo) ProtoReflect() protoreflect.Message {
	mi := &file_util_util_node_info_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoaRoundInfo.ProtoReflect.Descriptor instead.
func (*PoaRoundInfo) Descriptor() ([]byte, []int) {
	return file_util_util_node_info_service_proto_rawDescGZIP(), []int{1}
}

func (x *PoaRoundInfo) GetRoundDuration() *durationpb.Duration {
	if x != nil {
		return x.RoundDuration
	}
	return nil
}

func (x *PoaRoundInfo) GetSyncDuration() *durationpb.Duration {
	if x != nil {
		return x.SyncDuration
	}
	return nil
}

type PosRoundInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AverageBlockDelay *durationpb.Duration `protobuf:"bytes,1,opt,name=average_block_delay,json=averageBlockDelay,proto3" json:"average_block_delay,omitempty"`
}

func (x *PosRoundInfo) Reset() {
	*x = PosRoundInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_util_node_info_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosRoundInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosRoundInfo) ProtoMessage() {}

func (x *PosRoundInfo) ProtoReflect() protoreflect.Message {
	mi := &file_util_util_node_info_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosRoundInfo.ProtoReflect.Descriptor instead.
func (*PosRoundInfo) Descriptor() ([]byte, []int) {
	return file_util_util_node_info_service_proto_rawDescGZIP(), []int{2}
}

func (x *PosRoundInfo) GetAverageBlockDelay() *durationpb.Duration {
	if x != nil {
		return x.AverageBlockDelay
	}
	return nil
}

type AddressWithPubKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *AddressWithPubKeyResponse) Reset() {
	*x = AddressWithPubKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_util_node_info_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressWithPubKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressWithPubKeyResponse) ProtoMessage() {}

func (x *AddressWithPubKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_util_util_node_info_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressWithPubKeyResponse.ProtoReflect.Descriptor instead.
func (*AddressWithPubKeyResponse) Descriptor() ([]byte, []int) {
	return file_util_util_node_info_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddressWithPubKeyResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressWithPubKeyResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_util_util_node_info_service_proto protoreflect.FileDescriptor

var file_util_util_node_info_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfc, 0x07, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x54, 0x0a, 0x0b, 0x6d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x46, 0x65, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x46, 0x65, 0x65,
	0x12, 0x5d, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12,
	0x44, 0x0a, 0x1f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x48, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d,
	0x69, 0x6e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x67, 0x65, 0x12,
	0x4b, 0x0a, 0x14, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08,
	0x70, 0x6b, 0x69, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x50, 0x6b, 0x69, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x70, 0x6b, 0x69, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x69,
	0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x4f, 0x69, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x6c, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x63, 0x72, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x6f, 0x61, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x61,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x6f, 0x61,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x6f, 0x73,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x3d, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x46, 0x65, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x40, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x69, 0x6e,
	0x67, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x61, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x13, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22,
	0x54, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x32, 0xad, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x23, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x77, 0x61, 0x76, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x53, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x76,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x50, 0x01, 0x5a, 0x12, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0xaa, 0x02, 0x0f, 0x57, 0x61, 0x76, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_util_util_node_info_service_proto_rawDescOnce sync.Once
	file_util_util_node_info_service_proto_rawDescData = file_util_util_node_info_service_proto_rawDesc
)

func file_util_util_node_info_service_proto_rawDescGZIP() []byte {
	file_util_util_node_info_service_proto_rawDescOnce.Do(func() {
		file_util_util_node_info_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_util_util_node_info_service_proto_rawDescData)
	})
	return file_util_util_node_info_service_proto_rawDescData
}

var file_util_util_node_info_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_util_util_node_info_service_proto_goTypes = []interface{}{
	(*NodeConfigResponse)(nil),        // 0: wavesenterprise.NodeConfigResponse
	(*PoaRoundInfo)(nil),              // 1: wavesenterprise.PoaRoundInfo
	(*PosRoundInfo)(nil),              // 2: wavesenterprise.PosRoundInfo
	(*AddressWithPubKeyResponse)(nil), // 3: wavesenterprise.AddressWithPubKeyResponse
	nil,                               // 4: wavesenterprise.NodeConfigResponse.MinimumFeeEntry
	nil,                               // 5: wavesenterprise.NodeConfigResponse.AdditionalFeeEntry
	(grpc.CryptoType)(0),              // 6: wavesenterprise.CryptoType
	(grpc.ConsensusType)(0),           // 7: wavesenterprise.ConsensusType
	(*durationpb.Duration)(nil),       // 8: google.protobuf.Duration
	(grpc.PkiMode)(0),                 // 9: wavesenterprise.PkiMode
	(*emptypb.Empty)(nil),             // 10: google.protobuf.Empty
}
var file_util_util_node_info_service_proto_depIdxs = []int32{
	6,  // 0: wavesenterprise.NodeConfigResponse.crypto_type:type_name -> wavesenterprise.CryptoType
	7,  // 1: wavesenterprise.NodeConfigResponse.consensus:type_name -> wavesenterprise.ConsensusType
	4,  // 2: wavesenterprise.NodeConfigResponse.minimum_fee:type_name -> wavesenterprise.NodeConfigResponse.MinimumFeeEntry
	5,  // 3: wavesenterprise.NodeConfigResponse.additional_fee:type_name -> wavesenterprise.NodeConfigResponse.AdditionalFeeEntry
	8,  // 4: wavesenterprise.NodeConfigResponse.min_micro_block_age:type_name -> google.protobuf.Duration
	8,  // 5: wavesenterprise.NodeConfigResponse.micro_block_interval:type_name -> google.protobuf.Duration
	9,  // 6: wavesenterprise.NodeConfigResponse.pki_mode:type_name -> wavesenterprise.PkiMode
	1,  // 7: wavesenterprise.NodeConfigResponse.poa_round_info:type_name -> wavesenterprise.PoaRoundInfo
	2,  // 8: wavesenterprise.NodeConfigResponse.pos_round_info:type_name -> wavesenterprise.PosRoundInfo
	8,  // 9: wavesenterprise.PoaRoundInfo.round_duration:type_name -> google.protobuf.Duration
	8,  // 10: wavesenterprise.PoaRoundInfo.sync_duration:type_name -> google.protobuf.Duration
	8,  // 11: wavesenterprise.PosRoundInfo.average_block_delay:type_name -> google.protobuf.Duration
	10, // 12: wavesenterprise.NodeInfoService.NodeConfig:input_type -> google.protobuf.Empty
	10, // 13: wavesenterprise.NodeInfoService.NodeOwner:input_type -> google.protobuf.Empty
	0,  // 14: wavesenterprise.NodeInfoService.NodeConfig:output_type -> wavesenterprise.NodeConfigResponse
	3,  // 15: wavesenterprise.NodeInfoService.NodeOwner:output_type -> wavesenterprise.AddressWithPubKeyResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_util_util_node_info_service_proto_init() }
func file_util_util_node_info_service_proto_init() {
	if File_util_util_node_info_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_util_util_node_info_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeConfigResponse); i {
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
		file_util_util_node_info_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoaRoundInfo); i {
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
		file_util_util_node_info_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosRoundInfo); i {
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
		file_util_util_node_info_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressWithPubKeyResponse); i {
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
	file_util_util_node_info_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NodeConfigResponse_PoaRoundInfo)(nil),
		(*NodeConfigResponse_PosRoundInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_util_util_node_info_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_util_util_node_info_service_proto_goTypes,
		DependencyIndexes: file_util_util_node_info_service_proto_depIdxs,
		MessageInfos:      file_util_util_node_info_service_proto_msgTypes,
	}.Build()
	File_util_util_node_info_service_proto = out.File
	file_util_util_node_info_service_proto_rawDesc = nil
	file_util_util_node_info_service_proto_goTypes = nil
	file_util_util_node_info_service_proto_depIdxs = nil
}