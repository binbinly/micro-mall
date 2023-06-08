package model

import "pkg/mysql"

// AttrRelGroupModel 属性关联属性分组
type AttrRelGroupModel struct {
	mysql.PriID
	AttrID  int `json:"attr_id" gorm:"column:attr_id;not null;type:int(11) unsigned;comment:属性id"`
	GroupID int `json:"group_id" gorm:"column:group_id;not null;type:int(11) unsigned;comment:属性分组id"`
}

// TableName 表名
func (u *AttrRelGroupModel) TableName() string {
	return "pms_attr_rel_group"
}
