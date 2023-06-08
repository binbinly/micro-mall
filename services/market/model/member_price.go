package model

import "pkg/mysql"

// MemberPriceModel 商品会员价格
type MemberPriceModel struct {
	mysql.PriID
	mysql.Sku
	LevelID   int    `json:"level_id" gorm:"column:level_id;not null;type:int;comment:会员等级"`
	LevelName string `json:"level_name" gorm:"column:level_name;not null;type:varchar(64);comment:会员等级名"`
	Price     int    `json:"price" gorm:"column:price;not null;type:int;comment:会员对应价格"`
	IsSuper   int8   `json:"is_super" gorm:"column:is_super;not null;default:0;comment:是否叠加优惠"`
}

// TableName 表名
func (u *MemberPriceModel) TableName() string {
	return "sms_member_price"
}
