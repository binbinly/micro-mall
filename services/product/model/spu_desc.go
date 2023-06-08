package model

import (
	"pkg/mysql"
)

// SpuDescModel spu介绍
type SpuDescModel struct {
	mysql.PriID
	mysql.Spu
	Content string `json:"content" gorm:"column:content;comment:商品介绍"`
}

// TableName 表名
func (u *SpuDescModel) TableName() string {
	return "pms_spu_desc"
}
