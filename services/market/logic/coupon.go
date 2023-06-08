package logic

import (
	"context"
	"fmt"
	"github.com/binbinly/pkg/repo"
	"github.com/redis/go-redis/v9"
	"pkg/errno"
	"pkg/mysql"
	"time"

	"github.com/pkg/errors"

	"market/model"
)

// GetCouponList 优惠券列表
func (l *logic) GetCouponList(ctx context.Context, memberID, skuID int64) ([]*model.CouponModel, []int64, error) {
	list, err := l.repo.GetCouponList(ctx, 0, 0)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[logic.coupon] get list")
	}
	if len(list) == 0 {
		return []*model.CouponModel{}, nil, nil
	}
	//已领取的优惠券
	ids, err := l.repo.GetCouponUserDrawIds(ctx, memberID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[logic.coupon] get draw ids uid: %v", memberID)
	}
	return list, ids, nil
}

// GetMyCouponList 我的优惠券
func (l *logic) GetMyCouponList(ctx context.Context, memberID int64) ([]*model.Coupon, error) {
	list, err := l.repo.GetCouponMemberList(ctx, memberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[logic.coupon] get my list")
	}
	return list, nil
}

// CouponDraw 领取优惠券
func (l *logic) CouponDraw(ctx context.Context, memberID, id int64) error {
	coupon, err := l.repo.GetCouponByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[logic.coupon] find id: %v", id)
	}
	now := time.Now().Unix()
	if coupon == nil || coupon.ID == 0 || coupon.StartAt > now || coupon.EndAt < now {
		return errno.ErrCouponNotFound
	}
	exist, err := l.repo.CheckReceived(ctx, memberID, id)
	if err != nil {
		return errors.Wrapf(err, "[logic.coupon] check uid: %v id: %v", memberID, id)
	}
	if exist {
		return errno.ErrCouponReceived
	}
	num, err := l.rdb.Decr(ctx, fmt.Sprintf("coupon_num:%d", id)).Result()
	if err == redis.Nil {
		return errno.ErrCouponFinished
	} else if err != nil {
		return errors.Wrapf(err, "[logic.coupon] redis decr")
	}
	if num < 0 { //已领取完
		return errno.ErrCouponFinished
	}
	couponMember := &model.CouponMemberModel{
		MID:      mysql.MID{MemberID: memberID},
		CouponID: id,
		GetType:  model.CouponGetTypeDraw,
		Status:   model.CouponStatusInit,
	}
	err = l.repo.CouponUserSave(ctx, couponMember)
	if err != nil {
		return errors.Wrapf(err, "[logic.coupon] save")
	}
	return nil
}

// CouponUsed 使用优惠券
func (l *logic) CouponUsed(ctx context.Context, memberID, id, orderID int64) error {
	mCoupon, err := l.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[logic.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return errno.ErrCouponNotFound
	}
	err = l.repo.SetCouponMemberUsed(ctx, id, memberID, orderID)
	if errors.Is(err, repo.ErrRecordNotModified) {
		return errno.ErrCouponNotFound
	} else if err != nil {
		return errors.Wrapf(err, "[logic.coupon] set used id: %v uid: %v, oid: %v", id, memberID, orderID)
	}
	return nil
}

// GetCouponInfo 获取优惠券详情
func (l *logic) GetCouponInfo(ctx context.Context, memberID, id int64) (*model.CouponMemberModel, *model.CouponModel, error) {
	mCoupon, err := l.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[logic.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return nil, nil, errno.ErrCouponNotFound
	}
	coupon, err := l.repo.GetCouponByID(ctx, mCoupon.CouponID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[logic.coupon] get coupon by id: %v", mCoupon.CouponID)
	}
	now := time.Now().Unix()
	if coupon == nil || coupon.ID == 0 || now > coupon.EndAt || now < coupon.StartAt {
		return nil, nil, errno.ErrCouponNotFound
	}

	return mCoupon, coupon, nil
}
