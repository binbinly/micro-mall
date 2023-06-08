package model

import (
	"pkg/mysql"
)

// OrderItemModel 订单项信息
type OrderItemModel struct {
	mysql.PriID
	OID
	ONo
	mysql.Sku
	SkuName           string `json:"sku_name" gorm:"column:sku_name;not null;type:varchar(60);comment:商品sku名"`
	SkuImg            string `json:"sku_img" gorm:"column:sku_img;not null;type:varchar(128);comment:商品sku图片"`
	SkuPrice          int    `json:"sku_price" gorm:"column:sku_price;not null;type:int;comment:商品单价/分"`
	SkuAttrs          string `json:"sku_attrs" gorm:"column:sku_attrs;not null;type:varchar(255);comment:商品销售属性组合（JSON）"`
	Num               int    `json:"sku_num" gorm:"column:sku_num;not null;type:int;comment:购买数量"`
	PromotionAmount   int    `json:"promotion_amount" gorm:"column:promotion_amount;not null;type:int;default:0;comment:促销优惠金额（促销价、满减、阶梯价）"`
	IntegrationAmount int    `json:"integration_amount" gorm:"column:integration_amount;not null;type:int;default:0;comment:积分抵扣金额"`
	CouponAmount      int    `json:"coupon_amount" gorm:"column:coupon_amount;not null;type:int;default:0;comment:优惠券折扣金额"`
	RealAmount        int    `json:"real_amount" gorm:"column:real_amount;not null;type:int;comment:优惠后的真实金额"`
	GiveIntegration   int    `json:"give_integration" gorm:"column:give_integration;not null;type:int;default:0;comment:赠送积分"`
	GiveGrowth        int    `json:"give_growth" gorm:"column:give_growth;not null;type:int;default:0;comment:赠送成长值"`
}

// TableName 表名
func (u *OrderItemModel) TableName() string {
	return "oms_order_item"
}
