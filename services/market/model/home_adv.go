package model

import "pkg/mysql"

// HomeAdvModel 首页轮播广告
type HomeAdvModel struct {
	mysql.PriID
	Title      string `json:"title" gorm:"column:title;not null;type:varchar(128);comment:标题"`
	Img        string `json:"img" gorm:"column:img;not null;type:varchar(128);comment:图片"`
	StartAt    int64  `json:"start_at" gorm:"column:start_at;not null;type:int;comment:开始时间"`
	EndAt      int64  `json:"end_at" gorm:"column:end_at;not null;type:int;comment:结束时间"`
	ClickCount int    `json:"click_count" gorm:"column:click_count;not null;type:int;default:0;comment:点击数"`
	Url        string `json:"url" gorm:"column:url;not null;type:varchar(128);default:'';comment:链接地址"`
	Note       string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	mysql.Release
	mysql.OrderBy
	mysql.CUBy
	mysql.CUT
}

// TableName 表名
func (u *HomeAdvModel) TableName() string {
	return "sms_home_adv"
}
