package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"order/logic"
	"order/resource"
	"pkg/errno"
	"pkg/handler"
	cpb "pkg/proto/common"
	pb "pkg/proto/order"
	"reflect"
)

// Order 订单服务处理器
type Order struct {
	logic logic.Logic
}

func New(l logic.Logic) *Order {
	return &Order{logic: l}
}

// SubmitOrder 提交订单
func (o *Order) SubmitOrder(ctx context.Context, req *pb.SubmitReq, reply *pb.OrderIDReply) error {
	if err := o.logic.OrderSubmit(ctx, handler.GetUserID(ctx), req.AddressId, req.CouponId, req.SkuIds, req.Note); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// SubmitSkuOrder 不通过购物车，直接在商品页面提交sku订单
func (o *Order) SubmitSkuOrder(ctx context.Context, req *pb.SkuSubmitReq, reply *pb.OrderIDReply) error {
	if err := o.logic.SubmitSkuOrder(ctx, handler.GetUserID(ctx), req.SkuId, req.AddressId, req.CouponId,
		int(req.Num), req.Note); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// OrderDetail 订单详情
func (o *Order) OrderDetail(ctx context.Context, req *pb.OrderIDReq, reply *pb.OrderInfoReply) error {
	info, err := o.logic.OrderDetail(ctx, handler.GetUserID(ctx), req.Id)
	if err != nil {
		return errno.OrderReplyErr(err)
	}

	reply.Data = resource.OrderResource(info)
	return nil
}

// OrderCancel 订单取消
func (o *Order) OrderCancel(ctx context.Context, req *pb.OrderIDReq, empty *empty.Empty) error {
	if err := o.logic.OrderCancel(ctx, handler.GetUserID(ctx), req.Id); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// OrderList 订单列表
func (o *Order) OrderList(ctx context.Context, req *pb.ListReq, reply *pb.ListReply) error {
	list, err := o.logic.MyOrderList(ctx, handler.GetUserID(ctx), int(req.Status), req.Page)
	if err != nil {
		return errno.OrderReplyErr(err)
	}

	reply.Data = resource.OrderListResource(list)
	return nil
}

// OrderPayNotify 订单支付成功通知
func (o *Order) OrderPayNotify(ctx context.Context, req *pb.PayNotifyReq, empty *empty.Empty) error {
	if err := o.logic.OrderPayNotify(ctx, handler.GetUserID(ctx), int(req.PayAmount), req.PayType,
		req.OrderNo, req.TradeNo, req.TransHash); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// OrderRefund 订单退款
func (o *Order) OrderRefund(ctx context.Context, req *pb.RefundReq, empty *empty.Empty) error {
	if err := o.logic.OrderRefund(ctx, handler.GetUserID(ctx), req.OrderId, req.Content); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// OrderConfirmReceipt 确认收货
func (o *Order) OrderConfirmReceipt(ctx context.Context, req *pb.OrderIDReq, empty *empty.Empty) error {
	if err := o.logic.OrderConfirmReceipt(ctx, handler.GetUserID(ctx), req.Id); err != nil {
		return errno.OrderReplyErr(err)
	}

	return nil
}

// OrderComment 订单评价
func (o *Order) OrderComment(ctx context.Context, req *pb.CommentReq, empty *empty.Empty) error {
	if err := o.logic.OrderComment(ctx, handler.GetUserID(ctx), req.OrderId, req.SkuIds, req.Star,
		req.Content, req.Resources); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil

}

// GetOrderByID 订单信息
func (o *Order) GetOrderByID(ctx context.Context, req *pb.OrderIDReq, reply *cpb.Order) error {
	info, err := o.logic.OrderInfo(ctx, req.Id)
	if err != nil {
		return errno.OrderReplyErr(err)
	}

	reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(resource.OrderResource(info)).Elem())
	return nil
}
