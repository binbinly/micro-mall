package model

import (
	"pkg/dbs"
)

// PaymentModel 支付模型
type PaymentModel struct {
	dbs.PriID
	OID
	ONo
	TradeNo         string `json:"trade_no" gorm:"column:trade_no;not null;type:varchar(32);comment:交易号"`
	Amount          int    `json:"amount" gorm:"column:amount;not null;type:int;comment:交易金额/分"`
	Subject         string `json:"subject" gorm:"column:subject;not null;type:varchar(255);comment:交易内容"`
	ConfirmAt       int64  `json:"confirm_at" gorm:"column:confirm_at;not null;default:0;comment:确认时间"`
	CallbackContent string `json:"callback_content" gorm:"column:callback_content;not null;type:varchar(3000);default:'';comment:回调内容"`
	CallbackAt      int64  `json:"callback_at" gorm:"column:callback_at;not null;default:0;comment:回调时间"`
	dbs.CUT
}

// TableName 表名
func (u *PaymentModel) TableName() string {
	return "oms_payment"
}
