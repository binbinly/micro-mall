package model

import (
	"pkg/dbs"
)

// MemberLevelModel 会员等级
type MemberLevelModel struct {
	dbs.PriID
	Name          string `json:"name" gorm:"column:name;not null;type:varchar(100);comment:等级名称"`
	Growth        int    `json:"growth" gorm:"column:growth;not null;type:int;comment:等级需要的成长值"`
	FreeFreight   int    `json:"free_freight" gorm:"column:free_freight;not null;type:int;comment:免运费标准"`
	CommentGrowth int    `json:"comment_growth" gorm:"column:comment_growth;not null;type:int;comment:评论获取的成长值"`
	IsFreight     int8   `json:"is_freight" gorm:"is_freight;not null;default:0;comment:是否有免邮特权"`
	IsMember      int8   `json:"is_member" gorm:"is_member;not null;default:0;comment:是否有会员价格特权"`
	IsBirthday    int8   `json:"is_birthday" gorm:"is_birthday;not null;default:0;comment:是否有生日特权"`
	IsDefault     int8   `json:"is_default" gorm:"is_default;not null;default:0;comment:是否默认等级"`
	Note          string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	dbs.CUT
}

// TableName 表名
func (u *MemberLevelModel) TableName() string {
	return "ums_member_level"
}
