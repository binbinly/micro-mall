package model

import "pkg/dbs"

const (
	// CouponStatusNotReceive 未领取
	CouponStatusNotReceive = iota
	// CouponStatusInit 未使用
	CouponStatusInit
	// CouponStatusUsed 已使用
	CouponStatusUsed
	// CouponStatusExpire 已过期
	CouponStatusExpire
)

const (
	// CouponGetTypeDraw 主动领取
	CouponGetTypeDraw = 1
)

// CouponMemberModel 优惠券会员领取
type CouponMemberModel struct {
	dbs.PriID
	dbs.MID
	CouponID int64  `json:"coupon_id" gorm:"column:coupon_id;not null;index;type:int;comment:优惠券id"`
	Nickname string `json:"nickname" gorm:"column:nickname;not null;type:varchar(64);default:'';comment:会员昵称"`
	GetType  int8   `json:"get_type" gorm:"column:get_type;not null;comment:获取方式[0->后台赠送；1->主动领取]"`
	Status   int8   `json:"status" gorm:"column:status;not null;default:1;comment:使用状态[1->未使用；2->已使用；3->已过期]"`
	UsedAt   int64  `json:"used_at" gorm:"column:used_at;not null;type:int;default:0;comment:使用时间"`
	OrderID  int64  `json:"order_id" gorm:"column:order_id;not null;type:int;default:0;comment:订单id"`
	dbs.CUT
}

// TableName 表名
func (u *CouponMemberModel) TableName() string {
	return "sms_coupon_member"
}
