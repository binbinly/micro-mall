package model

import "pkg/dbs"

// SeckillSessionModel 秒杀活动场次
type SeckillSessionModel struct {
	dbs.PriID
	Name    string `json:"name" gorm:"column:name;not null;type:varchar(128);comment:场次名"`
	StartAt int64  `json:"start_at" gorm:"column:start_at;not null;type:int;comment:开始时间"`
	EndAt   int64  `json:"end_at" gorm:"column:end_at;not null;type:int;comment:结束时间"`
	dbs.Release
	dbs.CUT
}

// TableName 表名
func (u *SeckillSessionModel) TableName() string {
	return "sms_seckill_session"
}
