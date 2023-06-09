//*
// 营销服务

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: market/market.proto

package market

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
	Market_GetHomeData_FullMethodName     = "/market.Market/GetHomeData"
	Market_GetHomeCatData_FullMethodName  = "/market.Market/GetHomeCatData"
	Market_GetNoticeList_FullMethodName   = "/market.Market/GetNoticeList"
	Market_GetSearchData_FullMethodName   = "/market.Market/GetSearchData"
	Market_GetPayConfig_FullMethodName    = "/market.Market/GetPayConfig"
	Market_GetCouponList_FullMethodName   = "/market.Market/GetCouponList"
	Market_GetMyCouponList_FullMethodName = "/market.Market/GetMyCouponList"
	Market_CouponDraw_FullMethodName      = "/market.Market/CouponDraw"
)

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	// / 获取首页配置数据
	GetHomeData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HomeDataReply, error)
	// / 获取首页分类下配置数据
	GetHomeCatData(ctx context.Context, in *CatReq, opts ...grpc.CallOption) (*AppSettingReply, error)
	// / 获取公告列表
	GetNoticeList(ctx context.Context, in *common.PageReq, opts ...grpc.CallOption) (*NoticeReply, error)
	// / 获取搜索页配置数据
	GetSearchData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchReply, error)
	// / 获取支付配置
	GetPayConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PayReply, error)
	// / 商品可以领取的优惠券列表
	GetCouponList(ctx context.Context, in *common.SkuIDReq, opts ...grpc.CallOption) (*CouponListReply, error)
	// / 我的优惠券列表
	GetMyCouponList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CouponListReply, error)
	// / 领取优惠券
	CouponDraw(ctx context.Context, in *CouponReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetHomeData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HomeDataReply, error) {
	out := new(HomeDataReply)
	err := c.cc.Invoke(ctx, Market_GetHomeData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetHomeCatData(ctx context.Context, in *CatReq, opts ...grpc.CallOption) (*AppSettingReply, error) {
	out := new(AppSettingReply)
	err := c.cc.Invoke(ctx, Market_GetHomeCatData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetNoticeList(ctx context.Context, in *common.PageReq, opts ...grpc.CallOption) (*NoticeReply, error) {
	out := new(NoticeReply)
	err := c.cc.Invoke(ctx, Market_GetNoticeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetSearchData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, Market_GetSearchData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetPayConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, Market_GetPayConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetCouponList(ctx context.Context, in *common.SkuIDReq, opts ...grpc.CallOption) (*CouponListReply, error) {
	out := new(CouponListReply)
	err := c.cc.Invoke(ctx, Market_GetCouponList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetMyCouponList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CouponListReply, error) {
	out := new(CouponListReply)
	err := c.cc.Invoke(ctx, Market_GetMyCouponList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CouponDraw(ctx context.Context, in *CouponReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Market_CouponDraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	// / 获取首页配置数据
	GetHomeData(context.Context, *emptypb.Empty) (*HomeDataReply, error)
	// / 获取首页分类下配置数据
	GetHomeCatData(context.Context, *CatReq) (*AppSettingReply, error)
	// / 获取公告列表
	GetNoticeList(context.Context, *common.PageReq) (*NoticeReply, error)
	// / 获取搜索页配置数据
	GetSearchData(context.Context, *emptypb.Empty) (*SearchReply, error)
	// / 获取支付配置
	GetPayConfig(context.Context, *emptypb.Empty) (*PayReply, error)
	// / 商品可以领取的优惠券列表
	GetCouponList(context.Context, *common.SkuIDReq) (*CouponListReply, error)
	// / 我的优惠券列表
	GetMyCouponList(context.Context, *emptypb.Empty) (*CouponListReply, error)
	// / 领取优惠券
	CouponDraw(context.Context, *CouponReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) GetHomeData(context.Context, *emptypb.Empty) (*HomeDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeData not implemented")
}
func (UnimplementedMarketServer) GetHomeCatData(context.Context, *CatReq) (*AppSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeCatData not implemented")
}
func (UnimplementedMarketServer) GetNoticeList(context.Context, *common.PageReq) (*NoticeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoticeList not implemented")
}
func (UnimplementedMarketServer) GetSearchData(context.Context, *emptypb.Empty) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchData not implemented")
}
func (UnimplementedMarketServer) GetPayConfig(context.Context, *emptypb.Empty) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayConfig not implemented")
}
func (UnimplementedMarketServer) GetCouponList(context.Context, *common.SkuIDReq) (*CouponListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponList not implemented")
}
func (UnimplementedMarketServer) GetMyCouponList(context.Context, *emptypb.Empty) (*CouponListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCouponList not implemented")
}
func (UnimplementedMarketServer) CouponDraw(context.Context, *CouponReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CouponDraw not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_GetHomeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetHomeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetHomeData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetHomeData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetHomeCatData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetHomeCatData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetHomeCatData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetHomeCatData(ctx, req.(*CatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetNoticeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetNoticeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetNoticeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetNoticeList(ctx, req.(*common.PageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetSearchData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetSearchData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetSearchData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetSearchData(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetPayConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetPayConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetPayConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetPayConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetCouponList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.SkuIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetCouponList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetCouponList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetCouponList(ctx, req.(*common.SkuIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetMyCouponList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetMyCouponList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_GetMyCouponList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetMyCouponList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CouponDraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CouponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CouponDraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_CouponDraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CouponDraw(ctx, req.(*CouponReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHomeData",
			Handler:    _Market_GetHomeData_Handler,
		},
		{
			MethodName: "GetHomeCatData",
			Handler:    _Market_GetHomeCatData_Handler,
		},
		{
			MethodName: "GetNoticeList",
			Handler:    _Market_GetNoticeList_Handler,
		},
		{
			MethodName: "GetSearchData",
			Handler:    _Market_GetSearchData_Handler,
		},
		{
			MethodName: "GetPayConfig",
			Handler:    _Market_GetPayConfig_Handler,
		},
		{
			MethodName: "GetCouponList",
			Handler:    _Market_GetCouponList_Handler,
		},
		{
			MethodName: "GetMyCouponList",
			Handler:    _Market_GetMyCouponList_Handler,
		},
		{
			MethodName: "CouponDraw",
			Handler:    _Market_CouponDraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/market.proto",
}
