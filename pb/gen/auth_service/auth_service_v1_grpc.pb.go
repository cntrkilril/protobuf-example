// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: auth_service_v1.proto

package auth_service

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

const (
	AuthServiceV1_AuthenticateAdmin_FullMethodName   = "/auth.AuthServiceV1/AuthenticateAdmin"
	AuthServiceV1_ConfirmAdminSession_FullMethodName = "/auth.AuthServiceV1/ConfirmAdminSession"
)

// AuthServiceV1Client is the client API for AuthServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceV1Client interface {
	AuthenticateAdmin(ctx context.Context, in *AuthenticateAdminRequestV1, opts ...grpc.CallOption) (*AuthenticateAdminResponseV1, error)
	ConfirmAdminSession(ctx context.Context, in *ConfirmAdminSessionRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceV1Client(cc grpc.ClientConnInterface) AuthServiceV1Client {
	return &authServiceV1Client{cc}
}

func (c *authServiceV1Client) AuthenticateAdmin(ctx context.Context, in *AuthenticateAdminRequestV1, opts ...grpc.CallOption) (*AuthenticateAdminResponseV1, error) {
	out := new(AuthenticateAdminResponseV1)
	err := c.cc.Invoke(ctx, AuthServiceV1_AuthenticateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) ConfirmAdminSession(ctx context.Context, in *ConfirmAdminSessionRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthServiceV1_ConfirmAdminSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceV1Server is the server API for AuthServiceV1 service.
// All implementations must embed UnimplementedAuthServiceV1Server
// for forward compatibility
type AuthServiceV1Server interface {
	AuthenticateAdmin(context.Context, *AuthenticateAdminRequestV1) (*AuthenticateAdminResponseV1, error)
	ConfirmAdminSession(context.Context, *ConfirmAdminSessionRequestV1) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServiceV1Server()
}

// UnimplementedAuthServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceV1Server struct {
}

func (UnimplementedAuthServiceV1Server) AuthenticateAdmin(context.Context, *AuthenticateAdminRequestV1) (*AuthenticateAdminResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateAdmin not implemented")
}
func (UnimplementedAuthServiceV1Server) ConfirmAdminSession(context.Context, *ConfirmAdminSessionRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmAdminSession not implemented")
}
func (UnimplementedAuthServiceV1Server) mustEmbedUnimplementedAuthServiceV1Server() {}

// UnsafeAuthServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceV1Server will
// result in compilation errors.
type UnsafeAuthServiceV1Server interface {
	mustEmbedUnimplementedAuthServiceV1Server()
}

func RegisterAuthServiceV1Server(s grpc.ServiceRegistrar, srv AuthServiceV1Server) {
	s.RegisterService(&AuthServiceV1_ServiceDesc, srv)
}

func _AuthServiceV1_AuthenticateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateAdminRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).AuthenticateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceV1_AuthenticateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).AuthenticateAdmin(ctx, req.(*AuthenticateAdminRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_ConfirmAdminSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmAdminSessionRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).ConfirmAdminSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthServiceV1_ConfirmAdminSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).ConfirmAdminSession(ctx, req.(*ConfirmAdminSessionRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServiceV1_ServiceDesc is the grpc.ServiceDesc for AuthServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthServiceV1",
	HandlerType: (*AuthServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateAdmin",
			Handler:    _AuthServiceV1_AuthenticateAdmin_Handler,
		},
		{
			MethodName: "ConfirmAdminSession",
			Handler:    _AuthServiceV1_ConfirmAdminSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service_v1.proto",
}
