package model

import "pkg/mysql"

// CollectSpuModel 用户收藏商品
type CollectSpuModel struct {
	mysql.PriID
	mysql.MID
	mysql.Spu
	SpuName string `json:"spu_name" gorm:"column:spu_name;not null;type:varchar(255);comment:spu_name"`
	SpuImg  string `json:"spu_img" gorm:"column:spu_img;not null;type:varchar(255);comment:spu_img"`
	mysql.CT
	mysql.DT
}

// TableName 表名
func (u *CollectSpuModel) TableName() string {
	return "ums_collect_spu"
}
