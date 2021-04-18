// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cybervector

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

// CyberVectorProxyServiceClient is the client API for CyberVectorProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CyberVectorProxyServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (CyberVectorProxyService_SubscribeClient, error)
	UnSubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*ProxyMessaage, error)
	InsertIntent(ctx context.Context, in *InsertIntentRequest, opts ...grpc.CallOption) (*InsertIntentResponse, error)
	SelectIntents(ctx context.Context, in *SelectIntentRequest, opts ...grpc.CallOption) (*SelectIntentResponse, error)
	DeleteIntent(ctx context.Context, in *DeleteIntentRequest, opts ...grpc.CallOption) (*DeleteIntentResponse, error)
}

type cyberVectorProxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCyberVectorProxyServiceClient(cc grpc.ClientConnInterface) CyberVectorProxyServiceClient {
	return &cyberVectorProxyServiceClient{cc}
}

func (c *cyberVectorProxyServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (CyberVectorProxyService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CyberVectorProxyService_ServiceDesc.Streams[0], "/cybervector.CyberVectorProxyService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &cyberVectorProxyServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CyberVectorProxyService_SubscribeClient interface {
	Recv() (*ProxyMessaage, error)
	grpc.ClientStream
}

type cyberVectorProxyServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *cyberVectorProxyServiceSubscribeClient) Recv() (*ProxyMessaage, error) {
	m := new(ProxyMessaage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cyberVectorProxyServiceClient) UnSubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*ProxyMessaage, error) {
	out := new(ProxyMessaage)
	err := c.cc.Invoke(ctx, "/cybervector.CyberVectorProxyService/UnSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberVectorProxyServiceClient) InsertIntent(ctx context.Context, in *InsertIntentRequest, opts ...grpc.CallOption) (*InsertIntentResponse, error) {
	out := new(InsertIntentResponse)
	err := c.cc.Invoke(ctx, "/cybervector.CyberVectorProxyService/InsertIntent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberVectorProxyServiceClient) SelectIntents(ctx context.Context, in *SelectIntentRequest, opts ...grpc.CallOption) (*SelectIntentResponse, error) {
	out := new(SelectIntentResponse)
	err := c.cc.Invoke(ctx, "/cybervector.CyberVectorProxyService/SelectIntents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cyberVectorProxyServiceClient) DeleteIntent(ctx context.Context, in *DeleteIntentRequest, opts ...grpc.CallOption) (*DeleteIntentResponse, error) {
	out := new(DeleteIntentResponse)
	err := c.cc.Invoke(ctx, "/cybervector.CyberVectorProxyService/DeleteIntent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CyberVectorProxyServiceServer is the server API for CyberVectorProxyService service.
// All implementations must embed UnimplementedCyberVectorProxyServiceServer
// for forward compatibility
type CyberVectorProxyServiceServer interface {
	Subscribe(*SubscribeRequest, CyberVectorProxyService_SubscribeServer) error
	UnSubscribe(context.Context, *UnsubscribeRequest) (*ProxyMessaage, error)
	InsertIntent(context.Context, *InsertIntentRequest) (*InsertIntentResponse, error)
	SelectIntents(context.Context, *SelectIntentRequest) (*SelectIntentResponse, error)
	DeleteIntent(context.Context, *DeleteIntentRequest) (*DeleteIntentResponse, error)
	mustEmbedUnimplementedCyberVectorProxyServiceServer()
}

// UnimplementedCyberVectorProxyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCyberVectorProxyServiceServer struct {
}

func (UnimplementedCyberVectorProxyServiceServer) Subscribe(*SubscribeRequest, CyberVectorProxyService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedCyberVectorProxyServiceServer) UnSubscribe(context.Context, *UnsubscribeRequest) (*ProxyMessaage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribe not implemented")
}
func (UnimplementedCyberVectorProxyServiceServer) InsertIntent(context.Context, *InsertIntentRequest) (*InsertIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertIntent not implemented")
}
func (UnimplementedCyberVectorProxyServiceServer) SelectIntents(context.Context, *SelectIntentRequest) (*SelectIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectIntents not implemented")
}
func (UnimplementedCyberVectorProxyServiceServer) DeleteIntent(context.Context, *DeleteIntentRequest) (*DeleteIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIntent not implemented")
}
func (UnimplementedCyberVectorProxyServiceServer) mustEmbedUnimplementedCyberVectorProxyServiceServer() {
}

// UnsafeCyberVectorProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CyberVectorProxyServiceServer will
// result in compilation errors.
type UnsafeCyberVectorProxyServiceServer interface {
	mustEmbedUnimplementedCyberVectorProxyServiceServer()
}

func RegisterCyberVectorProxyServiceServer(s grpc.ServiceRegistrar, srv CyberVectorProxyServiceServer) {
	s.RegisterService(&CyberVectorProxyService_ServiceDesc, srv)
}

func _CyberVectorProxyService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CyberVectorProxyServiceServer).Subscribe(m, &cyberVectorProxyServiceSubscribeServer{stream})
}

type CyberVectorProxyService_SubscribeServer interface {
	Send(*ProxyMessaage) error
	grpc.ServerStream
}

type cyberVectorProxyServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *cyberVectorProxyServiceSubscribeServer) Send(m *ProxyMessaage) error {
	return x.ServerStream.SendMsg(m)
}

func _CyberVectorProxyService_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberVectorProxyServiceServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cybervector.CyberVectorProxyService/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberVectorProxyServiceServer).UnSubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CyberVectorProxyService_InsertIntent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertIntentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberVectorProxyServiceServer).InsertIntent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cybervector.CyberVectorProxyService/InsertIntent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberVectorProxyServiceServer).InsertIntent(ctx, req.(*InsertIntentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CyberVectorProxyService_SelectIntents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectIntentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberVectorProxyServiceServer).SelectIntents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cybervector.CyberVectorProxyService/SelectIntents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberVectorProxyServiceServer).SelectIntents(ctx, req.(*SelectIntentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CyberVectorProxyService_DeleteIntent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIntentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CyberVectorProxyServiceServer).DeleteIntent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cybervector.CyberVectorProxyService/DeleteIntent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CyberVectorProxyServiceServer).DeleteIntent(ctx, req.(*DeleteIntentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CyberVectorProxyService_ServiceDesc is the grpc.ServiceDesc for CyberVectorProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CyberVectorProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cybervector.CyberVectorProxyService",
	HandlerType: (*CyberVectorProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnSubscribe",
			Handler:    _CyberVectorProxyService_UnSubscribe_Handler,
		},
		{
			MethodName: "InsertIntent",
			Handler:    _CyberVectorProxyService_InsertIntent_Handler,
		},
		{
			MethodName: "SelectIntents",
			Handler:    _CyberVectorProxyService_SelectIntents_Handler,
		},
		{
			MethodName: "DeleteIntent",
			Handler:    _CyberVectorProxyService_DeleteIntent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _CyberVectorProxyService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cybervector_proxy.proto",
}