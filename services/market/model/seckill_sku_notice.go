package model

import "pkg/dbs"

// SeckillSkuNoticeModel 秒杀商品通知订阅
type SeckillSkuNoticeModel struct {
	dbs.PriID
	dbs.Sku
	MemberID  int   `json:"member_id" gorm:"column:member_id;not null;type:int;comment:会员id"`
	SessionID int   `json:"session_id" gorm:"column:session_id;not null;type:int;comment:场次id"`
	SendAt    int64 `json:"send_at" gorm:"column:send_at;not null;type:int;comment:发送时间"`
	Type      int8  `json:"type" gorm:"column:type;not null;default:0;comment:通知方式[0-短信，1-邮件]"`
	dbs.CT
}

// TableName 表名
func (u *SeckillSkuNoticeModel) TableName() string {
	return "sms_seckill_sku_notice"
}
