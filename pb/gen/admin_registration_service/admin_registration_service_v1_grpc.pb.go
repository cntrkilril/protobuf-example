// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: admin_registration_service_v1.proto

package admin_registration_service

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
	AdminRegistrationServiceV1_Register_FullMethodName = "/admin_registration.AdminRegistrationServiceV1/Register"
)

// AdminRegistrationServiceV1Client is the client API for AdminRegistrationServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminRegistrationServiceV1Client interface {
	Register(ctx context.Context, in *RegisterAdminRequestV1, opts ...grpc.CallOption) (*RegisterAdminResponseV1, error)
}

type adminRegistrationServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAdminRegistrationServiceV1Client(cc grpc.ClientConnInterface) AdminRegistrationServiceV1Client {
	return &adminRegistrationServiceV1Client{cc}
}

func (c *adminRegistrationServiceV1Client) Register(ctx context.Context, in *RegisterAdminRequestV1, opts ...grpc.CallOption) (*RegisterAdminResponseV1, error) {
	out := new(RegisterAdminResponseV1)
	err := c.cc.Invoke(ctx, AdminRegistrationServiceV1_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminRegistrationServiceV1Server is the server API for AdminRegistrationServiceV1 service.
// All implementations must embed UnimplementedAdminRegistrationServiceV1Server
// for forward compatibility
type AdminRegistrationServiceV1Server interface {
	Register(context.Context, *RegisterAdminRequestV1) (*RegisterAdminResponseV1, error)
	mustEmbedUnimplementedAdminRegistrationServiceV1Server()
}

// UnimplementedAdminRegistrationServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedAdminRegistrationServiceV1Server struct {
}

func (UnimplementedAdminRegistrationServiceV1Server) Register(context.Context, *RegisterAdminRequestV1) (*RegisterAdminResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAdminRegistrationServiceV1Server) mustEmbedUnimplementedAdminRegistrationServiceV1Server() {
}

// UnsafeAdminRegistrationServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminRegistrationServiceV1Server will
// result in compilation errors.
type UnsafeAdminRegistrationServiceV1Server interface {
	mustEmbedUnimplementedAdminRegistrationServiceV1Server()
}

func RegisterAdminRegistrationServiceV1Server(s grpc.ServiceRegistrar, srv AdminRegistrationServiceV1Server) {
	s.RegisterService(&AdminRegistrationServiceV1_ServiceDesc, srv)
}

func _AdminRegistrationServiceV1_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAdminRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminRegistrationServiceV1Server).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminRegistrationServiceV1_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminRegistrationServiceV1Server).Register(ctx, req.(*RegisterAdminRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminRegistrationServiceV1_ServiceDesc is the grpc.ServiceDesc for AdminRegistrationServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminRegistrationServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_registration.AdminRegistrationServiceV1",
	HandlerType: (*AdminRegistrationServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AdminRegistrationServiceV1_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin_registration_service_v1.proto",
}
