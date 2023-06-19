//*
// 购物车服务

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: cart/cart.proto

package cart

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

// 添加购物车请求结构
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId int64 `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` /// sku_id
	Num   int32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`                  /// 数量
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *AddReq) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

// 修改购物车购微项请求结构
type EditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldSkuId int64 `protobuf:"varint,2,opt,name=old_sku_id,json=oldSkuId,proto3" json:"old_sku_id,omitempty"` /// 修改前商品id
	NewSkuId int64 `protobuf:"varint,3,opt,name=new_sku_id,json=newSkuId,proto3" json:"new_sku_id,omitempty"` /// 修改后商品id
	Num      int32 `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`                             /// 数量
}

func (x *EditReq) Reset() {
	*x = EditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditReq) ProtoMessage() {}

func (x *EditReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditReq.ProtoReflect.Descriptor instead.
func (*EditReq) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *EditReq) GetOldSkuId() int64 {
	if x != nil {
		return x.OldSkuId
	}
	return 0
}

func (x *EditReq) GetNewSkuId() int64 {
	if x != nil {
		return x.NewSkuId
	}
	return 0
}

func (x *EditReq) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

// 多sku_id请求结构
type SkusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuIds []int64 `protobuf:"varint,2,rep,packed,name=sku_ids,json=skuIds,proto3" json:"sku_ids,omitempty"` /// sku_id数组
}

func (x *SkusReq) Reset() {
	*x = SkusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkusReq) ProtoMessage() {}

func (x *SkusReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkusReq.ProtoReflect.Descriptor instead.
func (*SkusReq) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *SkusReq) GetSkuIds() []int64 {
	if x != nil {
		return x.SkuIds
	}
	return nil
}

// 购物车列表
type CartsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*common.Cart `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CartsReply) Reset() {
	*x = CartsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartsReply) ProtoMessage() {}

func (x *CartsReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartsReply.ProtoReflect.Descriptor instead.
func (*CartsReply) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartsReply) GetData() []*common.Cart {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_cart_cart_proto protoreflect.FileDescriptor

var file_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x06,
	0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02,
	0x20, 0x00, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x72, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x25, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x08, 0x6f, 0x6c, 0x64, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x6e, 0x65, 0x77,
	0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x53, 0x6b, 0x75, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2c, 0x0a, 0x07, 0x53,
	0x6b, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x06, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x0a, 0x43, 0x61, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xaf, 0x03, 0x0a, 0x04, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0c, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x0c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x44,
	0x65, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x6b, 0x75, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3b, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a,
	0x06, 0x4d, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x30, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x43,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x53, 0x6b, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cart_cart_proto_rawDescOnce sync.Once
	file_cart_cart_proto_rawDescData = file_cart_cart_proto_rawDesc
)

func file_cart_cart_proto_rawDescGZIP() []byte {
	file_cart_cart_proto_rawDescOnce.Do(func() {
		file_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_cart_proto_rawDescData)
	})
	return file_cart_cart_proto_rawDescData
}

var file_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cart_cart_proto_goTypes = []interface{}{
	(*AddReq)(nil),          // 0: cart.AddReq
	(*EditReq)(nil),         // 1: cart.EditReq
	(*SkusReq)(nil),         // 2: cart.SkusReq
	(*CartsReply)(nil),      // 3: cart.CartsReply
	(*common.Cart)(nil),     // 4: common.Cart
	(*common.SkuIDReq)(nil), // 5: common.SkuIDReq
	(*empty.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_cart_cart_proto_depIdxs = []int32{
	4, // 0: cart.CartsReply.data:type_name -> common.Cart
	0, // 1: cart.Cart.AddCart:input_type -> cart.AddReq
	1, // 2: cart.Cart.EditCart:input_type -> cart.EditReq
	0, // 3: cart.Cart.EditCartNum:input_type -> cart.AddReq
	5, // 4: cart.Cart.DelCart:input_type -> common.SkuIDReq
	6, // 5: cart.Cart.ClearCart:input_type -> google.protobuf.Empty
	6, // 6: cart.Cart.MyCart:input_type -> google.protobuf.Empty
	2, // 7: cart.Cart.BatchGetCarts:input_type -> cart.SkusReq
	2, // 8: cart.Cart.BatchDelCarts:input_type -> cart.SkusReq
	6, // 9: cart.Cart.AddCart:output_type -> google.protobuf.Empty
	6, // 10: cart.Cart.EditCart:output_type -> google.protobuf.Empty
	6, // 11: cart.Cart.EditCartNum:output_type -> google.protobuf.Empty
	6, // 12: cart.Cart.DelCart:output_type -> google.protobuf.Empty
	6, // 13: cart.Cart.ClearCart:output_type -> google.protobuf.Empty
	3, // 14: cart.Cart.MyCart:output_type -> cart.CartsReply
	3, // 15: cart.Cart.BatchGetCarts:output_type -> cart.CartsReply
	6, // 16: cart.Cart.BatchDelCarts:output_type -> google.protobuf.Empty
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_cart_proto_init() }
func file_cart_cart_proto_init() {
	if File_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditReq); i {
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
		file_cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkusReq); i {
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
		file_cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartsReply); i {
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
			RawDescriptor: file_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_cart_proto_goTypes,
		DependencyIndexes: file_cart_cart_proto_depIdxs,
		MessageInfos:      file_cart_cart_proto_msgTypes,
	}.Build()
	File_cart_cart_proto = out.File
	file_cart_cart_proto_rawDesc = nil
	file_cart_cart_proto_goTypes = nil
	file_cart_cart_proto_depIdxs = nil
}