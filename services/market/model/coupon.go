package model

import "pkg/dbs"

const (
	// CouponUseTypeAll 全场通用
	CouponUseTypeAll = iota
	// CouponUseTypeCat 指定分类
	CouponUseTypeCat
	// CouponUseTypeSpu 指定商品
	CouponUseTypeSpu
)

// CouponModel 优惠券
type CouponModel struct {
	dbs.PriID
	Name          string `json:"name" gorm:"column:name;not null;type:varchar(128);comment:优惠券名"`
	Cover         string `json:"cover" gorm:"column:cover;not null;type:varchar(128);default:'';comment:优惠券封面"`
	Type          int8   `json:"type" gorm:"column:type;not null;comment:优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]"`
	Num           int    `json:"num" gorm:"column:num;not null;type:int;comment:数量"`
	Amount        int    `json:"amount" gorm:"column:amount;not null;type:int;comment:金额/分"`
	PerLimit      int    `json:"per_limit" gorm:"column:per_limit;not null;type:int;default:1;comment:每人限领张数"`
	MinPoint      int    `json:"min_point" gorm:"column:min_point;not null;type:int;comment:使用门槛"`
	StartAt       int64  `json:"start_at" gorm:"column:start_at;not null;type:int;comment:开始时间"`
	EndAt         int64  `json:"end_at" gorm:"column:end_at;not null;type:int;comment:结束时间"`
	UseType       int8   `json:"use_type" gorm:"column:use_type;not null;comment:使用类型[0->全场通用；1->指定分类；2->指定商品]"`
	Note          string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	PublishCount  int    `json:"publish_count" gorm:"column:publish_count;not null;type:int;default:0;comment:发行数量"`
	UseCount      int    `json:"use_count" gorm:"column:use_count;not null;type:int;default:0;comment:已使用数量"`
	ReceiveCount  int    `json:"receive_count" gorm:"column:receive_count;not null;type:int;default:0;comment:已领取数"`
	EnableStartAt int    `json:"enable_start_at" gorm:"column:enable_start_at;not null;type:int;default:0;comment:可以领取的开始时间"`
	EnableEndAt   int    `json:"enable_end_at" gorm:"column:enable_end_at;not null;type:int;default:0;comment:可以领取的结束时间"`
	Code          string `json:"code" gorm:"column:code;not null;type:varchar(20);comment:优惠码"`
	MemberLevel   int8   `json:"member_level" gorm:"column:member_level;not null;default:0;comment:可以领取的会员等级[0->不限等级，其他-对应等级]"`
	dbs.Release
	dbs.CUT
}

// TableName 表名
func (u *CouponModel) TableName() string {
	return "sms_coupon"
}

// Coupon 领取的优惠券结构
type Coupon struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	MinPoint int    `json:"min_point"`
	StartAt  int64  `json:"start_at"`
	EndAt    int64  `json:"end_at"`
	Note     string `json:"note"`
	Status   int    `json:"status"`
}
