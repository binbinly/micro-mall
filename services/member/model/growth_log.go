package model

import "pkg/mysql"

// GrowthLogModel 用户成长值变化记录模型
type GrowthLogModel struct {
	mysql.PriID
	mysql.MID
	ChangeCount int    `json:"change_count" gorm:"column:change_count;type:int;not null;comment:改变的值（正负计数）"`
	Note        string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	SourceType  int8   `json:"source_type" gorm:"column:source_type;not null;type:tinyint unsigned;default:0;comment:积分来源[0-购物，1-管理员修改]"`
	mysql.CT
}

// TableName 表名
func (u *GrowthLogModel) TableName() string {
	return "ums_growth_log"
}
