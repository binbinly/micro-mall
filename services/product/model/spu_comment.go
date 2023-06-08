package model

import (
	"pkg/mysql"
)

// SpuCommentModel 商品评价
type SpuCommentModel struct {
	mysql.PriID
	mysql.Spu
	mysql.Sku
	SkuName    string `json:"sku_name" gorm:"column:sku_name;not null;type:varchar(255);comment:商品名"`
	MemberID   int64  `json:"member_id" gorm:"column:member_id;not null;type:int(11) unsigned;comment:会员id"`
	ReplyID    int64  `json:"reply_id" gorm:"column:reply_id;index;not null;type:int unsigned;default:0;comment:回复内容id"`
	OrderID    int64  `json:"order_id" gorm:"column:order_id;not null;type:int;comment:订单id"`
	Star       int8   `json:"star" gorm:"column:star;not null;comment:星级"`
	SkuAttrs   string `json:"sku_attrs" gorm:"column:sku_attrs;not null;type:varchar(255);comment:购买属性组合"`
	LikeCount  int    `json:"like_count" gorm:"column:like_count;not null;type:int(11) unsigned;default:0;comment:点赞数"`
	ReplyCount int    `json:"reply_count" gorm:"column:reply_count;not null;type:int(11) unsigned;default:0;comment:回复数"`
	Resources  string `json:"resources" gorm:"column:resources;not null;type:varchar(1000);default:'';comment:评论图片/视频[json数据；[{type:文件类型,url:资源路径}]]"`
	Content    string `json:"content" gorm:"column:content;not null;type:varchar(5000);comment:内容"`
	mysql.Release
	mysql.CUT
}

// TableName 表名
func (u *SpuCommentModel) TableName() string {
	return "pms_spu_comment"
}
