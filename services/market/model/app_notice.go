package model

import "pkg/mysql"

// AppNoticeModel 公告模型
type AppNoticeModel struct {
	mysql.PriID
	Title   string `json:"title" gorm:"column:title;not null;type:varchar(255);comment:标题"`
	Content string `json:"content" gorm:"column:content;not null;type:varchar(1000);comment:内容"`
	mysql.CUBy
	mysql.CUT
}

// TableName 表名
func (u *AppNoticeModel) TableName() string {
	return "sms_app_notice"
}

// AppNotice 对外公告结构
type AppNotice struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}
