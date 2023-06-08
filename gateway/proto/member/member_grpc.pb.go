//*
// 会员服务

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: member/member.proto

package gateway

import (
	context "context"
	common "gateway/proto/common"
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
	Member_Register_FullMethodName       = "/member.Member/Register"
	Member_Login_FullMethodName          = "/member.Member/Login"
	Member_PhoneLogin_FullMethodName     = "/member.Member/PhoneLogin"
	Member_MemberEdit_FullMethodName     = "/member.Member/MemberEdit"
	Member_MemberPwdEdit_FullMethodName  = "/member.Member/MemberPwdEdit"
	Member_MemberProfile_FullMethodName  = "/member.Member/MemberProfile"
	Member_Logout_FullMethodName         = "/member.Member/Logout"
	Member_AddressAdd_FullMethodName     = "/member.Member/AddressAdd"
	Member_AddressEdit_FullMethodName    = "/member.Member/AddressEdit"
	Member_GetAddressList_FullMethodName = "/member.Member/GetAddressList"
	Member_AddressDel_FullMethodName     = "/member.Member/AddressDel"
	Member_SendCode_FullMethodName       = "/member.Member/SendCode"
)

// MemberClient is the client API for Member service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberClient interface {
	// / 注册
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 用户名密码登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*MemberTokenReply, error)
	// / 手机号登录
	PhoneLogin(ctx context.Context, in *PhoneLoginReq, opts ...grpc.CallOption) (*MemberTokenReply, error)
	// / 修改会员信息
	MemberEdit(ctx context.Context, in *MemberEditReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 修改密码
	MemberPwdEdit(ctx context.Context, in *PwdEditReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 获取会员信息
	MemberProfile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MemberReply, error)
	// / 登出
	Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 添加收货地址
	AddressAdd(ctx context.Context, in *AddressAddReq, opts ...grpc.CallOption) (*AddressIDReply, error)
	// / 修改收货地址
	AddressEdit(ctx context.Context, in *common.Address, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 收货地址列表
	GetAddressList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AddressReply, error)
	// / 删除收货地址
	AddressDel(ctx context.Context, in *AddressIDReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// / 发送短信验证码
	SendCode(ctx context.Context, in *PhoneReq, opts ...grpc.CallOption) (*CodeReply, error)
}

type memberClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberClient(cc grpc.ClientConnInterface) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*MemberTokenReply, error) {
	out := new(MemberTokenReply)
	err := c.cc.Invoke(ctx, Member_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) PhoneLogin(ctx context.Context, in *PhoneLoginReq, opts ...grpc.CallOption) (*MemberTokenReply, error) {
	out := new(MemberTokenReply)
	err := c.cc.Invoke(ctx, Member_PhoneLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) MemberEdit(ctx context.Context, in *MemberEditReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_MemberEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) MemberPwdEdit(ctx context.Context, in *PwdEditReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_MemberPwdEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) MemberProfile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MemberReply, error) {
	out := new(MemberReply)
	err := c.cc.Invoke(ctx, Member_MemberProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) AddressAdd(ctx context.Context, in *AddressAddReq, opts ...grpc.CallOption) (*AddressIDReply, error) {
	out := new(AddressIDReply)
	err := c.cc.Invoke(ctx, Member_AddressAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) AddressEdit(ctx context.Context, in *common.Address, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_AddressEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) GetAddressList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AddressReply, error) {
	out := new(AddressReply)
	err := c.cc.Invoke(ctx, Member_GetAddressList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) AddressDel(ctx context.Context, in *AddressIDReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Member_AddressDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) SendCode(ctx context.Context, in *PhoneReq, opts ...grpc.CallOption) (*CodeReply, error) {
	out := new(CodeReply)
	err := c.cc.Invoke(ctx, Member_SendCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServer is the server API for Member service.
// All implementations must embed UnimplementedMemberServer
// for forward compatibility
type MemberServer interface {
	// / 注册
	Register(context.Context, *RegisterReq) (*emptypb.Empty, error)
	// / 用户名密码登录
	Login(context.Context, *LoginReq) (*MemberTokenReply, error)
	// / 手机号登录
	PhoneLogin(context.Context, *PhoneLoginReq) (*MemberTokenReply, error)
	// / 修改会员信息
	MemberEdit(context.Context, *MemberEditReq) (*emptypb.Empty, error)
	// / 修改密码
	MemberPwdEdit(context.Context, *PwdEditReq) (*emptypb.Empty, error)
	// / 获取会员信息
	MemberProfile(context.Context, *emptypb.Empty) (*MemberReply, error)
	// / 登出
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// / 添加收货地址
	AddressAdd(context.Context, *AddressAddReq) (*AddressIDReply, error)
	// / 修改收货地址
	AddressEdit(context.Context, *common.Address) (*emptypb.Empty, error)
	// / 收货地址列表
	GetAddressList(context.Context, *emptypb.Empty) (*AddressReply, error)
	// / 删除收货地址
	AddressDel(context.Context, *AddressIDReq) (*emptypb.Empty, error)
	// / 发送短信验证码
	SendCode(context.Context, *PhoneReq) (*CodeReply, error)
	mustEmbedUnimplementedMemberServer()
}

// UnimplementedMemberServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServer struct {
}

func (UnimplementedMemberServer) Register(context.Context, *RegisterReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMemberServer) Login(context.Context, *LoginReq) (*MemberTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMemberServer) PhoneLogin(context.Context, *PhoneLoginReq) (*MemberTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneLogin not implemented")
}
func (UnimplementedMemberServer) MemberEdit(context.Context, *MemberEditReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberEdit not implemented")
}
func (UnimplementedMemberServer) MemberPwdEdit(context.Context, *PwdEditReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberPwdEdit not implemented")
}
func (UnimplementedMemberServer) MemberProfile(context.Context, *emptypb.Empty) (*MemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberProfile not implemented")
}
func (UnimplementedMemberServer) Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedMemberServer) AddressAdd(context.Context, *AddressAddReq) (*AddressIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressAdd not implemented")
}
func (UnimplementedMemberServer) AddressEdit(context.Context, *common.Address) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressEdit not implemented")
}
func (UnimplementedMemberServer) GetAddressList(context.Context, *emptypb.Empty) (*AddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressList not implemented")
}
func (UnimplementedMemberServer) AddressDel(context.Context, *AddressIDReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressDel not implemented")
}
func (UnimplementedMemberServer) SendCode(context.Context, *PhoneReq) (*CodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedMemberServer) mustEmbedUnimplementedMemberServer() {}

// UnsafeMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServer will
// result in compilation errors.
type UnsafeMemberServer interface {
	mustEmbedUnimplementedMemberServer()
}

func RegisterMemberServer(s grpc.ServiceRegistrar, srv MemberServer) {
	s.RegisterService(&Member_ServiceDesc, srv)
}

func _Member_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_PhoneLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).PhoneLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_PhoneLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).PhoneLogin(ctx, req.(*PhoneLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_MemberEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberEditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).MemberEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_MemberEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).MemberEdit(ctx, req.(*MemberEditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_MemberPwdEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PwdEditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).MemberPwdEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_MemberPwdEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).MemberPwdEdit(ctx, req.(*PwdEditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_MemberProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).MemberProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_MemberProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).MemberProfile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).Logout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_AddressAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).AddressAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_AddressAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).AddressAdd(ctx, req.(*AddressAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_AddressEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).AddressEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_AddressEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).AddressEdit(ctx, req.(*common.Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_GetAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).GetAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_GetAddressList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).GetAddressList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_AddressDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).AddressDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_AddressDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).AddressDel(ctx, req.(*AddressIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Member_SendCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).SendCode(ctx, req.(*PhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Member_ServiceDesc is the grpc.ServiceDesc for Member service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Member_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "member.Member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Member_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Member_Login_Handler,
		},
		{
			MethodName: "PhoneLogin",
			Handler:    _Member_PhoneLogin_Handler,
		},
		{
			MethodName: "MemberEdit",
			Handler:    _Member_MemberEdit_Handler,
		},
		{
			MethodName: "MemberPwdEdit",
			Handler:    _Member_MemberPwdEdit_Handler,
		},
		{
			MethodName: "MemberProfile",
			Handler:    _Member_MemberProfile_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Member_Logout_Handler,
		},
		{
			MethodName: "AddressAdd",
			Handler:    _Member_AddressAdd_Handler,
		},
		{
			MethodName: "AddressEdit",
			Handler:    _Member_AddressEdit_Handler,
		},
		{
			MethodName: "GetAddressList",
			Handler:    _Member_GetAddressList_Handler,
		},
		{
			MethodName: "AddressDel",
			Handler:    _Member_AddressDel_Handler,
		},
		{
			MethodName: "SendCode",
			Handler:    _Member_SendCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "member/member.proto",
}
