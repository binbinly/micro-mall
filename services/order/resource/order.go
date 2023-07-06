package resource

import (
	"strings"

	"github.com/binbinly/pkg/util"

	"order/config"
	"order/model"
	cpb "pkg/proto/common"
)

// OrderResource 转换订单详情输出
func OrderResource(order *model.OrderModel) *cpb.Order {
	res := &cpb.Order{
		Id:              int32(order.ID),
		OrderNo:         order.OrderNo,
		Note:            order.Note,
		TotalAmount:     util.ParseAmount(order.TotalAmount),
		Amount:          util.ParseAmount(order.Amount),
		CouponAmount:    util.ParseAmount(order.CouponAmount),
		FreightAmount:   util.ParseAmount(order.FreightAmount),
		PayAmount:       util.ParseAmount(order.PayAmount),
		PayType:         int32(order.PayType),
		PayAt:           int32(order.PayAt),
		CreateAt:        int32(order.CreatedAt),
		Status:          int32(order.Status),
		TradeNo:         order.TradeNo,
		DeliveryCompany: order.DeliveryCompany,
		DeliveryNo:      order.DeliveryNo,
		Integration:     int32(order.Integration),
		Growth:          int32(order.Growth),
		DeliveryAt:      int32(order.DeliveryAt),
		ReceiveAt:       int32(order.ReceiveAt),
		CommentAt:       int32(order.CommentAt),
		Address: &cpb.OrderAddress{
			Name:   order.AddressName,
			Phone:  order.AddressPhone,
			Area:   strings.Join([]string{order.AddressProvince, order.AddressCity, order.AddressCounty}, " "),
			Detail: order.AddressDetail,
		},
		Items: make([]*cpb.OrderSku, 0, len(order.Items)),
	}
	for _, item := range order.Items {
		res.Items = append(res.Items, &cpb.OrderSku{
			SkuId:     int32(item.SkuID),
			Title:     item.SkuName,
			Cover:     util.FormatResUrl(config.Cfg.DFS.Endpoint, item.SkuImg),
			Price:     util.ParseAmount(item.SkuPrice),
			Num:       int32(item.Num),
			AttrValue: item.SkuAttrs,
		})
	}
	return res
}

// OrderListResource 转换订单列表输出
func OrderListResource(list []*model.OrderModel) []*cpb.OrderList {
	if len(list) == 0 {
		return []*cpb.OrderList{}
	}
	res := make([]*cpb.OrderList, 0, len(list))
	for _, order := range list {
		ol := &cpb.OrderList{
			Id:      int32(order.ID),
			OrderNo: order.OrderNo,
			Amount:  util.ParseAmount(order.Amount),
			Status:  int32(order.Status),
			Time:    int32(order.CreatedAt),
			Items:   make([]*cpb.OrderSku, 0, len(order.Items)),
		}

		for _, item := range order.Items {
			ol.Items = append(ol.Items, &cpb.OrderSku{
				SkuId:     int32(item.SkuID),
				Title:     item.SkuName,
				Cover:     util.FormatResUrl(config.Cfg.DFS.Endpoint, item.SkuImg),
				Price:     util.ParseAmount(item.SkuPrice),
				Num:       int32(item.Num),
				AttrValue: item.SkuAttrs,
			})
		}
		res = append(res, ol)
	}
	return res
}
