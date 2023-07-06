package model

import (
	"pkg/dbs"
)

// SpuDescModel spu介绍
type SpuDescModel struct {
	dbs.PriID
	dbs.Spu
	Content string `json:"content" gorm:"column:content;comment:商品介绍"`
}

// TableName 表名
func (u *SpuDescModel) TableName() string {
	return "pms_spu_desc"
}
