package model

import (
	"pkg/mysql"
)

// AttrValueModel spu属性值
type AttrValueModel struct {
	mysql.PriID
	mysql.Spu
	AttrID    int    `json:"attr_id" gorm:"column:attr_id;not null;type:int(11) unsigned;comment:属性id"`
	AttrName  string `json:"attr_name" gorm:"column:attr_name;not null;type:varchar(100);comment:属性名"`
	AttrValue string `json:"attr_value" gorm:"column:attr_value;not null;type:varchar(255);comment:属性值"`
	IsShow    int8   `json:"is_show" gorm:"column:is_show;not null;default:0;comment:是否展示在介绍上"`
	mysql.OrderBy
}

// TableName 表名
func (u *AttrValueModel) TableName() string {
	return "pms_attr_value"
}

// Attrs 规格属性结构
type Attrs struct {
	GroupID   int64  `json:"group_id"`
	AttrID    int64  `json:"attr_id"`
	AttrName  string `json:"attr_name"`
	AttrValue string `json:"attr_value"`
}
