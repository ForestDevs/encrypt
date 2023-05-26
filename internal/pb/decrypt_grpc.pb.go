// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: decrypt.proto

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
	DecryptService_DecryptContent_FullMethodName = "/decrypt.DecryptService/DecryptContent"
	DecryptService_DecryptFile_FullMethodName    = "/decrypt.DecryptService/DecryptFile"
)

// DecryptServiceClient is the client API for DecryptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecryptServiceClient interface {
	DecryptContent(ctx context.Context, in *DecryptContentRequest, opts ...grpc.CallOption) (*DecryptContentResponse, error)
	DecryptFile(ctx context.Context, in *DecryptFileRequest, opts ...grpc.CallOption) (*DecryptFileResponse, error)
}

type decryptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDecryptServiceClient(cc grpc.ClientConnInterface) DecryptServiceClient {
	return &decryptServiceClient{cc}
}

func (c *decryptServiceClient) DecryptContent(ctx context.Context, in *DecryptContentRequest, opts ...grpc.CallOption) (*DecryptContentResponse, error) {
	out := new(DecryptContentResponse)
	err := c.cc.Invoke(ctx, DecryptService_DecryptContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decryptServiceClient) DecryptFile(ctx context.Context, in *DecryptFileRequest, opts ...grpc.CallOption) (*DecryptFileResponse, error) {
	out := new(DecryptFileResponse)
	err := c.cc.Invoke(ctx, DecryptService_DecryptFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecryptServiceServer is the server API for DecryptService service.
// All implementations must embed UnimplementedDecryptServiceServer
// for forward compatibility
type DecryptServiceServer interface {
	DecryptContent(context.Context, *DecryptContentRequest) (*DecryptContentResponse, error)
	DecryptFile(context.Context, *DecryptFileRequest) (*DecryptFileResponse, error)
	mustEmbedUnimplementedDecryptServiceServer()
}

// UnimplementedDecryptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDecryptServiceServer struct {
}

func (UnimplementedDecryptServiceServer) DecryptContent(context.Context, *DecryptContentRequest) (*DecryptContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptContent not implemented")
}
func (UnimplementedDecryptServiceServer) DecryptFile(context.Context, *DecryptFileRequest) (*DecryptFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptFile not implemented")
}
func (UnimplementedDecryptServiceServer) mustEmbedUnimplementedDecryptServiceServer() {}

// UnsafeDecryptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecryptServiceServer will
// result in compilation errors.
type UnsafeDecryptServiceServer interface {
	mustEmbedUnimplementedDecryptServiceServer()
}

func RegisterDecryptServiceServer(s grpc.ServiceRegistrar, srv DecryptServiceServer) {
	s.RegisterService(&DecryptService_ServiceDesc, srv)
}

func _DecryptService_DecryptContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptServiceServer).DecryptContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecryptService_DecryptContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptServiceServer).DecryptContent(ctx, req.(*DecryptContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecryptService_DecryptFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptServiceServer).DecryptFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecryptService_DecryptFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptServiceServer).DecryptFile(ctx, req.(*DecryptFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DecryptService_ServiceDesc is the grpc.ServiceDesc for DecryptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecryptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "decrypt.DecryptService",
	HandlerType: (*DecryptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DecryptContent",
			Handler:    _DecryptService_DecryptContent_Handler,
		},
		{
			MethodName: "DecryptFile",
			Handler:    _DecryptService_DecryptFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decrypt.proto",
}