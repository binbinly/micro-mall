package model

import (
	"pkg/mysql"
)

// OrderReturnReasonModel 退货原因
type OrderReturnReasonModel struct {
	mysql.PriID
	Name   string `json:"name" gorm:"column:name;not null;type:varchar(255);comment:退货原因名"`
	Status int8   `json:"status" gorm:"column:status;not null;default:0;comment:订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】"`
	mysql.CUBy
	mysql.CUT
}

// TableName 表名
func (u *OrderReturnReasonModel) TableName() string {
	return "oms_order_return_reason"
}
