package model

import "pkg/dbs"

// CouponRelCatModel 优惠券关联分类
type CouponRelCatModel struct {
	dbs.PriID
	CouponID int `json:"coupon_id" gorm:"column:coupon_id;not null;index;type:int;comment:优惠券id"`
	dbs.Cat
	CatName string `json:"cat_name" gorm:"column:cat_name;not null;type:varchar(255);comment:产品分类名"`
}

// TableName 表名
func (u *CouponRelCatModel) TableName() string {
	return "sms_coupon_rel_cat"
}
