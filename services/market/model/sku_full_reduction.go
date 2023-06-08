package model

import "pkg/mysql"

// SkuFullReductionModel 商品满减信息
type SkuFullReductionModel struct {
	mysql.PriID
	mysql.Sku
	FullPrice   int  `json:"full_price" gorm:"column:full_price;not null;type:int;comment:满多少"`
	ReducePrice int  `json:"reduce_price" gorm:"column:reduce_price;not null;type:int;comment:减多少"`
	IsSuper     int8 `json:"is_super" gorm:"column:is_super;not null;default:0;comment:是否叠加优惠"`
}

// TableName 表名
func (u *SkuFullReductionModel) TableName() string {
	return "sms_sku_full_reduction"
}
