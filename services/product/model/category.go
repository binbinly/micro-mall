package model

import (
	"pkg/mysql"
)

// CategoryModel 商品三级分类
type CategoryModel struct {
	mysql.PriID
	Name         string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:分类名"`
	ParentID     int64  `json:"parent_id" gorm:"column:parent_id;not null;type:int(11) unsigned;default:0;comment:上级分类"`
	Level        int8   `json:"level" gorm:"column:level;not null;default:1;comment:层级"`
	Icon         string `json:"icon" gorm:"column:icon;not null;default:'';comment:图标"`
	ProductUnit  string `json:"product_unit" gorm:"column:product_unit;not null;default:'';comment:计量单位"`
	ProductCount int    `json:"product_count" gorm:"column:product_count;not null;type:int(11) unsigned;default:0;comment:商品数量"`
	mysql.Release
	mysql.OrderBy
}

// TableName 表名
func (u *CategoryModel) TableName() string {
	return "pms_category"
}

// CategoryEs es分类结构
type CategoryEs struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CategoryTree 分类树
type CategoryTree struct {
	ID       int64           `json:"id"`
	ParentID int64           `json:"parent_id"`
	Name     string          `json:"name"`
	Sort     int16           `json:"sort"`
	Child    []*CategoryTree `json:"child"`
}

// CategorySort 按照 sort 从大到小排序
type CategorySort []*CategoryTree

func (a CategorySort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a CategorySort) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a CategorySort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Sort < a[i].Sort
}
