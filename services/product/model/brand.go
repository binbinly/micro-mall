package model

import (
	"pkg/dbs"
)

// BrandModel 品牌
type BrandModel struct {
	dbs.PriID
	Name  string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:品牌"`
	Desc  string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	Logo  string `json:"logo" gorm:"column:logo;not null;type:varchar(128);default:'';comment:品牌logo"`
	Cover string `json:"cover" gorm:"column:cover;not null;type:varchar(128);default:'';comment:品牌封面"`
	dbs.Release
	dbs.OrderBy
	dbs.CUT
}

// TableName 表名
func (u *BrandModel) TableName() string {
	return "pms_brand"
}

// Brand 品牌结构
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
