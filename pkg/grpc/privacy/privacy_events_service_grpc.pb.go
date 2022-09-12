// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: privacy/privacy_events_service.proto

package privacy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	events "wego/pkg/grpc/util/events"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivacyEventsServiceClient is the client API for PrivacyEventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivacyEventsServiceClient interface {
	SubscribeOn(ctx context.Context, in *events.SubscribeOnRequest, opts ...grpc.CallOption) (PrivacyEventsService_SubscribeOnClient, error)
}

type privacyEventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyEventsServiceClient(cc grpc.ClientConnInterface) PrivacyEventsServiceClient {
	return &privacyEventsServiceClient{cc}
}

func (c *privacyEventsServiceClient) SubscribeOn(ctx context.Context, in *events.SubscribeOnRequest, opts ...grpc.CallOption) (PrivacyEventsService_SubscribeOnClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrivacyEventsService_ServiceDesc.Streams[0], "/wavesenterprise.PrivacyEventsService/SubscribeOn", opts...)
	if err != nil {
		return nil, err
	}
	x := &privacyEventsServiceSubscribeOnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrivacyEventsService_SubscribeOnClient interface {
	Recv() (*PrivacyEvent, error)
	grpc.ClientStream
}

type privacyEventsServiceSubscribeOnClient struct {
	grpc.ClientStream
}

func (x *privacyEventsServiceSubscribeOnClient) Recv() (*PrivacyEvent, error) {
	m := new(PrivacyEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrivacyEventsServiceServer is the server API for PrivacyEventsService service.
// All implementations must embed UnimplementedPrivacyEventsServiceServer
// for forward compatibility
type PrivacyEventsServiceServer interface {
	SubscribeOn(*events.SubscribeOnRequest, PrivacyEventsService_SubscribeOnServer) error
	mustEmbedUnimplementedPrivacyEventsServiceServer()
}

// UnimplementedPrivacyEventsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivacyEventsServiceServer struct {
}

func (UnimplementedPrivacyEventsServiceServer) SubscribeOn(*events.SubscribeOnRequest, PrivacyEventsService_SubscribeOnServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOn not implemented")
}
func (UnimplementedPrivacyEventsServiceServer) mustEmbedUnimplementedPrivacyEventsServiceServer() {}

// UnsafePrivacyEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyEventsServiceServer will
// result in compilation errors.
type UnsafePrivacyEventsServiceServer interface {
	mustEmbedUnimplementedPrivacyEventsServiceServer()
}

func RegisterPrivacyEventsServiceServer(s grpc.ServiceRegistrar, srv PrivacyEventsServiceServer) {
	s.RegisterService(&PrivacyEventsService_ServiceDesc, srv)
}

func _PrivacyEventsService_SubscribeOn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(events.SubscribeOnRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrivacyEventsServiceServer).SubscribeOn(m, &privacyEventsServiceSubscribeOnServer{stream})
}

type PrivacyEventsService_SubscribeOnServer interface {
	Send(*PrivacyEvent) error
	grpc.ServerStream
}

type privacyEventsServiceSubscribeOnServer struct {
	grpc.ServerStream
}

func (x *privacyEventsServiceSubscribeOnServer) Send(m *PrivacyEvent) error {
	return x.ServerStream.SendMsg(m)
}

// PrivacyEventsService_ServiceDesc is the grpc.ServiceDesc for PrivacyEventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivacyEventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wavesenterprise.PrivacyEventsService",
	HandlerType: (*PrivacyEventsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOn",
			Handler:       _PrivacyEventsService_SubscribeOn_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "privacy/privacy_events_service.proto",
}
