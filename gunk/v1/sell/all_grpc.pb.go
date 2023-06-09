// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sellpb

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

// SellServiceClient is the client API for SellService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellServiceClient interface {
	CreateSell(ctx context.Context, in *CreateSellRequest, opts ...grpc.CallOption) (*CreateSellResponse, error)
	ListSell(ctx context.Context, in *ListSellRequest, opts ...grpc.CallOption) (*ListSellResponse, error)
	GetSell(ctx context.Context, in *GetSellRequest, opts ...grpc.CallOption) (*GetSellResponse, error)
}

type sellServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSellServiceClient(cc grpc.ClientConnInterface) SellServiceClient {
	return &sellServiceClient{cc}
}

func (c *sellServiceClient) CreateSell(ctx context.Context, in *CreateSellRequest, opts ...grpc.CallOption) (*CreateSellResponse, error) {
	out := new(CreateSellResponse)
	err := c.cc.Invoke(ctx, "/sellpb.SellService/CreateSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellServiceClient) ListSell(ctx context.Context, in *ListSellRequest, opts ...grpc.CallOption) (*ListSellResponse, error) {
	out := new(ListSellResponse)
	err := c.cc.Invoke(ctx, "/sellpb.SellService/ListSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellServiceClient) GetSell(ctx context.Context, in *GetSellRequest, opts ...grpc.CallOption) (*GetSellResponse, error) {
	out := new(GetSellResponse)
	err := c.cc.Invoke(ctx, "/sellpb.SellService/GetSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellServiceServer is the server API for SellService service.
// All implementations must embed UnimplementedSellServiceServer
// for forward compatibility
type SellServiceServer interface {
	CreateSell(context.Context, *CreateSellRequest) (*CreateSellResponse, error)
	ListSell(context.Context, *ListSellRequest) (*ListSellResponse, error)
	GetSell(context.Context, *GetSellRequest) (*GetSellResponse, error)
	mustEmbedUnimplementedSellServiceServer()
}

// UnimplementedSellServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSellServiceServer struct {
}

func (UnimplementedSellServiceServer) CreateSell(context.Context, *CreateSellRequest) (*CreateSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSell not implemented")
}
func (UnimplementedSellServiceServer) ListSell(context.Context, *ListSellRequest) (*ListSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSell not implemented")
}
func (UnimplementedSellServiceServer) GetSell(context.Context, *GetSellRequest) (*GetSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSell not implemented")
}
func (UnimplementedSellServiceServer) mustEmbedUnimplementedSellServiceServer() {}

// UnsafeSellServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellServiceServer will
// result in compilation errors.
type UnsafeSellServiceServer interface {
	mustEmbedUnimplementedSellServiceServer()
}

func RegisterSellServiceServer(s grpc.ServiceRegistrar, srv SellServiceServer) {
	s.RegisterService(&SellService_ServiceDesc, srv)
}

func _SellService_CreateSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellServiceServer).CreateSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sellpb.SellService/CreateSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellServiceServer).CreateSell(ctx, req.(*CreateSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellService_ListSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellServiceServer).ListSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sellpb.SellService/ListSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellServiceServer).ListSell(ctx, req.(*ListSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellService_GetSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellServiceServer).GetSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sellpb.SellService/GetSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellServiceServer).GetSell(ctx, req.(*GetSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SellService_ServiceDesc is the grpc.ServiceDesc for SellService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SellService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sellpb.SellService",
	HandlerType: (*SellServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSell",
			Handler:    _SellService_CreateSell_Handler,
		},
		{
			MethodName: "ListSell",
			Handler:    _SellService_ListSell_Handler,
		},
		{
			MethodName: "GetSell",
			Handler:    _SellService_GetSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "codemen.org/inventory/gunk/v1/sell/all.proto",
}
