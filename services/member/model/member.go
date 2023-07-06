package model

import (
	"pkg/dbs"
)

const (
	// MemberStatusNormal 状态 - 正常
	MemberStatusNormal = iota + 1
	// MemberStatusDisable 状态 - 禁用
	MemberStatusDisable
)

// MemberModel 会员模型
type MemberModel struct {
	dbs.PriID
	LevelID     int    `json:"level_id" gorm:"column:level_id;type:int;not null;default:0;comment:会员等级id"`
	Username    string `json:"username" gorm:"column:username;not null;uniqueIndex;type:varchar(64);comment:用户名"`
	Nickname    string `json:"nickname" gorm:"column:nickname;not null;type:varchar(64);default:'';comment:昵称"`
	Phone       int64  `gorm:"column:phone;not null;uniqueIndex;comment:手机号" json:"phone"`
	Email       string `gorm:"column:email;not null;type:varchar(60);default:'';comment:邮箱" json:"email"`
	Avatar      string `gorm:"column:avatar;not null;type:varchar(128);default:'';comment:头像" json:"avatar"`
	Gender      int8   `gorm:"column:gender;not null;default:1;comment:性别" json:"gender"`
	Birth       string `json:"birth" gorm:"column:birth;type:date;default:null;comment:生日"`
	Area        string `json:"area" gorm:"column:area;not null;type:varchar(255);default:'';comment:城市"`
	Job         string `json:"job" gorm:"column:job;not null;type:varchar(255);default:'';comment:职业"`
	SourceType  int8   `json:"source_type" gorm:"column:source_type;not null;default:0;comment:用户来源"`
	Integration int    `json:"integration" gorm:"column:integration;not null;default:0;comment:积分"`
	Growth      int    `json:"growth" gorm:"column:growth;not null;default:0;comment:成长值"`
	Status      int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	Sign        string `gorm:"column:sign;not null;type:varchar(255);default:'';comment:签名" json:"sign"`
	dbs.CUT
}

// TableName 表名
func (u *MemberModel) TableName() string {
	return "ums_member"
}
