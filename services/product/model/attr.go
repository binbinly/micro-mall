package model

import (
	"pkg/mysql"
)

// AttrModel 商品属性
type AttrModel struct {
	mysql.PriID
	mysql.Cat
	Name     string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:属性名"`
	Icon     string `json:"icon" gorm:"column:icon;not null;type:varchar(128);default:'';comment:属性图标"`
	Values   string `json:"values" gorm:"column:values;not null;type:varchar(255);comment:可选值多个逗号连接"`
	Type     int8   `json:"type" gorm:"column:type;not null;comment:属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]"`
	IsMany   int8   `json:"is_many" gorm:"column:is_many;not null;default:0;comment:值类型，是否多个值"`
	IsSearch int8   `json:"is_search" gorm:"column:is_search;not null;default:0;comment:是否需要搜索"`
	IsShow   int8   `json:"is_show" gorm:"column:is_show;not null;default:0;comment:是否展示在介绍上"`
	mysql.Release
	mysql.CUT
}

// TableName 表名
func (u *AttrModel) TableName() string {
	return "pms_attr"
}
