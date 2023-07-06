package model

import (
	"pkg/dbs"
)

// SkuAttrModel sku销售属性&值
type SkuAttrModel struct {
	dbs.PriID
	dbs.Sku
	AttrID    int64  `json:"attr_id" gorm:"column:attr_id;not null;type:int;comment:属性id"`
	AttrName  string `json:"attr_name" gorm:"column:attr_name;not null;type:varchar(60);comment:销售属性名"`
	AttrValue string `json:"attr_value" gorm:"column:attr_value;not null;type:varchar(255);comment:销售属性值"`
	dbs.OrderBy
}

// TableName 表名
func (u *SkuAttrModel) TableName() string {
	return "pms_sku_attr"
}
