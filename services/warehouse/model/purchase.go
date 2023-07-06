package model

import "pkg/dbs"

// PurchaseModel 采购信息
type PurchaseModel struct {
	dbs.PriID
	AssigneeID   int    `json:"assignee_id" gorm:"column:assignee_id;not null;type:int;default:0;comment:采购人id"`
	AssigneeName string `json:"assignee_name" gorm:"column:assignee_name;not null;type:varchar(64);default:'';comment:采购人"`
	Phone        string `json:"phone" gorm:"column:phone;not null;type:char(11);default:'';comment:手机号"`
	Priority     int8   `json:"priority" gorm:"column:priority;not null;default:0;comment:优先级"`
	Status       int8   `json:"status" gorm:"column:status;not null;default:0;comment:状态"`
	WareID       int    `json:"ware_id" gorm:"column:ware_id;not null;type:int;comment:仓库id"`
	Amount       int    `json:"amount" gorm:"column:amount;not null;type:int;default:0;comment:总金额/分"`
	dbs.CUT
}

// TableName 表名
func (u *PurchaseModel) TableName() string {
	return "wms_purchase"
}
