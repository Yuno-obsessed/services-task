// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: services.proto

package servicespb

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

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	Provide(ctx context.Context, in *SymbolsRequest, opts ...grpc.CallOption) (*SymbolsResponse, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) Provide(ctx context.Context, in *SymbolsRequest, opts ...grpc.CallOption) (*SymbolsResponse, error) {
	out := new(SymbolsResponse)
	err := c.cc.Invoke(ctx, "/pkg.Provider/Provide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	Provide(context.Context, *SymbolsRequest) (*SymbolsResponse, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) Provide(context.Context, *SymbolsRequest) (*SymbolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provide not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_Provide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Provide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.Provider/Provide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Provide(ctx, req.(*SymbolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provide",
			Handler:    _Provider_Provide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// ReceiverClient is the client API for Receiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiverClient interface {
	Receive(ctx context.Context, in *SymbolsResponse, opts ...grpc.CallOption) (*ProcessedSymbols, error)
}

type receiverClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiverClient(cc grpc.ClientConnInterface) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) Receive(ctx context.Context, in *SymbolsResponse, opts ...grpc.CallOption) (*ProcessedSymbols, error) {
	out := new(ProcessedSymbols)
	err := c.cc.Invoke(ctx, "/pkg.Receiver/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverServer is the server API for Receiver service.
// All implementations must embed UnimplementedReceiverServer
// for forward compatibility
type ReceiverServer interface {
	Receive(context.Context, *SymbolsResponse) (*ProcessedSymbols, error)
	mustEmbedUnimplementedReceiverServer()
}

// UnimplementedReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedReceiverServer struct {
}

func (UnimplementedReceiverServer) Receive(context.Context, *SymbolsResponse) (*ProcessedSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedReceiverServer) mustEmbedUnimplementedReceiverServer() {}

// UnsafeReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiverServer will
// result in compilation errors.
type UnsafeReceiverServer interface {
	mustEmbedUnimplementedReceiverServer()
}

func RegisterReceiverServer(s grpc.ServiceRegistrar, srv ReceiverServer) {
	s.RegisterService(&Receiver_ServiceDesc, srv)
}

func _Receiver_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolsResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.Receiver/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).Receive(ctx, req.(*SymbolsResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// Receiver_ServiceDesc is the grpc.ServiceDesc for Receiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Receiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _Receiver_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// VisualizerClient is the client API for Visualizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VisualizerClient interface {
	Visualize(ctx context.Context, in *ProcessedSymbols, opts ...grpc.CallOption) (*PresentedSymbols, error)
}

type visualizerClient struct {
	cc grpc.ClientConnInterface
}

func NewVisualizerClient(cc grpc.ClientConnInterface) VisualizerClient {
	return &visualizerClient{cc}
}

func (c *visualizerClient) Visualize(ctx context.Context, in *ProcessedSymbols, opts ...grpc.CallOption) (*PresentedSymbols, error) {
	out := new(PresentedSymbols)
	err := c.cc.Invoke(ctx, "/pkg.Visualizer/Visualize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisualizerServer is the server API for Visualizer service.
// All implementations must embed UnimplementedVisualizerServer
// for forward compatibility
type VisualizerServer interface {
	Visualize(context.Context, *ProcessedSymbols) (*PresentedSymbols, error)
	mustEmbedUnimplementedVisualizerServer()
}

// UnimplementedVisualizerServer must be embedded to have forward compatible implementations.
type UnimplementedVisualizerServer struct {
}

func (UnimplementedVisualizerServer) Visualize(context.Context, *ProcessedSymbols) (*PresentedSymbols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Visualize not implemented")
}
func (UnimplementedVisualizerServer) mustEmbedUnimplementedVisualizerServer() {}

// UnsafeVisualizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VisualizerServer will
// result in compilation errors.
type UnsafeVisualizerServer interface {
	mustEmbedUnimplementedVisualizerServer()
}

func RegisterVisualizerServer(s grpc.ServiceRegistrar, srv VisualizerServer) {
	s.RegisterService(&Visualizer_ServiceDesc, srv)
}

func _Visualizer_Visualize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessedSymbols)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisualizerServer).Visualize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.Visualizer/Visualize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisualizerServer).Visualize(ctx, req.(*ProcessedSymbols))
	}
	return interceptor(ctx, in, info, handler)
}

// Visualizer_ServiceDesc is the grpc.ServiceDesc for Visualizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Visualizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.Visualizer",
	HandlerType: (*VisualizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Visualize",
			Handler:    _Visualizer_Visualize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}