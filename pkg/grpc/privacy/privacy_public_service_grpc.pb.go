// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: privacy/privacy_public_service.proto

package privacy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	entity "wego/pkg/grpc/entity"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivacyPublicServiceClient is the client API for PrivacyPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivacyPublicServiceClient interface {
	GetPolicyItemData(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemDataResponse, error)
	GetPolicyItemLargeData(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (PrivacyPublicService_GetPolicyItemLargeDataClient, error)
	GetPolicyItemInfo(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemInfoResponse, error)
	PolicyItemDataExists(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemDataExistsResponse, error)
	SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error)
	SendLargeData(ctx context.Context, opts ...grpc.CallOption) (PrivacyPublicService_SendLargeDataClient, error)
	Recipients(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*entity.AddressesResponse, error)
	Owners(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*entity.AddressesResponse, error)
	Hashes(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*HashesResponse, error)
	ForceSync(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*ForceSyncResponse, error)
}

type privacyPublicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyPublicServiceClient(cc grpc.ClientConnInterface) PrivacyPublicServiceClient {
	return &privacyPublicServiceClient{cc}
}

func (c *privacyPublicServiceClient) GetPolicyItemData(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemDataResponse, error) {
	out := new(PolicyItemDataResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/GetPolicyItemData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) GetPolicyItemLargeData(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (PrivacyPublicService_GetPolicyItemLargeDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrivacyPublicService_ServiceDesc.Streams[0], "/wavesenterprise.grpc.PrivacyPublicService/GetPolicyItemLargeData", opts...)
	if err != nil {
		return nil, err
	}
	x := &privacyPublicServiceGetPolicyItemLargeDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrivacyPublicService_GetPolicyItemLargeDataClient interface {
	Recv() (*PolicyItemDataResponse, error)
	grpc.ClientStream
}

type privacyPublicServiceGetPolicyItemLargeDataClient struct {
	grpc.ClientStream
}

func (x *privacyPublicServiceGetPolicyItemLargeDataClient) Recv() (*PolicyItemDataResponse, error) {
	m := new(PolicyItemDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *privacyPublicServiceClient) GetPolicyItemInfo(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemInfoResponse, error) {
	out := new(PolicyItemInfoResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/GetPolicyItemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) PolicyItemDataExists(ctx context.Context, in *PolicyItemRequest, opts ...grpc.CallOption) (*PolicyItemDataExistsResponse, error) {
	out := new(PolicyItemDataExistsResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/PolicyItemDataExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error) {
	out := new(SendDataResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/SendData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) SendLargeData(ctx context.Context, opts ...grpc.CallOption) (PrivacyPublicService_SendLargeDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrivacyPublicService_ServiceDesc.Streams[1], "/wavesenterprise.grpc.PrivacyPublicService/SendLargeData", opts...)
	if err != nil {
		return nil, err
	}
	x := &privacyPublicServiceSendLargeDataClient{stream}
	return x, nil
}

type PrivacyPublicService_SendLargeDataClient interface {
	Send(*SendLargeDataRequest) error
	CloseAndRecv() (*SendDataResponse, error)
	grpc.ClientStream
}

type privacyPublicServiceSendLargeDataClient struct {
	grpc.ClientStream
}

func (x *privacyPublicServiceSendLargeDataClient) Send(m *SendLargeDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *privacyPublicServiceSendLargeDataClient) CloseAndRecv() (*SendDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *privacyPublicServiceClient) Recipients(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*entity.AddressesResponse, error) {
	out := new(entity.AddressesResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/Recipients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) Owners(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*entity.AddressesResponse, error) {
	out := new(entity.AddressesResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/Owners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) Hashes(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*HashesResponse, error) {
	out := new(HashesResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/Hashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyPublicServiceClient) ForceSync(ctx context.Context, in *PolicyIdRequest, opts ...grpc.CallOption) (*ForceSyncResponse, error) {
	out := new(ForceSyncResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.PrivacyPublicService/ForceSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivacyPublicServiceServer is the server API for PrivacyPublicService service.
// All implementations must embed UnimplementedPrivacyPublicServiceServer
// for forward compatibility
type PrivacyPublicServiceServer interface {
	GetPolicyItemData(context.Context, *PolicyItemRequest) (*PolicyItemDataResponse, error)
	GetPolicyItemLargeData(*PolicyItemRequest, PrivacyPublicService_GetPolicyItemLargeDataServer) error
	GetPolicyItemInfo(context.Context, *PolicyItemRequest) (*PolicyItemInfoResponse, error)
	PolicyItemDataExists(context.Context, *PolicyItemRequest) (*PolicyItemDataExistsResponse, error)
	SendData(context.Context, *SendDataRequest) (*SendDataResponse, error)
	SendLargeData(PrivacyPublicService_SendLargeDataServer) error
	Recipients(context.Context, *PolicyIdRequest) (*entity.AddressesResponse, error)
	Owners(context.Context, *PolicyIdRequest) (*entity.AddressesResponse, error)
	Hashes(context.Context, *PolicyIdRequest) (*HashesResponse, error)
	ForceSync(context.Context, *PolicyIdRequest) (*ForceSyncResponse, error)
	mustEmbedUnimplementedPrivacyPublicServiceServer()
}

// UnimplementedPrivacyPublicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivacyPublicServiceServer struct {
}

func (UnimplementedPrivacyPublicServiceServer) GetPolicyItemData(context.Context, *PolicyItemRequest) (*PolicyItemDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyItemData not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) GetPolicyItemLargeData(*PolicyItemRequest, PrivacyPublicService_GetPolicyItemLargeDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPolicyItemLargeData not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) GetPolicyItemInfo(context.Context, *PolicyItemRequest) (*PolicyItemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyItemInfo not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) PolicyItemDataExists(context.Context, *PolicyItemRequest) (*PolicyItemDataExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyItemDataExists not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) SendData(context.Context, *SendDataRequest) (*SendDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) SendLargeData(PrivacyPublicService_SendLargeDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendLargeData not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) Recipients(context.Context, *PolicyIdRequest) (*entity.AddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recipients not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) Owners(context.Context, *PolicyIdRequest) (*entity.AddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owners not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) Hashes(context.Context, *PolicyIdRequest) (*HashesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hashes not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) ForceSync(context.Context, *PolicyIdRequest) (*ForceSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceSync not implemented")
}
func (UnimplementedPrivacyPublicServiceServer) mustEmbedUnimplementedPrivacyPublicServiceServer() {}

// UnsafePrivacyPublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyPublicServiceServer will
// result in compilation errors.
type UnsafePrivacyPublicServiceServer interface {
	mustEmbedUnimplementedPrivacyPublicServiceServer()
}

func RegisterPrivacyPublicServiceServer(s grpc.ServiceRegistrar, srv PrivacyPublicServiceServer) {
	s.RegisterService(&PrivacyPublicService_ServiceDesc, srv)
}

func _PrivacyPublicService_GetPolicyItemData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).GetPolicyItemData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/GetPolicyItemData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).GetPolicyItemData(ctx, req.(*PolicyItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_GetPolicyItemLargeData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PolicyItemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrivacyPublicServiceServer).GetPolicyItemLargeData(m, &privacyPublicServiceGetPolicyItemLargeDataServer{stream})
}

type PrivacyPublicService_GetPolicyItemLargeDataServer interface {
	Send(*PolicyItemDataResponse) error
	grpc.ServerStream
}

type privacyPublicServiceGetPolicyItemLargeDataServer struct {
	grpc.ServerStream
}

func (x *privacyPublicServiceGetPolicyItemLargeDataServer) Send(m *PolicyItemDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PrivacyPublicService_GetPolicyItemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).GetPolicyItemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/GetPolicyItemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).GetPolicyItemInfo(ctx, req.(*PolicyItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_PolicyItemDataExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).PolicyItemDataExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/PolicyItemDataExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).PolicyItemDataExists(ctx, req.(*PolicyItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_SendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).SendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/SendData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).SendData(ctx, req.(*SendDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_SendLargeData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrivacyPublicServiceServer).SendLargeData(&privacyPublicServiceSendLargeDataServer{stream})
}

type PrivacyPublicService_SendLargeDataServer interface {
	SendAndClose(*SendDataResponse) error
	Recv() (*SendLargeDataRequest, error)
	grpc.ServerStream
}

type privacyPublicServiceSendLargeDataServer struct {
	grpc.ServerStream
}

func (x *privacyPublicServiceSendLargeDataServer) SendAndClose(m *SendDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *privacyPublicServiceSendLargeDataServer) Recv() (*SendLargeDataRequest, error) {
	m := new(SendLargeDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PrivacyPublicService_Recipients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).Recipients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/Recipients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).Recipients(ctx, req.(*PolicyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_Owners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).Owners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/Owners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).Owners(ctx, req.(*PolicyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_Hashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).Hashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/Hashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).Hashes(ctx, req.(*PolicyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivacyPublicService_ForceSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyPublicServiceServer).ForceSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.PrivacyPublicService/ForceSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyPublicServiceServer).ForceSync(ctx, req.(*PolicyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivacyPublicService_ServiceDesc is the grpc.ServiceDesc for PrivacyPublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivacyPublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wavesenterprise.grpc.PrivacyPublicService",
	HandlerType: (*PrivacyPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicyItemData",
			Handler:    _PrivacyPublicService_GetPolicyItemData_Handler,
		},
		{
			MethodName: "GetPolicyItemInfo",
			Handler:    _PrivacyPublicService_GetPolicyItemInfo_Handler,
		},
		{
			MethodName: "PolicyItemDataExists",
			Handler:    _PrivacyPublicService_PolicyItemDataExists_Handler,
		},
		{
			MethodName: "SendData",
			Handler:    _PrivacyPublicService_SendData_Handler,
		},
		{
			MethodName: "Recipients",
			Handler:    _PrivacyPublicService_Recipients_Handler,
		},
		{
			MethodName: "Owners",
			Handler:    _PrivacyPublicService_Owners_Handler,
		},
		{
			MethodName: "Hashes",
			Handler:    _PrivacyPublicService_Hashes_Handler,
		},
		{
			MethodName: "ForceSync",
			Handler:    _PrivacyPublicService_ForceSync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPolicyItemLargeData",
			Handler:       _PrivacyPublicService_GetPolicyItemLargeData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendLargeData",
			Handler:       _PrivacyPublicService_SendLargeData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "privacy/privacy_public_service.proto",
}