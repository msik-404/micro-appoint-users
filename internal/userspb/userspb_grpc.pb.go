// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: userspb.proto

package userspb

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	FindOneCustomerCredentials(ctx context.Context, in *CustomerCredentialsRequest, opts ...grpc.CallOption) (*CredentialsReply, error)
	FindOneOwnerCredentials(ctx context.Context, in *OwnerCredentialsRequest, opts ...grpc.CallOption) (*CredentialsReply, error)
	FindManyOwnerCompanies(ctx context.Context, in *OwnerCompaniesRequest, opts ...grpc.CallOption) (*OwnerCompaniesReply, error)
	AddCustomer(ctx context.Context, in *AddCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddOwner(ctx context.Context, in *AddOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddOwnerCompany(ctx context.Context, in *AddOwnerCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteOwnerCompany(ctx context.Context, in *DeleteOwnerCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateOwner(ctx context.Context, in *UpdateOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteOwner(ctx context.Context, in *DeleteOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) FindOneCustomerCredentials(ctx context.Context, in *CustomerCredentialsRequest, opts ...grpc.CallOption) (*CredentialsReply, error) {
	out := new(CredentialsReply)
	err := c.cc.Invoke(ctx, "/userspb.Api/FindOneCustomerCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) FindOneOwnerCredentials(ctx context.Context, in *OwnerCredentialsRequest, opts ...grpc.CallOption) (*CredentialsReply, error) {
	out := new(CredentialsReply)
	err := c.cc.Invoke(ctx, "/userspb.Api/FindOneOwnerCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) FindManyOwnerCompanies(ctx context.Context, in *OwnerCompaniesRequest, opts ...grpc.CallOption) (*OwnerCompaniesReply, error) {
	out := new(OwnerCompaniesReply)
	err := c.cc.Invoke(ctx, "/userspb.Api/FindManyOwnerCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddCustomer(ctx context.Context, in *AddCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/AddCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddOwner(ctx context.Context, in *AddOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/AddOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddOwnerCompany(ctx context.Context, in *AddOwnerCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/AddOwnerCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteOwnerCompany(ctx context.Context, in *DeleteOwnerCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/DeleteOwnerCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateOwner(ctx context.Context, in *UpdateOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/UpdateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteOwner(ctx context.Context, in *DeleteOwnerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/userspb.Api/DeleteOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	FindOneCustomerCredentials(context.Context, *CustomerCredentialsRequest) (*CredentialsReply, error)
	FindOneOwnerCredentials(context.Context, *OwnerCredentialsRequest) (*CredentialsReply, error)
	FindManyOwnerCompanies(context.Context, *OwnerCompaniesRequest) (*OwnerCompaniesReply, error)
	AddCustomer(context.Context, *AddCustomerRequest) (*emptypb.Empty, error)
	AddOwner(context.Context, *AddOwnerRequest) (*emptypb.Empty, error)
	AddOwnerCompany(context.Context, *AddOwnerCompanyRequest) (*emptypb.Empty, error)
	DeleteOwnerCompany(context.Context, *DeleteOwnerCompanyRequest) (*emptypb.Empty, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*emptypb.Empty, error)
	UpdateOwner(context.Context, *UpdateOwnerRequest) (*emptypb.Empty, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*emptypb.Empty, error)
	DeleteOwner(context.Context, *DeleteOwnerRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) FindOneCustomerCredentials(context.Context, *CustomerCredentialsRequest) (*CredentialsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneCustomerCredentials not implemented")
}
func (UnimplementedApiServer) FindOneOwnerCredentials(context.Context, *OwnerCredentialsRequest) (*CredentialsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneOwnerCredentials not implemented")
}
func (UnimplementedApiServer) FindManyOwnerCompanies(context.Context, *OwnerCompaniesRequest) (*OwnerCompaniesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindManyOwnerCompanies not implemented")
}
func (UnimplementedApiServer) AddCustomer(context.Context, *AddCustomerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomer not implemented")
}
func (UnimplementedApiServer) AddOwner(context.Context, *AddOwnerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOwner not implemented")
}
func (UnimplementedApiServer) AddOwnerCompany(context.Context, *AddOwnerCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOwnerCompany not implemented")
}
func (UnimplementedApiServer) DeleteOwnerCompany(context.Context, *DeleteOwnerCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwnerCompany not implemented")
}
func (UnimplementedApiServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedApiServer) UpdateOwner(context.Context, *UpdateOwnerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOwner not implemented")
}
func (UnimplementedApiServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedApiServer) DeleteOwner(context.Context, *DeleteOwnerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwner not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_FindOneCustomerCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).FindOneCustomerCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/FindOneCustomerCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).FindOneCustomerCredentials(ctx, req.(*CustomerCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_FindOneOwnerCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnerCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).FindOneOwnerCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/FindOneOwnerCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).FindOneOwnerCredentials(ctx, req.(*OwnerCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_FindManyOwnerCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnerCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).FindManyOwnerCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/FindManyOwnerCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).FindManyOwnerCompanies(ctx, req.(*OwnerCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/AddCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddCustomer(ctx, req.(*AddCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/AddOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddOwner(ctx, req.(*AddOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddOwnerCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOwnerCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddOwnerCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/AddOwnerCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddOwnerCompany(ctx, req.(*AddOwnerCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteOwnerCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOwnerCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteOwnerCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/DeleteOwnerCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteOwnerCompany(ctx, req.(*DeleteOwnerCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/UpdateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateOwner(ctx, req.(*UpdateOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Api/DeleteOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteOwner(ctx, req.(*DeleteOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userspb.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneCustomerCredentials",
			Handler:    _Api_FindOneCustomerCredentials_Handler,
		},
		{
			MethodName: "FindOneOwnerCredentials",
			Handler:    _Api_FindOneOwnerCredentials_Handler,
		},
		{
			MethodName: "FindManyOwnerCompanies",
			Handler:    _Api_FindManyOwnerCompanies_Handler,
		},
		{
			MethodName: "AddCustomer",
			Handler:    _Api_AddCustomer_Handler,
		},
		{
			MethodName: "AddOwner",
			Handler:    _Api_AddOwner_Handler,
		},
		{
			MethodName: "AddOwnerCompany",
			Handler:    _Api_AddOwnerCompany_Handler,
		},
		{
			MethodName: "DeleteOwnerCompany",
			Handler:    _Api_DeleteOwnerCompany_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _Api_UpdateCustomer_Handler,
		},
		{
			MethodName: "UpdateOwner",
			Handler:    _Api_UpdateOwner_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _Api_DeleteCustomer_Handler,
		},
		{
			MethodName: "DeleteOwner",
			Handler:    _Api_DeleteOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userspb.proto",
}