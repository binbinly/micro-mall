package model

import (
	"pkg/mysql"
)

// CategoryRelBrandModel 分类关联品牌
type CategoryRelBrandModel struct {
	mysql.PriID
	mysql.Cat
	BrandID int `json:"brand_id" gorm:"column:brand_id;not null;type:int(11) unsigned;comment:品牌id"`
}

// TableName 表名
func (u *CategoryRelBrandModel) TableName() string {
	return "pms_category_rel_brand"
}
