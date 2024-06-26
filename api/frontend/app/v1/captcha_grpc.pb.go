// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: v1/captcha.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CaptCha_Get_FullMethodName    = "/api.frontend.app.v1.CaptCha/Get"
	CaptCha_Verify_FullMethodName = "/api.frontend.app.v1.CaptCha/Verify"
)

// CaptChaClient is the client API for CaptCha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptChaClient interface {
	// 获取验证码
	Get(ctx context.Context, in *GetCaptcha, opts ...grpc.CallOption) (*GetCaptchaReply, error)
	// 验证验证码
	Verify(ctx context.Context, in *VerifyCaptcha, opts ...grpc.CallOption) (*VerifyCaptchaReply, error)
}

type captChaClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptChaClient(cc grpc.ClientConnInterface) CaptChaClient {
	return &captChaClient{cc}
}

func (c *captChaClient) Get(ctx context.Context, in *GetCaptcha, opts ...grpc.CallOption) (*GetCaptchaReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCaptchaReply)
	err := c.cc.Invoke(ctx, CaptCha_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captChaClient) Verify(ctx context.Context, in *VerifyCaptcha, opts ...grpc.CallOption) (*VerifyCaptchaReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyCaptchaReply)
	err := c.cc.Invoke(ctx, CaptCha_Verify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptChaServer is the server API for CaptCha service.
// All implementations must embed UnimplementedCaptChaServer
// for forward compatibility
type CaptChaServer interface {
	// 获取验证码
	Get(context.Context, *GetCaptcha) (*GetCaptchaReply, error)
	// 验证验证码
	Verify(context.Context, *VerifyCaptcha) (*VerifyCaptchaReply, error)
	mustEmbedUnimplementedCaptChaServer()
}

// UnimplementedCaptChaServer must be embedded to have forward compatible implementations.
type UnimplementedCaptChaServer struct {
}

func (UnimplementedCaptChaServer) Get(context.Context, *GetCaptcha) (*GetCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCaptChaServer) Verify(context.Context, *VerifyCaptcha) (*VerifyCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedCaptChaServer) mustEmbedUnimplementedCaptChaServer() {}

// UnsafeCaptChaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptChaServer will
// result in compilation errors.
type UnsafeCaptChaServer interface {
	mustEmbedUnimplementedCaptChaServer()
}

func RegisterCaptChaServer(s grpc.ServiceRegistrar, srv CaptChaServer) {
	s.RegisterService(&CaptCha_ServiceDesc, srv)
}

func _CaptCha_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptcha)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptChaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptCha_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptChaServer).Get(ctx, req.(*GetCaptcha))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptCha_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptcha)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptChaServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptCha_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptChaServer).Verify(ctx, req.(*VerifyCaptcha))
	}
	return interceptor(ctx, in, info, handler)
}

// CaptCha_ServiceDesc is the grpc.ServiceDesc for CaptCha service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaptCha_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.frontend.app.v1.CaptCha",
	HandlerType: (*CaptChaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CaptCha_Get_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _CaptCha_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/captcha.proto",
}
