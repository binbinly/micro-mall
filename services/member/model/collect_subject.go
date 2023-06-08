package model

import "pkg/mysql"

// CollectSubjectModel 用户收藏专题活动
type CollectSubjectModel struct {
	mysql.PriID
	mysql.MID
	SubjectID   int    `json:"subject_id" gorm:"column:subject_id;not null;type:int;comment:活动id"`
	SubjectName string `json:"subject_name" gorm:"column:subject_name;not null;type:varchar(255);comment:活动名"`
	SubjectImg  string `json:"subject_img" gorm:"column:subject_img;not null;type:varchar(255);comment:活动封面"`
	SubjectUrl  string `json:"subject_url" gorm:"column:subject_url;not null;type:varchar(255);comment:活动url"`
	mysql.CT
	mysql.DT
}

// TableName 表名
func (u *CollectSubjectModel) TableName() string {
	return "ums_collect_subject"
}
