package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"pkg/mysql"
	"time"

	"github.com/pkg/errors"

	"market/model"
)

// ICoupon 优惠券接口
type ICoupon interface {
	GetCouponList(ctx context.Context, catID, spuID int64) (list []*model.CouponModel, err error)
	GetCouponByID(ctx context.Context, id int64) (coupon *model.CouponModel, err error)
	GetCouponMemberList(ctx context.Context, memberID int64) (list []*model.Coupon, err error)
	GetCouponMemberByID(ctx context.Context, id int64) (coupon *model.CouponMemberModel, err error)
	SetCouponMemberUsed(ctx context.Context, id, memberID, orderID int64) error
	CouponUserSave(ctx context.Context, coupon *model.CouponMemberModel) error
	GetCouponUserDrawIds(ctx context.Context, memberID int64) (ids []int64, err error)
	CheckReceived(ctx context.Context, memberID, couponID int64) (bool, error)
}

// GetCouponList 可以被领取的优惠券列表
func (r *Repo) GetCouponList(ctx context.Context, catID, spuID int64) (list []*model.CouponModel, err error) {
	now := time.Now().Unix()
	err = r.DB.WithContext(ctx).Model(&model.CouponModel{}).Scopes(mysql.WhereRelease).
		Where("enable_start_at<=?", now).Where("enable_end_at>=?", now).
		Where("use_type=?", model.CouponUseTypeAll).
		Order("amount desc,id desc").Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] by db list")
	}
	return
}

// GetCouponByID 获取优惠券详情
func (r *Repo) GetCouponByID(ctx context.Context, id int64) (coupon *model.CouponModel, err error) {
	doKey := fmt.Sprintf("coupon:%d", id)
	if err = r.QueryCache(ctx, doKey, &coupon, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).First(data, id).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] query cache")
	}

	return coupon, nil
}
