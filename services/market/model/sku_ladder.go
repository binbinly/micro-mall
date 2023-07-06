package model

import "pkg/dbs"

// SkuLadderModel 商品阶梯价格
type SkuLadderModel struct {
	dbs.PriID
	dbs.Sku
	FullCount int  `json:"full_count" gorm:"column:full_count;not null;type:int;comment:满几件"`
	Discount  int  `json:"discount" gorm:"column:discount;not null;type:int;comment:打几折"`
	Price     int  `json:"price" gorm:"column:price;not null;type:int;comment:折后价"`
	IsSuper   int8 `json:"is_super" gorm:"column:is_super;not null;default:0;comment:是否叠加优惠"`
}

// TableName 表名
func (u *SkuLadderModel) TableName() string {
	return "sms_sku_ladder"
}
