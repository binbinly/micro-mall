package model

import "pkg/dbs"

// HomeSubjectModel 首页专题模型【首页下面很多专题，每个专题链接新的页面，展示专题商品信息】
type HomeSubjectModel struct {
	dbs.PriID
	Title    string `json:"title" gorm:"column:title;not null;type:varchar(128);comment:标题"`
	Subtitle string `json:"subtitle" gorm:"column:subtitle;not null;type:varchar(128);comment:副标题"`
	Img      string `json:"img" gorm:"column:img;not null;type:varchar(128);comment:图片"`
	Status   int8   `json:"status" gorm:"column:status;not null;default:0;comment:状态"`
	Url      string `json:"url" gorm:"column:url;not null;type:varchar(128);default:'';comment:链接地址"`
	Note     string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	dbs.CUBy
	dbs.CUT
}

// TableName 表名
func (u *HomeSubjectModel) TableName() string {
	return "sms_home_subject"
}
