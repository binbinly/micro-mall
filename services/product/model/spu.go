package model

import (
	"pkg/mysql"
)

// SpuModel spu信息
type SpuModel struct {
	mysql.PriID
	mysql.Cat
	BrandID       int    `json:"brand_id" gorm:"column:brand_id;not null;type:int(11) unsigned;comment:品牌id"`
	Name          string `json:"name" gorm:"column:name;not null;type:varchar(255);comment:商品名"`
	OriginalPrice int    `json:"original_price" gorm:"column:original_price;not null;type:int unsigned;comment:原价"`
	Desc          string `json:"desc" gorm:"column:desc;not null;type:varchar(1000);default:'';comment:商品描述"`
	Weight        int    `json:"weight" gorm:"column:weight;not null;type:int(11) unsigned;comment:重量/g"`
	IsMany        int8   `json:"is_many" gorm:"column:is_many;not null;type:tinyint unsigned;default:0;comment:是否多规格"`
	Status        int8   `json:"status" gorm:"column:status;not null;default:0;comment:状态"`
	mysql.CUT
}

// TableName 表名
func (u *SpuModel) TableName() string {
	return "pms_spu"
}
