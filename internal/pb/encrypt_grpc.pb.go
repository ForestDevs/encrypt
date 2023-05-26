// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: encrypt.proto

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
	EncryptService_EncryptContent_FullMethodName = "/encrypt.EncryptService/EncryptContent"
	EncryptService_EncryptFile_FullMethodName    = "/encrypt.EncryptService/EncryptFile"
)

// EncryptServiceClient is the client API for EncryptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncryptServiceClient interface {
	EncryptContent(ctx context.Context, in *EncryptContentRequest, opts ...grpc.CallOption) (*EncryptContentResponse, error)
	EncryptFile(ctx context.Context, in *EncryptFileRequest, opts ...grpc.CallOption) (*EncryptFileResponse, error)
}

type encryptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncryptServiceClient(cc grpc.ClientConnInterface) EncryptServiceClient {
	return &encryptServiceClient{cc}
}

func (c *encryptServiceClient) EncryptContent(ctx context.Context, in *EncryptContentRequest, opts ...grpc.CallOption) (*EncryptContentResponse, error) {
	out := new(EncryptContentResponse)
	err := c.cc.Invoke(ctx, EncryptService_EncryptContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptServiceClient) EncryptFile(ctx context.Context, in *EncryptFileRequest, opts ...grpc.CallOption) (*EncryptFileResponse, error) {
	out := new(EncryptFileResponse)
	err := c.cc.Invoke(ctx, EncryptService_EncryptFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptServiceServer is the server API for EncryptService service.
// All implementations must embed UnimplementedEncryptServiceServer
// for forward compatibility
type EncryptServiceServer interface {
	EncryptContent(context.Context, *EncryptContentRequest) (*EncryptContentResponse, error)
	EncryptFile(context.Context, *EncryptFileRequest) (*EncryptFileResponse, error)
	mustEmbedUnimplementedEncryptServiceServer()
}

// UnimplementedEncryptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncryptServiceServer struct {
}

func (UnimplementedEncryptServiceServer) EncryptContent(context.Context, *EncryptContentRequest) (*EncryptContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptContent not implemented")
}
func (UnimplementedEncryptServiceServer) EncryptFile(context.Context, *EncryptFileRequest) (*EncryptFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptFile not implemented")
}
func (UnimplementedEncryptServiceServer) mustEmbedUnimplementedEncryptServiceServer() {}

// UnsafeEncryptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncryptServiceServer will
// result in compilation errors.
type UnsafeEncryptServiceServer interface {
	mustEmbedUnimplementedEncryptServiceServer()
}

func RegisterEncryptServiceServer(s grpc.ServiceRegistrar, srv EncryptServiceServer) {
	s.RegisterService(&EncryptService_ServiceDesc, srv)
}

func _EncryptService_EncryptContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptServiceServer).EncryptContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncryptService_EncryptContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptServiceServer).EncryptContent(ctx, req.(*EncryptContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncryptService_EncryptFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptServiceServer).EncryptFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncryptService_EncryptFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptServiceServer).EncryptFile(ctx, req.(*EncryptFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncryptService_ServiceDesc is the grpc.ServiceDesc for EncryptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncryptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "encrypt.EncryptService",
	HandlerType: (*EncryptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncryptContent",
			Handler:    _EncryptService_EncryptContent_Handler,
		},
		{
			MethodName: "EncryptFile",
			Handler:    _EncryptService_EncryptFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encrypt.proto",
}