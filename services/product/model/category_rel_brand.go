package model

import (
	"pkg/dbs"
)

// CategoryRelBrandModel 分类关联品牌
type CategoryRelBrandModel struct {
	dbs.PriID
	dbs.Cat
	BrandID int `json:"brand_id" gorm:"column:brand_id;not null;type:int;comment:品牌id"`
}

// TableName 表名
func (u *CategoryRelBrandModel) TableName() string {
	return "pms_category_rel_brand"
}
