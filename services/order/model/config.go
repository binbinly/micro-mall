package model

import (
	"pkg/mysql"
)

const (
	// ConfigFlashOrderOvertime 秒杀订单超时关闭时间(分)
	ConfigFlashOrderOvertime = "flash_order_overtime"
	// ConfigNormalOrderOvertime 正常订单超时时间(分)
	ConfigNormalOrderOvertime = "normal_order_overtime"
	// ConfigConfirmOvertime 发货后自动确认收货时间（天）
	ConfigConfirmOvertime = "confirm_overtime"
	// ConfigFinishOvertime 自动完成交易时间，不能申请退货（天）
	ConfigFinishOvertime = "finish_overtime"
	// ConfigCommentOvertime 订单完成后自动好评时间（天）
	ConfigCommentOvertime = "comment_overtime"
	// ConfigMemberLevel 会员等级【0-不限会员等级，全部通用；其他-对应的其他会员等级】
	ConfigMemberLevel = "member_level"
)

// ConfigModel 配置模型
type ConfigModel struct {
	mysql.PriID
	Name  string `json:"name" gorm:"column:name;not null;type:char(15);comment:配置键"`
	Value string `json:"value" gorm:"column:value;not null;type:varchar(1000);comment:配置值"`
	Desc  string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:描述"`
	mysql.CUBy
	mysql.CUT
}

// TableName 表名
func (u *ConfigModel) TableName() string {
	return "oms_config"
}

// Config 对外配置结构
type Config struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
