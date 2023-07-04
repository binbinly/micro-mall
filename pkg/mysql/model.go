package mysql

import (
	"errors"

	"gorm.io/gorm"
)

const (
	// ReleaseYes 已发布
	ReleaseYes = 1
	// DefaultOrder 默认排序
	DefaultOrder = "id DESC"
	// DefaultOrderSort 排序字段排序
	DefaultOrderSort = "sort DESC"
)

const (
	// StatusInit 状态-初始化
	StatusInit = iota
	// StatusSuccess 状态-成功
	StatusSuccess
	// StatusError 状态-失败
	StatusError
)

// ErrRecordNotModified 记录未修改
var ErrRecordNotModified = errors.New("record not modified")

// CUT 公共时间字段
type CUT struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// CUBy 公共操作者字段
type CUBy struct {
	CreateBy int `gorm:"column:create_by;type:int(11) unsigned;not null;default:0;comment:创建者" json:"create_by"`
	UpdateBy int `gorm:"column:update_by;type:int(11) unsigned;not null;default:0;comment:更新者" json:"update_by"`
}

// PriID 主键
type PriID struct {
	ID int64 `gorm:"primaryKey;autoIncrement;type:int(11) unsigned auto_increment;column:id;comment:ID" json:"id"`
}

// CT 创建时间
type CT struct {
	CreatedAt int64 `gorm:"column:created_at;type:int(11) unsigned;not null;autoCreateTime;comment:创建时间" json:"created_at"`
}

// UT 更新时间
type UT struct {
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11) unsigned;not null;autoUpdateTime;comment:更新时间" json:"updated_at"`
}

// DT 删除时间
type DT struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
}

// MID 用户ID
type MID struct {
	MemberID int64 `gorm:"column:member_id;not null;type:int(11) unsigned;index;comment:用户id" json:"member_id"`
}

// Spu 商品spu
type Spu struct {
	SpuID int64 `json:"spu_id" gorm:"column:spu_id;not null;index;type:int unsigned;comment:spu_id"`
}

// Sku 商品sku
type Sku struct {
	SkuID int64 `json:"sku_id" gorm:"column:sku_id;not null;index;type:int unsigned;comment:sku_id"`
}

// Cat 商品分类
type Cat struct {
	CatID int64 `json:"cat_id" gorm:"column:cat_id;not null;type:int unsigned;comment:产品分类"`
}

// OrderBy 排序字段
type OrderBy struct {
	Sort int32 `json:"sort" gorm:"column:sort;not null;type:int unsigned;default:50;comment:排序"`
}

// Release 是否发布
type Release struct {
	IsRelease int8 `json:"is_release" gorm:"column:is_release;not null;type:tinyint unsigned;default:0;comment:是否发布上线"`
}

// OffsetPage 分页查询
func OffsetPage(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

// WhereRelease 发布
func WhereRelease(db *gorm.DB) *gorm.DB {
	return db.Where("is_release=?", ReleaseYes)
}
