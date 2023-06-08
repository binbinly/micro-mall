package model

import (
	"pkg/mysql"
)

// MemberAddressModel 用户收货地址模型
type MemberAddressModel struct {
	mysql.PriID
	mysql.MID
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:收货人姓名"`
	Phone     string `json:"phone" gorm:"column:phone;not null;type:char(11);comment:收货人手机号"`
	Province  string `json:"province" gorm:"column:province;not null;type:varchar(30);comment:省"`
	City      string `json:"city" gorm:"column:city;not null;type:varchar(30);comment:市"`
	County    string `json:"county" gorm:"column:county;not null;type:varchar(30);comment:区/县"`
	AreaCode  int    `json:"area_code" gorm:"column:area_code;not null;type:int unsigned;comment:最后一级地区编码"`
	Detail    string `json:"detail" gorm:"column:detail;not null;type:varchar(255);comment:详细地址"`
	IsDefault int8   `json:"is_default" gorm:"column:is_default;not null;default:0;comment:是否默认"`
	mysql.CUT
	mysql.DT
}

// TableName 表名
func (u *MemberAddressModel) TableName() string {
	return "ums_member_address"
}
