package model

import "pkg/dbs"

// AttrRelGroupModel 属性关联属性分组
type AttrRelGroupModel struct {
	dbs.PriID
	AttrID  int `json:"attr_id" gorm:"column:attr_id;not null;type:int;comment:属性id"`
	GroupID int `json:"group_id" gorm:"column:group_id;not null;type:int;comment:属性分组id"`
}

// TableName 表名
func (u *AttrRelGroupModel) TableName() string {
	return "pms_attr_rel_group"
}
