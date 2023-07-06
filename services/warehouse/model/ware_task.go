package model

import "pkg/dbs"

const (
	//TaskStatusLock 锁定
	TaskStatusLock = 1 + iota
	//TaskStatusUnlock 解锁
	TaskStatusUnlock
	//TaskStatusFinish 订单完成扣减库存
	TaskStatusFinish
)

// WareTaskModel 库存工作单
type WareTaskModel struct {
	dbs.PriID
	OrderID   int64  `json:"order_id" gorm:"column:order_id;not null;index;type:int;comment:order_id"`
	OrderNo   string `json:"order_no" gorm:"column:order_no;not null;index;type:varchar(32);comment:单号"`
	Consignee string `json:"consignee" gorm:"column:consignee;not null;type:varchar(64);comment:收货人"`
	Phone     string `json:"phone" gorm:"column:phone;not null;type:char(11);comment:收货人手机号"`
	Address   string `json:"address" gorm:"column:address;not null;type:varchar(255);comment:配送地址"`
	Note      string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	Remark    string `json:"remark" gorm:"column:remark;not null;type:varchar(255);default:'';comment:任务备注"`
	Status    int8   `json:"status" gorm:"column:status;not null;default:0;comment:状态 1=锁定 2=解锁 3=扣减"`
	dbs.CUT
}

// TableName 表名
func (u *WareTaskModel) TableName() string {
	return "wms_ware_task"
}
