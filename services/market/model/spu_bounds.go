package model

import "pkg/mysql"

// SpuBoundsModel 商品spu积分设置
type SpuBoundsModel struct {
	mysql.PriID
	mysql.Spu
	GrowBounds int  `json:"grow_bounds" gorm:"column:grow_bounds;not null;type:int;comment:成长积分"`
	BuyBounds  int  `json:"buy_bounds" gorm:"column:buy_bounds;not null;type:int;comment:购物积分"`
	Work       int8 `json:"work" gorm:"column:work;not null;default:0;comment:优惠生效情况[0 - 无优惠，成长积分是否赠送;1 - 无优惠，购物积分是否赠送;2 - 有优惠，成长积分是否赠送;3 - 有优惠，购物积分是否赠送【状态位0：不赠送，1：赠送】]"`
}

// TableName 表名
func (u *SpuBoundsModel) TableName() string {
	return "sms_spu_bounds"
}
