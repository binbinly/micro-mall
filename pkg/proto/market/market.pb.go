//*
// 营销服务

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: market/market.proto

package market

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "pkg/proto/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 分类id
type CatIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CatIDReq) Reset() {
	*x = CatIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatIDReq) ProtoMessage() {}

func (x *CatIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatIDReq.ProtoReflect.Descriptor instead.
func (*CatIDReq) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{0}
}

func (x *CatIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 优惠券请求结构
type CouponIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` /// 优惠券id
}

func (x *CouponIDReq) Reset() {
	*x = CouponIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponIDReq) ProtoMessage() {}

func (x *CouponIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponIDReq.ProtoReflect.Descriptor instead.
func (*CouponIDReq) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{1}
}

func (x *CouponIDReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 优惠券使用请求结构
type CouponUsedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponId int64 `protobuf:"varint,2,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"` /// 优惠券id
	OrderId  int64 `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`    /// 订单id
}

func (x *CouponUsedReq) Reset() {
	*x = CouponUsedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponUsedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponUsedReq) ProtoMessage() {}

func (x *CouponUsedReq) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponUsedReq.ProtoReflect.Descriptor instead.
func (*CouponUsedReq) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{2}
}

func (x *CouponUsedReq) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *CouponUsedReq) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

// 首页分类数据
type HomeCatDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.AppSetting `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` /// 分类下配置页面数据
}

func (x *HomeCatDataReply) Reset() {
	*x = HomeCatDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeCatDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeCatDataReply) ProtoMessage() {}

func (x *HomeCatDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeCatDataReply.ProtoReflect.Descriptor instead.
func (*HomeCatDataReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{3}
}

func (x *HomeCatDataReply) GetData() []*common.AppSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

// 首页数据
type HomeDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.HomeData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HomeDataReply) Reset() {
	*x = HomeDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeDataReply) ProtoMessage() {}

func (x *HomeDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeDataReply.ProtoReflect.Descriptor instead.
func (*HomeDataReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{4}
}

func (x *HomeDataReply) GetData() []*common.HomeData {
	if x != nil {
		return x.Data
	}
	return nil
}

// 页面配置列表
type AppSettingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.AppSetting `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AppSettingReply) Reset() {
	*x = AppSettingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSettingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSettingReply) ProtoMessage() {}

func (x *AppSettingReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppSettingReply.ProtoReflect.Descriptor instead.
func (*AppSettingReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{5}
}

func (x *AppSettingReply) GetData() []*common.AppSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

// 公告列表
type NoticeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Notice `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` /// 公告列表
}

func (x *NoticeReply) Reset() {
	*x = NoticeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeReply) ProtoMessage() {}

func (x *NoticeReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeReply.ProtoReflect.Descriptor instead.
func (*NoticeReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{6}
}

func (x *NoticeReply) GetData() []*common.Notice {
	if x != nil {
		return x.Data
	}
	return nil
}

// 搜索页配置
type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.AppSetting `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` /// 搜索页配置数据
	Hot  []string             `protobuf:"bytes,2,rep,name=hot,proto3" json:"hot,omitempty"`   /// 搜索热词
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReply.ProtoReflect.Descriptor instead.
func (*SearchReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{7}
}

func (x *SearchReply) GetData() []*common.AppSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SearchReply) GetHot() []string {
	if x != nil {
		return x.Hot
	}
	return nil
}

// 支付配置
type PayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Pay `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PayReply) Reset() {
	*x = PayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayReply) ProtoMessage() {}

func (x *PayReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayReply.ProtoReflect.Descriptor instead.
func (*PayReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{8}
}

func (x *PayReply) GetData() []*common.Pay {
	if x != nil {
		return x.Data
	}
	return nil
}

// 优惠券列表
type CouponListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Coupon `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CouponListReply) Reset() {
	*x = CouponListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponListReply) ProtoMessage() {}

func (x *CouponListReply) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponListReply.ProtoReflect.Descriptor instead.
func (*CouponListReply) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{9}
}

func (x *CouponListReply) GetData() []*common.Coupon {
	if x != nil {
		return x.Data
	}
	return nil
}

// ---- 内部响应 ----
// 优惠券详情
type CouponInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupon *common.Coupon `protobuf:"bytes,1,opt,name=coupon,proto3" json:"coupon,omitempty"`
}

func (x *CouponInternal) Reset() {
	*x = CouponInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_market_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponInternal) ProtoMessage() {}

func (x *CouponInternal) ProtoReflect() protoreflect.Message {
	mi := &file_market_market_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponInternal.ProtoReflect.Descriptor instead.
func (*CouponInternal) Descriptor() ([]byte, []int) {
	return file_market_market_proto_rawDescGZIP(), []int{10}
}

func (x *CouponInternal) GetCoupon() *common.Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

var File_market_market_proto protoreflect.FileDescriptor

var file_market_market_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x0d,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x61, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x35, 0x0a, 0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x68, 0x6f, 0x74, 0x22,
	0x2b, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0f,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x32, 0xe8, 0x04,
	0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x43, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x43, 0x61, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x72, 0x61, 0x77,
	0x12, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a,
	0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_market_market_proto_rawDescOnce sync.Once
	file_market_market_proto_rawDescData = file_market_market_proto_rawDesc
)

func file_market_market_proto_rawDescGZIP() []byte {
	file_market_market_proto_rawDescOnce.Do(func() {
		file_market_market_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_market_proto_rawDescData)
	})
	return file_market_market_proto_rawDescData
}

var file_market_market_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_market_market_proto_goTypes = []interface{}{
	(*CatIDReq)(nil),          // 0: market.CatIDReq
	(*CouponIDReq)(nil),       // 1: market.CouponIDReq
	(*CouponUsedReq)(nil),     // 2: market.CouponUsedReq
	(*HomeCatDataReply)(nil),  // 3: market.HomeCatDataReply
	(*HomeDataReply)(nil),     // 4: market.HomeDataReply
	(*AppSettingReply)(nil),   // 5: market.AppSettingReply
	(*NoticeReply)(nil),       // 6: market.NoticeReply
	(*SearchReply)(nil),       // 7: market.SearchReply
	(*PayReply)(nil),          // 8: market.PayReply
	(*CouponListReply)(nil),   // 9: market.CouponListReply
	(*CouponInternal)(nil),    // 10: market.CouponInternal
	(*common.AppSetting)(nil), // 11: common.AppSetting
	(*common.HomeData)(nil),   // 12: common.HomeData
	(*common.Notice)(nil),     // 13: common.Notice
	(*common.Pay)(nil),        // 14: common.Pay
	(*common.Coupon)(nil),     // 15: common.Coupon
	(*empty.Empty)(nil),       // 16: google.protobuf.Empty
	(*common.PageReq)(nil),    // 17: common.PageReq
	(*common.SkuIDReq)(nil),   // 18: common.SkuIDReq
}
var file_market_market_proto_depIdxs = []int32{
	11, // 0: market.HomeCatDataReply.data:type_name -> common.AppSetting
	12, // 1: market.HomeDataReply.data:type_name -> common.HomeData
	11, // 2: market.AppSettingReply.data:type_name -> common.AppSetting
	13, // 3: market.NoticeReply.data:type_name -> common.Notice
	11, // 4: market.SearchReply.data:type_name -> common.AppSetting
	14, // 5: market.PayReply.data:type_name -> common.Pay
	15, // 6: market.CouponListReply.data:type_name -> common.Coupon
	15, // 7: market.CouponInternal.coupon:type_name -> common.Coupon
	16, // 8: market.Market.GetHomeData:input_type -> google.protobuf.Empty
	0,  // 9: market.Market.GetHomeCatData:input_type -> market.CatIDReq
	17, // 10: market.Market.GetNoticeList:input_type -> common.PageReq
	16, // 11: market.Market.GetSearchData:input_type -> google.protobuf.Empty
	16, // 12: market.Market.GetPayConfig:input_type -> google.protobuf.Empty
	18, // 13: market.Market.GetCouponList:input_type -> common.SkuIDReq
	16, // 14: market.Market.GetMyCouponList:input_type -> google.protobuf.Empty
	1,  // 15: market.Market.CouponDraw:input_type -> market.CouponIDReq
	2,  // 16: market.Market.CouponUsed:input_type -> market.CouponUsedReq
	1,  // 17: market.Market.GetCouponInfo:input_type -> market.CouponIDReq
	4,  // 18: market.Market.GetHomeData:output_type -> market.HomeDataReply
	5,  // 19: market.Market.GetHomeCatData:output_type -> market.AppSettingReply
	6,  // 20: market.Market.GetNoticeList:output_type -> market.NoticeReply
	7,  // 21: market.Market.GetSearchData:output_type -> market.SearchReply
	8,  // 22: market.Market.GetPayConfig:output_type -> market.PayReply
	9,  // 23: market.Market.GetCouponList:output_type -> market.CouponListReply
	9,  // 24: market.Market.GetMyCouponList:output_type -> market.CouponListReply
	16, // 25: market.Market.CouponDraw:output_type -> google.protobuf.Empty
	16, // 26: market.Market.CouponUsed:output_type -> google.protobuf.Empty
	10, // 27: market.Market.GetCouponInfo:output_type -> market.CouponInternal
	18, // [18:28] is the sub-list for method output_type
	8,  // [8:18] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_market_market_proto_init() }
func file_market_market_proto_init() {
	if File_market_market_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_market_market_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatIDReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponIDReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponUsedReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeCatDataReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeDataReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppSettingReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponListReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_market_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponInternal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_market_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_market_proto_goTypes,
		DependencyIndexes: file_market_market_proto_depIdxs,
		MessageInfos:      file_market_market_proto_msgTypes,
	}.Build()
	File_market_market_proto = out.File
	file_market_market_proto_rawDesc = nil
	file_market_market_proto_goTypes = nil
	file_market_market_proto_depIdxs = nil
}
