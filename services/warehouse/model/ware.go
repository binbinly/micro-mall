package model

import "pkg/mysql"

// WareModel 仓库信息
type WareModel struct {
	mysql.PriID
	Name     string `json:"name" gorm:"column:name;not null;type:varchar(255);comment:仓库名"`
	Address  string `json:"address" gorm:"column:address;not null;type:varchar(255);comment:仓库地址"`
	AreaCode int    `json:"area_code" gorm:"column:area_code;not null;type:int unsigned;comment:最后一级地区编码"`
	mysql.CUT
}

// TableName 表名
func (u *WareModel) TableName() string {
	return "wms_ware"
}
