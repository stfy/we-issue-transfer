// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: transaction/transaction_public_service.proto

package transaction

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	managed "wego/pkg/grpc/managed"
	util "wego/pkg/grpc/util"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransactionPublicServiceClient is the client API for TransactionPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionPublicServiceClient interface {
	Broadcast(ctx context.Context, in *managed.Transaction, opts ...grpc.CallOption) (*managed.Transaction, error)
	BroadcastWithCerts(ctx context.Context, in *util.TransactionWithCertChain, opts ...grpc.CallOption) (*managed.Transaction, error)
	UtxInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TransactionPublicService_UtxInfoClient, error)
	TransactionInfo(ctx context.Context, in *util.TransactionInfoRequest, opts ...grpc.CallOption) (*util.TransactionInfoResponse, error)
	UnconfirmedTransactionInfo(ctx context.Context, in *util.TransactionInfoRequest, opts ...grpc.CallOption) (*managed.Transaction, error)
}

type transactionPublicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionPublicServiceClient(cc grpc.ClientConnInterface) TransactionPublicServiceClient {
	return &transactionPublicServiceClient{cc}
}

func (c *transactionPublicServiceClient) Broadcast(ctx context.Context, in *managed.Transaction, opts ...grpc.CallOption) (*managed.Transaction, error) {
	out := new(managed.Transaction)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.TransactionPublicService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPublicServiceClient) BroadcastWithCerts(ctx context.Context, in *util.TransactionWithCertChain, opts ...grpc.CallOption) (*managed.Transaction, error) {
	out := new(managed.Transaction)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.TransactionPublicService/BroadcastWithCerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPublicServiceClient) UtxInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TransactionPublicService_UtxInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransactionPublicService_ServiceDesc.Streams[0], "/wavesenterprise.grpc.TransactionPublicService/UtxInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionPublicServiceUtxInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransactionPublicService_UtxInfoClient interface {
	Recv() (*UtxSize, error)
	grpc.ClientStream
}

type transactionPublicServiceUtxInfoClient struct {
	grpc.ClientStream
}

func (x *transactionPublicServiceUtxInfoClient) Recv() (*UtxSize, error) {
	m := new(UtxSize)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transactionPublicServiceClient) TransactionInfo(ctx context.Context, in *util.TransactionInfoRequest, opts ...grpc.CallOption) (*util.TransactionInfoResponse, error) {
	out := new(util.TransactionInfoResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.TransactionPublicService/TransactionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPublicServiceClient) UnconfirmedTransactionInfo(ctx context.Context, in *util.TransactionInfoRequest, opts ...grpc.CallOption) (*managed.Transaction, error) {
	out := new(managed.Transaction)
	err := c.cc.Invoke(ctx, "/wavesenterprise.grpc.TransactionPublicService/UnconfirmedTransactionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionPublicServiceServer is the server API for TransactionPublicService service.
// All implementations must embed UnimplementedTransactionPublicServiceServer
// for forward compatibility
type TransactionPublicServiceServer interface {
	Broadcast(context.Context, *managed.Transaction) (*managed.Transaction, error)
	BroadcastWithCerts(context.Context, *util.TransactionWithCertChain) (*managed.Transaction, error)
	UtxInfo(*emptypb.Empty, TransactionPublicService_UtxInfoServer) error
	TransactionInfo(context.Context, *util.TransactionInfoRequest) (*util.TransactionInfoResponse, error)
	UnconfirmedTransactionInfo(context.Context, *util.TransactionInfoRequest) (*managed.Transaction, error)
	mustEmbedUnimplementedTransactionPublicServiceServer()
}

// UnimplementedTransactionPublicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionPublicServiceServer struct {
}

func (UnimplementedTransactionPublicServiceServer) Broadcast(context.Context, *managed.Transaction) (*managed.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedTransactionPublicServiceServer) BroadcastWithCerts(context.Context, *util.TransactionWithCertChain) (*managed.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastWithCerts not implemented")
}
func (UnimplementedTransactionPublicServiceServer) UtxInfo(*emptypb.Empty, TransactionPublicService_UtxInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method UtxInfo not implemented")
}
func (UnimplementedTransactionPublicServiceServer) TransactionInfo(context.Context, *util.TransactionInfoRequest) (*util.TransactionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionInfo not implemented")
}
func (UnimplementedTransactionPublicServiceServer) UnconfirmedTransactionInfo(context.Context, *util.TransactionInfoRequest) (*managed.Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnconfirmedTransactionInfo not implemented")
}
func (UnimplementedTransactionPublicServiceServer) mustEmbedUnimplementedTransactionPublicServiceServer() {
}

// UnsafeTransactionPublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionPublicServiceServer will
// result in compilation errors.
type UnsafeTransactionPublicServiceServer interface {
	mustEmbedUnimplementedTransactionPublicServiceServer()
}

func RegisterTransactionPublicServiceServer(s grpc.ServiceRegistrar, srv TransactionPublicServiceServer) {
	s.RegisterService(&TransactionPublicService_ServiceDesc, srv)
}

func _TransactionPublicService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(managed.Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPublicServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.TransactionPublicService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPublicServiceServer).Broadcast(ctx, req.(*managed.Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionPublicService_BroadcastWithCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(util.TransactionWithCertChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPublicServiceServer).BroadcastWithCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.TransactionPublicService/BroadcastWithCerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPublicServiceServer).BroadcastWithCerts(ctx, req.(*util.TransactionWithCertChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionPublicService_UtxInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionPublicServiceServer).UtxInfo(m, &transactionPublicServiceUtxInfoServer{stream})
}

type TransactionPublicService_UtxInfoServer interface {
	Send(*UtxSize) error
	grpc.ServerStream
}

type transactionPublicServiceUtxInfoServer struct {
	grpc.ServerStream
}

func (x *transactionPublicServiceUtxInfoServer) Send(m *UtxSize) error {
	return x.ServerStream.SendMsg(m)
}

func _TransactionPublicService_TransactionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(util.TransactionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPublicServiceServer).TransactionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.TransactionPublicService/TransactionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPublicServiceServer).TransactionInfo(ctx, req.(*util.TransactionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionPublicService_UnconfirmedTransactionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(util.TransactionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionPublicServiceServer).UnconfirmedTransactionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.grpc.TransactionPublicService/UnconfirmedTransactionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionPublicServiceServer).UnconfirmedTransactionInfo(ctx, req.(*util.TransactionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionPublicService_ServiceDesc is the grpc.ServiceDesc for TransactionPublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionPublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wavesenterprise.grpc.TransactionPublicService",
	HandlerType: (*TransactionPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _TransactionPublicService_Broadcast_Handler,
		},
		{
			MethodName: "BroadcastWithCerts",
			Handler:    _TransactionPublicService_BroadcastWithCerts_Handler,
		},
		{
			MethodName: "TransactionInfo",
			Handler:    _TransactionPublicService_TransactionInfo_Handler,
		},
		{
			MethodName: "UnconfirmedTransactionInfo",
			Handler:    _TransactionPublicService_UnconfirmedTransactionInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UtxInfo",
			Handler:       _TransactionPublicService_UtxInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transaction/transaction_public_service.proto",
}
