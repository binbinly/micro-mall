package model

import "pkg/dbs"

// WareSkuModel 商品库存
type WareSkuModel struct {
	dbs.PriID
	dbs.Sku
	WareID    int64  `json:"ware_id" gorm:"column:ware_id;not null;type:int;comment:仓库id"`
	SkuName   string `json:"sku_name" gorm:"column:sku_name;not null;type:varchar(255);default:'';comment:采购商品名"`
	Stock     int    `json:"stock" gorm:"column:stock;not null;type:int;default:0;comment:库存"`
	StockLock int    `json:"stock_lock" gorm:"column:stock_lock;not null;type:int;default:0;comment:锁定库存"`
	dbs.CUT
}

// TableName 表名
func (u *WareSkuModel) TableName() string {
	return "wms_ware_sku"
}

// WareSkuStock 商品库存对外结构
type WareSkuStock struct {
	SkuID     int64 `json:"sku_id"`
	Stock     int   `json:"stock"`
	StockLock int   `json:"stock_lock"`
}
