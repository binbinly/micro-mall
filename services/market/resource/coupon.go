package resource

import (
	"github.com/binbinly/pkg/util"
	"golang.org/x/exp/slices"
	"market/model"
	cpb "pkg/proto/common"
)

// CouponListResource 转换优惠券列表输出
func CouponListResource(list []*model.CouponModel, ids []int64) []*cpb.Coupon {
	if len(list) == 0 {
		return []*cpb.Coupon{}
	}
	res := make([]*cpb.Coupon, 0, len(list))
	for _, c := range list {
		status := model.CouponStatusNotReceive
		if slices.Contains(ids, c.ID) {
			status = model.CouponStatusInit
		}
		res = append(res, &cpb.Coupon{
			Id:       int32(c.ID),
			Name:     c.Name,
			Amount:   util.ParseAmount(c.Amount),
			MinPoint: util.ParseAmount(c.MinPoint),
			StartAt:  int32(c.StartAt),
			EndAt:    int32(c.EndAt),
			Note:     c.Note,
			Status:   int32(status),
		})
	}
	return res
}

// MyCouponListResource 转换优惠券列表输出
func MyCouponListResource(list []*model.Coupon) []*cpb.Coupon {
	if len(list) == 0 {
		return []*cpb.Coupon{}
	}
	res := make([]*cpb.Coupon, 0, len(list))
	for _, c := range list {
		res = append(res, &cpb.Coupon{
			Id:       int32(c.ID),
			Name:     c.Name,
			Amount:   util.ParseAmount(c.Amount),
			MinPoint: util.ParseAmount(c.MinPoint),
			StartAt:  int32(c.StartAt),
			EndAt:    int32(c.EndAt),
			Note:     c.Note,
			Status:   int32(c.Status),
		})
	}
	return res
}

// CouponResource 转换优惠券详情
func CouponResource(c *model.CouponModel, external bool) *cpb.Coupon {
	coupon := &cpb.Coupon{
		Id:      int32(c.ID),
		Name:    c.Name,
		StartAt: int32(c.StartAt),
		EndAt:   int32(c.EndAt),
		Note:    c.Note,
	}
	if external {
		coupon.Amount = util.ParseAmount(c.Amount)
		coupon.MinPoint = util.ParseAmount(c.MinPoint)
	} else {
		coupon.Amount = float64(c.Amount)
		coupon.MinPoint = float64(c.MinPoint)
	}

	return coupon
}
