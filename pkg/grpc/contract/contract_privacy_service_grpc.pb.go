// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: contract/contract_privacy_service.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivacyServiceClient is the client API for PrivacyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivacyServiceClient interface {
	GetPolicyRecipients(ctx context.Context, in *PolicyRecipientsRequest, opts ...grpc.CallOption) (*PolicyRecipientsResponse, error)
	GetPolicyOwners(ctx context.Context, in *PolicyOwnersRequest, opts ...grpc.CallOption) (*PolicyOwnersResponse, error)
	GetPolicyHashes(ctx context.Context, in *PolicyHashesRequest, opts ...grpc.CallOption) (*PolicyHashesResponse, error)
	GetPolicyItemData(ctx context.Context, in *PolicyItemDataRequest, opts ...grpc.CallOption) (*PolicyItemDataResponse, error)
	GetPolicyItemInfo(ctx context.Context, in *PolicyItemInfoRequest, opts ...grpc.CallOption) (*PolicyItemInfoResponse, error)
}

type privacyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyServiceClient(cc grpc.ClientConnInterface) PrivacyServiceClient {
	return &privacyServiceClient{cc}
}

func (c *privacyServiceClient) GetPolicyRecipients(ctx context.Context, in *PolicyRecipientsRequest, opts ...grpc.CallOption) (*PolicyRecipientsResponse, error) {
	out := new(PolicyRecipientsResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.PrivacyService/GetPolicyRecipients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyServiceClient) GetPolicyOwners(ctx context.Context, in *PolicyOwnersRequest, opts ...grpc.CallOption) (*PolicyOwnersResponse, error) {
	out := new(PolicyOwnersResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.PrivacyService/GetPolicyOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyServiceClient) GetPolicyHashes(ctx context.Context, in *PolicyHashesRequest, opts ...grpc.CallOption) (*PolicyHashesResponse, error) {
	out := new(PolicyHashesResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.PrivacyService/GetPolicyHashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyServiceClient) GetPolicyItemData(ctx context.Context, in *PolicyItemDataRequest, opts ...grpc.CallOption) (*PolicyItemDataResponse, error) {
	out := new(PolicyItemDataResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.PrivacyService/GetPolicyItemData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyServiceClient) GetPolicyItemInfo(ctx context.Context, in *PolicyItemInfoRequest, opts ...grpc.CallOption) (*PolicyItemInfoResponse, error) {
	out := new(PolicyItemInfoResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.PrivacyService/GetPolicyItemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivacyServiceServer is the server API for PrivacyService service.
// All implementations must embed UnimplementedPrivacyServiceServer
// for forward compatibility
type PrivacyServiceServer interface {
	GetPolicyRecipients(context.Context, *PolicyRecipientsRequest) (*PolicyRecipientsResponse, error)
	GetPolicyOwners(context.Context, *PolicyOwnersRequest) (*PolicyOwnersResponse, error)
	GetPolicyHashes(context.Context, *PolicyHashesRequest) (*PolicyHashesResponse, error)
	GetPolicyItemData(context.Context, *PolicyItemDataRequest) (*PolicyItemDataResponse, error)
	GetPolicyItemInfo(context.Context, *PolicyItemInfoRequest) (*PolicyItemInfoResponse, error)
	mustEmbedUnimplementedPrivacyServiceServer()
}

// UnimplementedPrivacyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivacyServiceServer struct {
}

func (UnimplementedPrivacyServiceServer) GetPolicyRecipients(context.Context, *PolicyRecipientsRequest) (*PolicyRecipientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyRecipients not implemented")
}
func (UnimplementedPrivacyServiceServer) GetPolicyOwners(context.Context, *PolicyOwnersRequest) (*PolicyOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyOwners not implemented")
}
func (UnimplementedPrivacyServiceServer) GetPolicyHashes(context.Context, *PolicyHashesRequest) (*PolicyHashesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyHashes not implemented")
}
func (UnimplementedPrivacyServiceServer) GetPolicyItemData(context.Context, *PolicyItemDataRequest) (*PolicyItemDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyItemData not implemented")
}
func (UnimplementedPrivacyServiceServer) GetPolicyItemInfo(context.Context, *PolicyItemInfoRequest) (*PolicyItemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyItemInfo not implemented")
}
func (UnimplementedPrivacyServiceServer) mustEmbedUnimplementedPrivacyServiceServer() {}

// UnsafePrivacyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyServiceServer will
// result in compilation errors.
type UnsafePrivacyServiceServer interface {
	mustEmbedUnimplementedPrivacyServiceServer()
}

func RegisterPrivacyServiceServer(s grpc.ServiceRegistrar, srv PrivacyServiceServer) {
	s.RegisterService(&PrivacyService_ServiceDesc, srv)
}

func _PrivacyService_GetPolicyRecipients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRecipientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServiceServer).GetPolicyRecipients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.PrivacyService/GetPolicyRecipients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServiceServer).GetPolicyRecipients(ctx, req.(*PolicyRecipientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyService_GetPolicyOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServiceServer).GetPolicyOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.PrivacyService/GetPolicyOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServiceServer).GetPolicyOwners(ctx, req.(*PolicyOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyService_GetPolicyHashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyHashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServiceServer).GetPolicyHashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.PrivacyService/GetPolicyHashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServiceServer).GetPolicyHashes(ctx, req.(*PolicyHashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyService_GetPolicyItemData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyItemDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServiceServer).GetPolicyItemData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.PrivacyService/GetPolicyItemData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServiceServer).GetPolicyItemData(ctx, req.(*PolicyItemDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyService_GetPolicyItemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyItemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServiceServer).GetPolicyItemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.PrivacyService/GetPolicyItemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServiceServer).GetPolicyItemInfo(ctx, req.(*PolicyItemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivacyService_ServiceDesc is the grpc.ServiceDesc for PrivacyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivacyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wavesenterprise.PrivacyService",
	HandlerType: (*PrivacyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicyRecipients",
			Handler:    _PrivacyService_GetPolicyRecipients_Handler,
		},
		{
			MethodName: "GetPolicyOwners",
			Handler:    _PrivacyService_GetPolicyOwners_Handler,
		},
		{
			MethodName: "GetPolicyHashes",
			Handler:    _PrivacyService_GetPolicyHashes_Handler,
		},
		{
			MethodName: "GetPolicyItemData",
			Handler:    _PrivacyService_GetPolicyItemData_Handler,
		},
		{
			MethodName: "GetPolicyItemInfo",
			Handler:    _PrivacyService_GetPolicyItemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/contract_privacy_service.proto",
}