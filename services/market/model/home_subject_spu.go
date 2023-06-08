package model

import (
	"pkg/mysql"
)

// HomeSubjectSpuModel 专题商品
type HomeSubjectSpuModel struct {
	mysql.PriID
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(128);comment:名称"`
	SubjectID int    `json:"subject_id" gorm:"column:subject_id;not null;type:int;comment:专题id"`
	mysql.Spu
	mysql.OrderBy
}

// TableName 表名
func (u *HomeSubjectSpuModel) TableName() string {
	return "sms_home_subject_spu"
}
