package model

import (
	"pkg/dbs"
)

// MemberStatModel 会员统计信息
type MemberStatModel struct {
	dbs.PriID
	dbs.MID
	TotalAmount         int `json:"total_amount" gorm:"column:total_amount;not null;default:0;comment:累计消费金额/分"`
	CouponAmount        int `json:"coupon_amount" gorm:"column:coupon_amount;not null;default:0;comment:累计优惠金额/分"`
	OrderCount          int `json:"order_count" gorm:"column:order_count;not null;default:0;comment:订单数"`
	CouponCount         int `json:"coupon_count" gorm:"column:coupon_count;not null;default:0;comment:优惠券数"`
	CommentCount        int `json:"comment_count" gorm:"column:comment_count;not null;default:0;comment:评价数"`
	ReturnCount         int `json:"return_count" gorm:"column:return_count;not null;default:0;comment:退货数"`
	LoginCount          int `json:"login_count" gorm:"column:login_count;not null;default:0;comment:登录次数"`
	FollowCount         int `json:"follow_count" gorm:"column:follow_count;not null;default:0;comment:关注数"`
	FensCount           int `json:"fens_count" gorm:"column:fens_count;not null;default:0;comment:粉丝数"`
	CollectProductCount int `json:"collect_product_count" gorm:"column:collect_product_count;not null;default:0;comment:收藏商品数"`
	CollectSubjectCount int `json:"collect_subject_count" gorm:"column:collect_subject_count;not null;default:0;comment:收藏专题活动数"`
	CollectCommentCount int `json:"collect_comment_count" gorm:"column:collect_comment_count;not null;default:0;comment:收藏评论数"`
	InviteCount         int `json:"invite_count" gorm:"column:invite_count;not null;default:0;comment:邀请好友数"`
}

// TableName 表名
func (u *MemberStatModel) TableName() string {
	return "ums_member_stat"
}
