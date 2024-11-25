// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: service.proto

package textgen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TextGenService_GenerateText_FullMethodName = "/textgen.TextGenService/GenerateText"
)

// TextGenServiceClient is the client API for TextGenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextGenServiceClient interface {
	// Метод для генерации текста
	GenerateText(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type textGenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextGenServiceClient(cc grpc.ClientConnInterface) TextGenServiceClient {
	return &textGenServiceClient{cc}
}

func (c *textGenServiceClient) GenerateText(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, TextGenService_GenerateText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextGenServiceServer is the server API for TextGenService service.
// All implementations must embed UnimplementedTextGenServiceServer
// for forward compatibility.
type TextGenServiceServer interface {
	// Метод для генерации текста
	GenerateText(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedTextGenServiceServer()
}

// UnimplementedTextGenServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTextGenServiceServer struct{}

func (UnimplementedTextGenServiceServer) GenerateText(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateText not implemented")
}
func (UnimplementedTextGenServiceServer) mustEmbedUnimplementedTextGenServiceServer() {}
func (UnimplementedTextGenServiceServer) testEmbeddedByValue()                        {}

// UnsafeTextGenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextGenServiceServer will
// result in compilation errors.
type UnsafeTextGenServiceServer interface {
	mustEmbedUnimplementedTextGenServiceServer()
}

func RegisterTextGenServiceServer(s grpc.ServiceRegistrar, srv TextGenServiceServer) {
	// If the following call pancis, it indicates UnimplementedTextGenServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TextGenService_ServiceDesc, srv)
}

func _TextGenService_GenerateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenServiceServer).GenerateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextGenService_GenerateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenServiceServer).GenerateText(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextGenService_ServiceDesc is the grpc.ServiceDesc for TextGenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextGenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "textgen.TextGenService",
	HandlerType: (*TextGenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateText",
			Handler:    _TextGenService_GenerateText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
