package model

import "pkg/dbs"

// CouponRelSpuModel 优惠券关联产品
type CouponRelSpuModel struct {
	dbs.PriID
	CouponID int `json:"coupon_id" gorm:"column:coupon_id;not null;index;type:int;comment:优惠券id"`
	dbs.Spu
	SpuName string `json:"spu_name" gorm:"column:spu_name;not null;type:varchar(255);comment:产品名"`
}

// TableName 表名
func (u *CouponRelSpuModel) TableName() string {
	return "sms_coupon_rel_spu"
}
