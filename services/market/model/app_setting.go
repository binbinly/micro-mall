package model

import (
	"encoding/json"
	"pkg/dbs"
)

const (
	// AppPageHome 首页
	AppPageHome = 1 + iota
	// AppPageSearch 搜索页
	AppPageSearch
)

const (
	// AppTypeSwiper 轮播图
	AppTypeSwiper = iota + 1
	// AppTypeNav 图标
	AppTypeNav
	// AppTypeThreeAdv 三图广告
	AppTypeThreeAdv
	// AppTypeOneAdv 单图广告
	AppTypeOneAdv
	// AppTypeProduct 商品列表
	AppTypeProduct
)

// AppSettingModel 应用页面设置模型
type AppSettingModel struct {
	dbs.PriID
	dbs.Cat
	Page int8   `json:"page" gorm:"column:page;not null;type:tinyint;comment:页面"`
	Type int8   `json:"type" gorm:"column:type;not null;type:tinyint;comment:类型"`
	Data string `json:"data" gorm:"column:data;not null;type:varchar(5000);default:'';comment:数据"`
	dbs.OrderBy
	dbs.CUBy
	dbs.CUT
}

// TableName 表名
func (u *AppSettingModel) TableName() string {
	return "sms_app_setting"
}

// AppSetting 对外页面配置数据结构
type AppSetting struct {
	Type int8            `json:"type"`
	Data json.RawMessage `json:"data"`
}
