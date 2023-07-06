package model

import "pkg/dbs"

// WareTaskDetailModel 库存工作单详情
type WareTaskDetailModel struct {
	dbs.PriID
	TaskID int64 `json:"task_id" gorm:"column:task_id;not null;type:int;comment:工作单id"`
	WareID int64 `json:"ware_id" gorm:"column:ware_id;not null;type:int;comment:仓库id"`
	dbs.Sku
	SkuName string `json:"sku_name" gorm:"column:sku_name;not null;type:varchar(255);comment:采购商品名"`
	SkuNum  int    `json:"sku_num" gorm:"column:sku_num;not null;type:int;comment:采购数量"`
	dbs.CUT
}

// TableName 表名
func (u *WareTaskDetailModel) TableName() string {
	return "wms_ware_task_detail"
}
