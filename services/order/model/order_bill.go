package model

import (
	"pkg/dbs"
)

// OrderBillModel 订单发票
type OrderBillModel struct {
	dbs.PriID
	dbs.MID
	OID
	Type        int8   `json:"type" gorm:"column:type;not null;type:tinyint;comment:类型：0:个人,1:企业"`
	Name        string `json:"name" gorm:"column:name;not null;type:varchar(120);comment:名称/公司名称"`
	Phone       string `json:"phone" gorm:"column:phone;not null;type:char(11);comment:手机号"`
	Email       string `json:"email" gorm:"column:email;not null;type:varchar(120);comment:邮箱"`
	Code        string `json:"code" gorm:"column:code;not null;type:varchar(60);default:'';comment:税号"`
	Address     string `json:"address" gorm:"column:address;not null;type:varchar(120);comment:地址"`
	BankName    string `json:"bank_name" gorm:"column:bank_name;not null;type:varchar(60);comment:开户行"`
	BankAccount string `json:"bank_account" gorm:"column:bank_account;not null;type:varchar(30);comment:银行账号"`
	IsFinish    int8   `json:"is_finish" gorm:"column:is_finish;not null;type:tinyint;default:0;comment:是否开票"`
	dbs.CUT
}

// TableName 表名
func (o *OrderBillModel) TableName() string {
	return "oms_order_bill"
}
