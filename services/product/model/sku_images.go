package model

import (
	"pkg/dbs"
)

// SkuImageModel sku图片
type SkuImageModel struct {
	dbs.PriID
	dbs.Sku
	Img       string `json:"img" gorm:"column:img;not null;type:varchar(128);comment:图片地址"`
	IsDefault int8   `json:"is_default" gorm:"column:is_default;not null;default:0;comment:是否默认"`
	dbs.OrderBy
}

// TableName 表名
func (u *SkuImageModel) TableName() string {
	return "pms_sku_image"
}
