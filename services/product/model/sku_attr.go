package model

import (
	"pkg/mysql"
)

// SkuAttrModel sku销售属性&值
type SkuAttrModel struct {
	mysql.PriID
	mysql.Sku
	AttrID    int64  `json:"attr_id" gorm:"column:attr_id;not null;type:int(11) unsigned;comment:属性id"`
	AttrName  string `json:"attr_name" gorm:"column:attr_name;not null;type:varchar(60);comment:销售属性名"`
	AttrValue string `json:"attr_value" gorm:"column:attr_value;not null;type:varchar(255);comment:销售属性值"`
	mysql.OrderBy
}

// TableName 表名
func (u *SkuAttrModel) TableName() string {
	return "pms_sku_attr"
}
