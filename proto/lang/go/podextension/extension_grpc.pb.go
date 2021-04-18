// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package podextension

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

// PodExtensionClient is the client API for PodExtension service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodExtensionClient interface {
	Unary(ctx context.Context, in *UnaryReq, opts ...grpc.CallOption) (*UnaryResp, error)
}

type podExtensionClient struct {
	cc grpc.ClientConnInterface
}

func NewPodExtensionClient(cc grpc.ClientConnInterface) PodExtensionClient {
	return &podExtensionClient{cc}
}

func (c *podExtensionClient) Unary(ctx context.Context, in *UnaryReq, opts ...grpc.CallOption) (*UnaryResp, error) {
	out := new(UnaryResp)
	err := c.cc.Invoke(ctx, "/podextension.PodExtension/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodExtensionServer is the server API for PodExtension service.
// All implementations must embed UnimplementedPodExtensionServer
// for forward compatibility
type PodExtensionServer interface {
	Unary(context.Context, *UnaryReq) (*UnaryResp, error)
	mustEmbedUnimplementedPodExtensionServer()
}

// UnimplementedPodExtensionServer must be embedded to have forward compatible implementations.
type UnimplementedPodExtensionServer struct {
}

func (UnimplementedPodExtensionServer) Unary(context.Context, *UnaryReq) (*UnaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (UnimplementedPodExtensionServer) mustEmbedUnimplementedPodExtensionServer() {}

// UnsafePodExtensionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodExtensionServer will
// result in compilation errors.
type UnsafePodExtensionServer interface {
	mustEmbedUnimplementedPodExtensionServer()
}

func RegisterPodExtensionServer(s grpc.ServiceRegistrar, srv PodExtensionServer) {
	s.RegisterService(&PodExtension_ServiceDesc, srv)
}

func _PodExtension_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodExtensionServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podextension.PodExtension/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodExtensionServer).Unary(ctx, req.(*UnaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PodExtension_ServiceDesc is the grpc.ServiceDesc for PodExtension service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodExtension_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "podextension.PodExtension",
	HandlerType: (*PodExtensionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _PodExtension_Unary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extension.proto",
}
