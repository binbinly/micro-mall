package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"market/model"
	"pkg/dbs"
)

// GetCouponMemberList 获取用户领取的优惠券列表
func (r *Repo) GetCouponMemberList(ctx context.Context, memberID int64) (list []*model.Coupon, err error) {
	err = r.DB.WithContext(ctx).Table("sms_coupon_member as m").
		Select("`name`,`amount`,min_point,start_at,end_at,`note`,m.id,m.status").
		Scopes(dbs.WhereRelease).Joins("left join sms_coupon as c on m.coupon_id = c.id").
		Where("member_id=?", memberID).Order("amount desc").Scan(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.couponMember] by db uid: %v", memberID)
	}
	return
}

// GetCouponMemberByID 会员优惠券详情
func (r *Repo) GetCouponMemberByID(ctx context.Context, id int64) (coupon *model.CouponMemberModel, err error) {
	coupon = new(model.CouponMemberModel)
	if err = r.DB.WithContext(ctx).First(coupon, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.couponMember] first db")
	}

	return coupon, nil
}

// SetCouponMemberUsed 设置优惠券已使用
func (r *Repo) SetCouponMemberUsed(ctx context.Context, id, memberID, orderID int64) error {
	result := r.DB.WithContext(ctx).Model(&model.CouponMemberModel{}).
		Where("id=? and member_id=? and status=?", id, memberID, model.CouponStatusInit).
		Updates(map[string]any{
			"status":   model.CouponStatusUsed,
			"used_at":  time.Now().Unix(),
			"order_id": orderID,
		})
	if result.Error != nil {
		return errors.Wrapf(result.Error, "[repo.couponMember] set used")
	}
	if result.RowsAffected == 0 { //没有记录更新
		return dbs.ErrRecordNotModified
	}

	return nil
}

// CouponUserSave 保存记录
func (r *Repo) CouponUserSave(ctx context.Context, coupon *model.CouponMemberModel) error {
	err := r.DB.WithContext(ctx).Save(coupon).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.couponMember] save db")
	}
	return nil
}

// GetCouponUserDrawIds 已领取的优惠券id列表
func (r *Repo) GetCouponUserDrawIds(ctx context.Context, memberID int64) (ids []int64, err error) {
	err = r.DB.WithContext(ctx).Model(&model.CouponMemberModel{}).Where("member_id=?", memberID).Pluck("coupon_id", &ids).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.couponMember] draw pluck db")
	}
	return
}

// CheckReceived 检查是否已领取过
func (r *Repo) CheckReceived(ctx context.Context, memberID, couponID int64) (bool, error) {
	var c int64
	err := r.DB.WithContext(ctx).Model(&model.CouponMemberModel{}).Where("member_id=? && coupon_id=?", memberID, couponID).Count(&c).Error
	if err != nil {
		return false, errors.Wrapf(err, "[repo.couponMember] count")
	}
	return c > 0, nil
}
