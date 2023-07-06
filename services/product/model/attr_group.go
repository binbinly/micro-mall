package model

import (
	"pkg/dbs"
)

// AttrGroupModel 属性分组
type AttrGroupModel struct {
	dbs.PriID
	dbs.Cat
	Name string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:分组名"`
	Desc string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	Icon string `json:"icon" gorm:"column:icon;not null;type:varchar(128);default:'';comment:图标"`
	dbs.OrderBy
}

// TableName 表名
func (u *AttrGroupModel) TableName() string {
	return "pms_attr_group"
}
