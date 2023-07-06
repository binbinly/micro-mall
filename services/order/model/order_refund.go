package model

import (
	"pkg/dbs"
)

// OrderRefundModel 退款模型
type OrderRefundModel struct {
	dbs.PriID
	OID
	Amount  int    `json:"amount" gorm:"column:amount;not null;type:int;comment:退款金额"`
	TradeNo string `json:"trade_no" gorm:"column:trade_no;not null;type:varchar(32);default:'';comment:交易号"`
	Status  int8   `json:"status" gorm:"column:status;not null;default:0;comment:订单状态"`
	Channel int8   `json:"channel" gorm:"column:channel;not null;default:0;comment:退款渠道[1-支付宝，2-微信，3-银联，4-汇款]"`
	Content string `json:"content" gorm:"column:content;not null;type:varchar(255);default:'';comment:退款内容"`
	dbs.CUT
}

// TableName 表名
func (u *OrderRefundModel) TableName() string {
	return "oms_order_refund"
}
