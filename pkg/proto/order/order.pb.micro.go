// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order/order.proto

package order

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	proto "google.golang.org/protobuf/proto"
	math "math"
	common "pkg/proto/common"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	// / 从购物车提交订单
	SubmitOrder(ctx context.Context, in *SubmitReq, opts ...client.CallOption) (*OrderIDReply, error)
	// / 商品直接提交订单
	SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, opts ...client.CallOption) (*OrderIDReply, error)
	// / 订单详情
	OrderDetail(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*OrderInfoReply, error)
	// / 订单取消
	OrderCancel(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*empty.Empty, error)
	// / 订单列表
	OrderList(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListReply, error)
	// / 订单支付成功回调
	OrderPayNotify(ctx context.Context, in *PayNotifyReq, opts ...client.CallOption) (*empty.Empty, error)
	// / 订单退款
	OrderRefund(ctx context.Context, in *RefundReq, opts ...client.CallOption) (*empty.Empty, error)
	// / 订单确认收货
	OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*empty.Empty, error)
	// / 订单评价
	OrderComment(ctx context.Context, in *CommentReq, opts ...client.CallOption) (*empty.Empty, error)
	// / 内部调用 订单信息
	GetOrderByID(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*common.Order, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) SubmitOrder(ctx context.Context, in *SubmitReq, opts ...client.CallOption) (*OrderIDReply, error) {
	req := c.c.NewRequest(c.name, "Order.SubmitOrder", in)
	out := new(OrderIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, opts ...client.CallOption) (*OrderIDReply, error) {
	req := c.c.NewRequest(c.name, "Order.SubmitSkuOrder", in)
	out := new(OrderIDReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderDetail(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*OrderInfoReply, error) {
	req := c.c.NewRequest(c.name, "Order.OrderDetail", in)
	out := new(OrderInfoReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderCancel(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.OrderCancel", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderList(ctx context.Context, in *ListReq, opts ...client.CallOption) (*ListReply, error) {
	req := c.c.NewRequest(c.name, "Order.OrderList", in)
	out := new(ListReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderPayNotify(ctx context.Context, in *PayNotifyReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.OrderPayNotify", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderRefund(ctx context.Context, in *RefundReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.OrderRefund", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.OrderConfirmReceipt", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) OrderComment(ctx context.Context, in *CommentReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "Order.OrderComment", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrderByID(ctx context.Context, in *OrderIDReq, opts ...client.CallOption) (*common.Order, error) {
	req := c.c.NewRequest(c.name, "Order.GetOrderByID", in)
	out := new(common.Order)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	// / 从购物车提交订单
	SubmitOrder(context.Context, *SubmitReq, *OrderIDReply) error
	// / 商品直接提交订单
	SubmitSkuOrder(context.Context, *SkuSubmitReq, *OrderIDReply) error
	// / 订单详情
	OrderDetail(context.Context, *OrderIDReq, *OrderInfoReply) error
	// / 订单取消
	OrderCancel(context.Context, *OrderIDReq, *empty.Empty) error
	// / 订单列表
	OrderList(context.Context, *ListReq, *ListReply) error
	// / 订单支付成功回调
	OrderPayNotify(context.Context, *PayNotifyReq, *empty.Empty) error
	// / 订单退款
	OrderRefund(context.Context, *RefundReq, *empty.Empty) error
	// / 订单确认收货
	OrderConfirmReceipt(context.Context, *OrderIDReq, *empty.Empty) error
	// / 订单评价
	OrderComment(context.Context, *CommentReq, *empty.Empty) error
	// / 内部调用 订单信息
	GetOrderByID(context.Context, *OrderIDReq, *common.Order) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		SubmitOrder(ctx context.Context, in *SubmitReq, out *OrderIDReply) error
		SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, out *OrderIDReply) error
		OrderDetail(ctx context.Context, in *OrderIDReq, out *OrderInfoReply) error
		OrderCancel(ctx context.Context, in *OrderIDReq, out *empty.Empty) error
		OrderList(ctx context.Context, in *ListReq, out *ListReply) error
		OrderPayNotify(ctx context.Context, in *PayNotifyReq, out *empty.Empty) error
		OrderRefund(ctx context.Context, in *RefundReq, out *empty.Empty) error
		OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, out *empty.Empty) error
		OrderComment(ctx context.Context, in *CommentReq, out *empty.Empty) error
		GetOrderByID(ctx context.Context, in *OrderIDReq, out *common.Order) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) SubmitOrder(ctx context.Context, in *SubmitReq, out *OrderIDReply) error {
	return h.OrderHandler.SubmitOrder(ctx, in, out)
}

func (h *orderHandler) SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, out *OrderIDReply) error {
	return h.OrderHandler.SubmitSkuOrder(ctx, in, out)
}

func (h *orderHandler) OrderDetail(ctx context.Context, in *OrderIDReq, out *OrderInfoReply) error {
	return h.OrderHandler.OrderDetail(ctx, in, out)
}

func (h *orderHandler) OrderCancel(ctx context.Context, in *OrderIDReq, out *empty.Empty) error {
	return h.OrderHandler.OrderCancel(ctx, in, out)
}

func (h *orderHandler) OrderList(ctx context.Context, in *ListReq, out *ListReply) error {
	return h.OrderHandler.OrderList(ctx, in, out)
}

func (h *orderHandler) OrderPayNotify(ctx context.Context, in *PayNotifyReq, out *empty.Empty) error {
	return h.OrderHandler.OrderPayNotify(ctx, in, out)
}

func (h *orderHandler) OrderRefund(ctx context.Context, in *RefundReq, out *empty.Empty) error {
	return h.OrderHandler.OrderRefund(ctx, in, out)
}

func (h *orderHandler) OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, out *empty.Empty) error {
	return h.OrderHandler.OrderConfirmReceipt(ctx, in, out)
}

func (h *orderHandler) OrderComment(ctx context.Context, in *CommentReq, out *empty.Empty) error {
	return h.OrderHandler.OrderComment(ctx, in, out)
}

func (h *orderHandler) GetOrderByID(ctx context.Context, in *OrderIDReq, out *common.Order) error {
	return h.OrderHandler.GetOrderByID(ctx, in, out)
}