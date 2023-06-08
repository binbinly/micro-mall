package handler

import (
	"context"
	"market/logic"
	"market/resource"
	"pkg/errno"
	"pkg/handler"

	"github.com/golang/protobuf/ptypes/empty"
	cpb "pkg/proto/common"
	pb "pkg/proto/market"
)

// Auth 营销服身份验证
func Auth(method string) bool {
	switch method {
	case "Market.GetCouponList", "Market.GetMyCouponList", "Market.CouponDraw",
		"Market.GetCouponInfo", "Market.CouponUsed":
		//这些路由需要身份验证
		return true
	}
	return false
}

// Market 营销服务
type Market struct {
	logic logic.Logic
}

func New(l logic.Logic) *Market {
	return &Market{logic: l}
}

// GetHomeData 首页数据
func (m *Market) GetHomeData(ctx context.Context, empty *empty.Empty, reply *pb.HomeDataReply) error {
	cats, err := m.logic.GetHomeCat(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	data, err := m.logic.GetHomeCatData(ctx, 0)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	catData := make([]*cpb.HomeData, 0, len(cats)+1)
	catData = append(catData, &cpb.HomeData{
		Id:   0,
		Name: "全部",
		List: resource.AppSettingResource(data),
	})

	for _, cat := range cats {
		catData = append(catData, &cpb.HomeData{
			Id:   int32(cat.ID),
			Name: cat.Name,
			List: []*cpb.AppSetting{},
		})
	}

	reply.Data = catData
	return nil
}

// GetHomeCatData 获取首页分类下数据
func (m *Market) GetHomeCatData(ctx context.Context, req *pb.CatIDReq, reply *pb.AppSettingReply) error {
	list, err := m.logic.GetHomeCatData(ctx, int(req.Id))
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	reply.Data = resource.AppSettingResource(list)
	return nil
}

// GetNoticeList 公告列表
func (m *Market) GetNoticeList(ctx context.Context, req *cpb.PageReq, reply *pb.NoticeReply) error {
	list, err := m.logic.GetNoticeList(ctx, req.Page)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = resource.AppNoticeResource(list)
	return nil
}

// GetSearchData 获取搜索页数据
func (m *Market) GetSearchData(ctx context.Context, empty *empty.Empty, reply *pb.SearchReply) error {
	data, hot, err := m.logic.GetSearchData(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	reply.Data = resource.AppSettingResource(data)
	reply.Hot = hot
	return nil
}

// GetPayConfig 获取支付配置
func (m *Market) GetPayConfig(ctx context.Context, empty *empty.Empty, reply *pb.PayReply) error {
	list, err := m.logic.GetPayConfig(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	reply.Data = resource.PayResource(list)
	return nil
}

// GetCouponList 优惠券列表
func (m *Market) GetCouponList(ctx context.Context, req *cpb.SkuIDReq, reply *pb.CouponListReply) error {
	list, ids, err := m.logic.GetCouponList(ctx, handler.GetUserID(ctx), req.Id)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	reply.Data = resource.CouponListResource(list, ids)
	return nil
}

// GetMyCouponList 我的优惠券
func (m *Market) GetMyCouponList(ctx context.Context, empty *empty.Empty, reply *pb.CouponListReply) error {
	list, err := m.logic.GetMyCouponList(ctx, handler.GetUserID(ctx))
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	reply.Data = resource.MyCouponListResource(list)
	return nil
}

// CouponDraw 领取优惠券
func (m *Market) CouponDraw(ctx context.Context, req *pb.CouponIDReq, empty *empty.Empty) error {
	if err := m.logic.CouponDraw(ctx, handler.GetUserID(ctx), req.Id); err != nil {
		return errno.MarketReplyErr(err)
	}

	return nil
}

// CouponUsed 使用优惠券
func (m *Market) CouponUsed(ctx context.Context, req *pb.CouponUsedReq, empty *empty.Empty) error {
	if err := m.logic.CouponUsed(ctx, handler.GetUserID(ctx), req.CouponId, req.OrderId); err != nil {
		return errno.MarketReplyErr(err)
	}
	return nil
}

// GetCouponInfo 获取优惠券信息
func (m *Market) GetCouponInfo(ctx context.Context, req *pb.CouponIDReq, internal *pb.CouponInternal) error {
	mCoupon, coupon, err := m.logic.GetCouponInfo(ctx, handler.GetUserID(ctx), req.Id)
	if err != nil {
		return errno.MarketReplyErr(err)
	}

	c := resource.CouponResource(coupon, false)
	c.Status = int32(mCoupon.Status)
	internal.Coupon = c
	return nil
}
