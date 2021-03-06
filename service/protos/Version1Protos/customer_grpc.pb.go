// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: service/protos/Version1Protos/customer.proto

package protos

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	AddCustomerDetails(ctx context.Context, in *CustomerMessageRequest, opts ...grpc.CallOption) (*CustomerMessageReply, error)
	GetCustomers(ctx context.Context, in *GetCustomerMessageRequest, opts ...grpc.CallOption) (*CustomerDetailsReply, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) AddCustomerDetails(ctx context.Context, in *CustomerMessageRequest, opts ...grpc.CallOption) (*CustomerMessageReply, error) {
	out := new(CustomerMessageReply)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/AddCustomerDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *GetCustomerMessageRequest, opts ...grpc.CallOption) (*CustomerDetailsReply, error) {
	out := new(CustomerDetailsReply)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/GetCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	AddCustomerDetails(context.Context, *CustomerMessageRequest) (*CustomerMessageReply, error)
	GetCustomers(context.Context, *GetCustomerMessageRequest) (*CustomerDetailsReply, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) AddCustomerDetails(context.Context, *CustomerMessageRequest) (*CustomerMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomerDetails not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomers(context.Context, *GetCustomerMessageRequest) (*CustomerDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_AddCustomerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).AddCustomerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/AddCustomerDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).AddCustomerDetails(ctx, req.(*CustomerMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/GetCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomers(ctx, req.(*GetCustomerMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCustomerDetails",
			Handler:    _CustomerService_AddCustomerDetails_Handler,
		},
		{
			MethodName: "GetCustomers",
			Handler:    _CustomerService_GetCustomers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/protos/Version1Protos/customer.proto",
}
