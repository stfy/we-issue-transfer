// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: contract/contract_util_service.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UtilServiceClient is the client API for UtilService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UtilServiceClient interface {
	GetNodeTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodeTimeResponse, error)
}

type utilServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilServiceClient(cc grpc.ClientConnInterface) UtilServiceClient {
	return &utilServiceClient{cc}
}

func (c *utilServiceClient) GetNodeTime(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodeTimeResponse, error) {
	out := new(NodeTimeResponse)
	err := c.cc.Invoke(ctx, "/wavesenterprise.UtilService/GetNodeTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilServiceServer is the server API for UtilService service.
// All implementations must embed UnimplementedUtilServiceServer
// for forward compatibility
type UtilServiceServer interface {
	GetNodeTime(context.Context, *emptypb.Empty) (*NodeTimeResponse, error)
	mustEmbedUnimplementedUtilServiceServer()
}

// UnimplementedUtilServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUtilServiceServer struct {
}

func (UnimplementedUtilServiceServer) GetNodeTime(context.Context, *emptypb.Empty) (*NodeTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeTime not implemented")
}
func (UnimplementedUtilServiceServer) mustEmbedUnimplementedUtilServiceServer() {}

// UnsafeUtilServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UtilServiceServer will
// result in compilation errors.
type UnsafeUtilServiceServer interface {
	mustEmbedUnimplementedUtilServiceServer()
}

func RegisterUtilServiceServer(s grpc.ServiceRegistrar, srv UtilServiceServer) {
	s.RegisterService(&UtilService_ServiceDesc, srv)
}

func _UtilService_GetNodeTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServiceServer).GetNodeTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wavesenterprise.UtilService/GetNodeTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServiceServer).GetNodeTime(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UtilService_ServiceDesc is the grpc.ServiceDesc for UtilService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UtilService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wavesenterprise.UtilService",
	HandlerType: (*UtilServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeTime",
			Handler:    _UtilService_GetNodeTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/contract_util_service.proto",
}