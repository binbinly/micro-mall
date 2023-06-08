package model

import "pkg/mysql"

// PurchaseDetailModel 采购详情
type PurchaseDetailModel struct {
	mysql.PriID
	mysql.Sku
	PurchaseID int  `json:"purchase_id" gorm:"column:purchase_id;not null;type:int;default:0;comment:采购单id"`
	SkuNum     int  `json:"sku_num" gorm:"column:sku_num;not null;type:int;comment:采购数量"`
	SkuPrice   int  `json:"sku_price" gorm:"column:sku_price;not null;type:int;default:0;comment:采购金额/分"`
	WareID     int  `json:"ware_id" gorm:"column:ware_id;not null;type:int;comment:仓库id"`
	Status     int8 `json:"status" gorm:"column:status;not null;default:0;comment:状态[0新建，1已分配，2正在采购，3已完成，4采购失败]"`
	mysql.CUT
}

// TableName 表名
func (u *PurchaseDetailModel) TableName() string {
	return "wms_purchase_detail"
}
