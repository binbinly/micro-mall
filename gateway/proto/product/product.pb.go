//*
// 产品服务

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: product/product.proto

package gateway

import (
	common "gateway/proto/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 商品列表请求结构
type SkuListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatId int64 `protobuf:"varint,1,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"` // 分类
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                // 分页
}

func (x *SkuListReq) Reset() {
	*x = SkuListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuListReq) ProtoMessage() {}

func (x *SkuListReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuListReq.ProtoReflect.Descriptor instead.
func (*SkuListReq) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *SkuListReq) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *SkuListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

// 搜索请求结构
type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword  string         `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`                        /// 关键字
	CatId    int64          `protobuf:"varint,2,opt,name=cat_id,json=catId,proto3" json:"cat_id,omitempty"`              /// 分类id
	Field    int32          `protobuf:"varint,3,opt,name=field,proto3" json:"field,omitempty"`                           /// 排序字段
	Order    int32          `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`                           /// 排序类型 0=asc 1=desc
	HasStock bool           `protobuf:"varint,5,opt,name=has_stock,json=hasStock,proto3" json:"has_stock,omitempty"`     // 是否有库存
	PriceS   int32          `protobuf:"varint,6,opt,name=price_s,json=priceS,proto3" json:"price_s,omitempty"`           /// 价格区间起始
	PriceE   int32          `protobuf:"varint,7,opt,name=price_e,json=priceE,proto3" json:"price_e,omitempty"`           /// 价格区间止
	BrandId  []int64        `protobuf:"varint,8,rep,packed,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"` /// 品牌,多选
	Attrs    []*SearchAttrs `protobuf:"bytes,9,rep,name=attrs,proto3" json:"attrs,omitempty"`                            /// 规格属性
	Page     int32          `protobuf:"varint,10,opt,name=page,proto3" json:"page,omitempty"`                            /// 分页
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *SearchReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchReq) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *SearchReq) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

func (x *SearchReq) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *SearchReq) GetHasStock() bool {
	if x != nil {
		return x.HasStock
	}
	return false
}

func (x *SearchReq) GetPriceS() int32 {
	if x != nil {
		return x.PriceS
	}
	return 0
}

func (x *SearchReq) GetPriceE() int32 {
	if x != nil {
		return x.PriceE
	}
	return 0
}

func (x *SearchReq) GetBrandId() []int64 {
	if x != nil {
		return x.BrandId
	}
	return nil
}

func (x *SearchReq) GetAttrs() []*SearchAttrs {
	if x != nil {
		return x.Attrs
	}
	return nil
}

func (x *SearchReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

// 搜索规格结构
type SearchAttrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`        // 规格名id
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"` // 规格值列表
}

func (x *SearchAttrs) Reset() {
	*x = SearchAttrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchAttrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAttrs) ProtoMessage() {}

func (x *SearchAttrs) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAttrs.ProtoReflect.Descriptor instead.
func (*SearchAttrs) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *SearchAttrs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchAttrs) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// 产品分类
type CategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32              `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string             `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*common.Category `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CategoryReply) Reset() {
	*x = CategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryReply) ProtoMessage() {}

func (x *CategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryReply.ProtoReflect.Descriptor instead.
func (*CategoryReply) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CategoryReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CategoryReply) GetData() []*common.Category {
	if x != nil {
		return x.Data
	}
	return nil
}

// 商品详情
type SkuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *common.Sku `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SkuReply) Reset() {
	*x = SkuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuReply) ProtoMessage() {}

func (x *SkuReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuReply.ProtoReflect.Descriptor instead.
func (*SkuReply) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *SkuReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SkuReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SkuReply) GetData() *common.Sku {
	if x != nil {
		return x.Data
	}
	return nil
}

// 产品列表响应结构
type SkuListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string          `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*common.SkuEs `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` /// 产品列表
}

func (x *SkuListReply) Reset() {
	*x = SkuListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuListReply) ProtoMessage() {}

func (x *SkuListReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuListReply.ProtoReflect.Descriptor instead.
func (*SkuListReply) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *SkuListReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SkuListReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SkuListReply) GetData() []*common.SkuEs {
	if x != nil {
		return x.Data
	}
	return nil
}

// 搜索结构
type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`      /// 状态码
	Message string            `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"` /// 消息
	Data    []*common.SkuEs   `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`       /// 搜索商品
	Brands  []*common.BrandEs `protobuf:"bytes,2,rep,name=brands,proto3" json:"brands,omitempty"`   /// 当前查询到的结果锁涉及到的品牌
	Attrs   []*common.AttrEs  `protobuf:"bytes,3,rep,name=attrs,proto3" json:"attrs,omitempty"`     /// 当前查询到的结果锁涉及到的所有属性
	Cats    []*common.CatEs   `protobuf:"bytes,4,rep,name=cats,proto3" json:"cats,omitempty"`       /// 当前查询到的结果锁涉及到的所有分类
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[6]
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
	return file_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *SearchReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SearchReply) GetData() []*common.SkuEs {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SearchReply) GetBrands() []*common.BrandEs {
	if x != nil {
		return x.Brands
	}
	return nil
}

func (x *SearchReply) GetAttrs() []*common.AttrEs {
	if x != nil {
		return x.Attrs
	}
	return nil
}

func (x *SearchReply) GetCats() []*common.CatEs {
	if x != nil {
		return x.Cats
	}
	return nil
}

// sku销售属性
type SkuSaleAttrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32               `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data    *common.SkuSaleAttr `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SkuSaleAttrReply) Reset() {
	*x = SkuSaleAttrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuSaleAttrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuSaleAttrReply) ProtoMessage() {}

func (x *SkuSaleAttrReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuSaleAttrReply.ProtoReflect.Descriptor instead.
func (*SkuSaleAttrReply) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *SkuSaleAttrReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SkuSaleAttrReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SkuSaleAttrReply) GetData() *common.SkuSaleAttr {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x37, 0x0a, 0x0a, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x15,
	0x0a, 0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x92, 0x02, 0x0a, 0x09, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x45, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74,
	0x74, 0x72, 0x73, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x35,
	0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x08, 0x53, 0x6b,
	0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x0c, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75, 0x45, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd0, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75, 0x45,
	0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x45, 0x73, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x24, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x45, 0x73, 0x52,
	0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x63, 0x61, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61,
	0x74, 0x45, 0x73, 0x52, 0x04, 0x63, 0x61, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x10, 0x53, 0x6b, 0x75,
	0x53, 0x61, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xe5, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x57, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x65, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x63, 0x61, 0x74, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x6b, 0x75,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x79, 0x0a, 0x07, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x5a, 0x28,
	0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x63, 0x61, 0x74, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x70, 0x2f, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x7d, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x09, 0x53, 0x6b,
	0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x6b, 0x75, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x73,
	0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x6b, 0x75,
	0x53, 0x61, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData = file_product_product_proto_rawDesc
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_product_proto_rawDescData)
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_product_product_proto_goTypes = []interface{}{
	(*SkuListReq)(nil),         // 0: product.SkuListReq
	(*SearchReq)(nil),          // 1: product.SearchReq
	(*SearchAttrs)(nil),        // 2: product.SearchAttrs
	(*CategoryReply)(nil),      // 3: product.CategoryReply
	(*SkuReply)(nil),           // 4: product.SkuReply
	(*SkuListReply)(nil),       // 5: product.SkuListReply
	(*SearchReply)(nil),        // 6: product.SearchReply
	(*SkuSaleAttrReply)(nil),   // 7: product.SkuSaleAttrReply
	(*common.Category)(nil),    // 8: common.Category
	(*common.Sku)(nil),         // 9: common.Sku
	(*common.SkuEs)(nil),       // 10: common.SkuEs
	(*common.BrandEs)(nil),     // 11: common.BrandEs
	(*common.AttrEs)(nil),      // 12: common.AttrEs
	(*common.CatEs)(nil),       // 13: common.CatEs
	(*common.SkuSaleAttr)(nil), // 14: common.SkuSaleAttr
	(*emptypb.Empty)(nil),      // 15: google.protobuf.Empty
	(*common.SkuIDReq)(nil),    // 16: common.SkuIDReq
}
var file_product_product_proto_depIdxs = []int32{
	2,  // 0: product.SearchReq.attrs:type_name -> product.SearchAttrs
	8,  // 1: product.CategoryReply.data:type_name -> common.Category
	9,  // 2: product.SkuReply.data:type_name -> common.Sku
	10, // 3: product.SkuListReply.data:type_name -> common.SkuEs
	10, // 4: product.SearchReply.data:type_name -> common.SkuEs
	11, // 5: product.SearchReply.brands:type_name -> common.BrandEs
	12, // 6: product.SearchReply.attrs:type_name -> common.AttrEs
	13, // 7: product.SearchReply.cats:type_name -> common.CatEs
	14, // 8: product.SkuSaleAttrReply.data:type_name -> common.SkuSaleAttr
	15, // 9: product.Product.CategoryTree:input_type -> google.protobuf.Empty
	1,  // 10: product.Product.SkuSearch:input_type -> product.SearchReq
	0,  // 11: product.Product.SkuList:input_type -> product.SkuListReq
	16, // 12: product.Product.SkuDetail:input_type -> common.SkuIDReq
	16, // 13: product.Product.GetSkuSaleAttrs:input_type -> common.SkuIDReq
	3,  // 14: product.Product.CategoryTree:output_type -> product.CategoryReply
	6,  // 15: product.Product.SkuSearch:output_type -> product.SearchReply
	5,  // 16: product.Product.SkuList:output_type -> product.SkuListReply
	4,  // 17: product.Product.SkuDetail:output_type -> product.SkuReply
	7,  // 18: product.Product.GetSkuSaleAttrs:output_type -> product.SkuSaleAttrReply
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuListReq); i {
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
		file_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_product_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchAttrs); i {
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
		file_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryReply); i {
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
		file_product_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuReply); i {
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
		file_product_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuListReply); i {
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
		file_product_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_product_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuSaleAttrReply); i {
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
			RawDescriptor: file_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_rawDesc = nil
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}
