package model

import "pkg/dbs"

// SeckillSkuModel 秒杀活动关联商品
type SeckillSkuModel struct {
	dbs.PriID
	dbs.Sku
	ActivityID int `json:"activity_id" gorm:"column:activity_id;not null;type:int;comment:活动id"`
	SessionID  int `json:"session_id" gorm:"column:session_id;not null;type:int;comment:场次id"`
	Price      int `json:"price" gorm:"column:price;not null;type:int;comment:秒杀价格"`
	Count      int `json:"count" gorm:"column:count;not null;type:int;comment:秒杀总量"`
	Limit      int `json:"limit" gorm:"column:limit;not null;type:int;default:1;comment:每人限购数量"`
	dbs.OrderBy
}

// TableName 表名
func (u *SeckillSkuModel) TableName() string {
	return "sms_seckill_sku"
}
