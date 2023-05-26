// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: livenessProbe.proto

package pb

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

const (
	LivenessProbe_LivenessProbe_FullMethodName = "/livenessProbe.LivenessProbe/LivenessProbe"
)

// LivenessProbeClient is the client API for LivenessProbe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivenessProbeClient interface {
	LivenessProbe(ctx context.Context, in *LivenessProbeRequest, opts ...grpc.CallOption) (*LivenessProbeResponse, error)
}

type livenessProbeClient struct {
	cc grpc.ClientConnInterface
}

func NewLivenessProbeClient(cc grpc.ClientConnInterface) LivenessProbeClient {
	return &livenessProbeClient{cc}
}

func (c *livenessProbeClient) LivenessProbe(ctx context.Context, in *LivenessProbeRequest, opts ...grpc.CallOption) (*LivenessProbeResponse, error) {
	out := new(LivenessProbeResponse)
	err := c.cc.Invoke(ctx, LivenessProbe_LivenessProbe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LivenessProbeServer is the server API for LivenessProbe service.
// All implementations must embed UnimplementedLivenessProbeServer
// for forward compatibility
type LivenessProbeServer interface {
	LivenessProbe(context.Context, *LivenessProbeRequest) (*LivenessProbeResponse, error)
	mustEmbedUnimplementedLivenessProbeServer()
}

// UnimplementedLivenessProbeServer must be embedded to have forward compatible implementations.
type UnimplementedLivenessProbeServer struct {
}

func (UnimplementedLivenessProbeServer) LivenessProbe(context.Context, *LivenessProbeRequest) (*LivenessProbeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LivenessProbe not implemented")
}
func (UnimplementedLivenessProbeServer) mustEmbedUnimplementedLivenessProbeServer() {}

// UnsafeLivenessProbeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivenessProbeServer will
// result in compilation errors.
type UnsafeLivenessProbeServer interface {
	mustEmbedUnimplementedLivenessProbeServer()
}

func RegisterLivenessProbeServer(s grpc.ServiceRegistrar, srv LivenessProbeServer) {
	s.RegisterService(&LivenessProbe_ServiceDesc, srv)
}

func _LivenessProbe_LivenessProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LivenessProbeServer).LivenessProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LivenessProbe_LivenessProbe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LivenessProbeServer).LivenessProbe(ctx, req.(*LivenessProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LivenessProbe_ServiceDesc is the grpc.ServiceDesc for LivenessProbe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivenessProbe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "livenessProbe.LivenessProbe",
	HandlerType: (*LivenessProbeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LivenessProbe",
			Handler:    _LivenessProbe_LivenessProbe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "livenessProbe.proto",
}