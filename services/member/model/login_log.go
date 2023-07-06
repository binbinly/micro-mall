package model

import (
	"pkg/dbs"
)

// LoginLogModel 会员登录日志
type LoginLogModel struct {
	dbs.PriID
	dbs.MID
	IP   string `json:"ip" gorm:"column:ip;not null;type:varchar(15);comment:ip地址"`
	City string `json:"city" gorm:"column:city;not null;type:varchar(64);default:'';comment:城市"`
	Type int8   `json:"type" gorm:"column:type;not null;comment:登录类型[1-web，2-app]"`
	dbs.CT
}

// TableName 表名
func (u *LoginLogModel) TableName() string {
	return "ums_login_log"
}
