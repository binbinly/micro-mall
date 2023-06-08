package model

import (
	"pkg/mysql"
)

const (
	// OrderStatusInit 初始化,待付款
	OrderStatusInit = 1 + iota
	// OrderStatusDelivered 已支付，待发货, 可退款
	OrderStatusDelivered
	// OrderStatusShipped 已发货
	OrderStatusShipped
	// OrderStatusReceived 已收货，可退款
	OrderStatusReceived
	// OrderStatusFinish 已完成，可评价
	OrderStatusFinish
	// OrderStatusPendingRefund 待退款
	OrderStatusPendingRefund
	// OrderStatusRefund 已退款
	OrderStatusRefund
)

const (
	// OrderSourceTypeApp 订单来源app
	OrderSourceTypeApp = 1
)

// OrderModel 订单
type OrderModel struct {
	mysql.PriID
	mysql.MID
	OrderNo           string `json:"order_no" gorm:"column:order_no;uniqueIndex;not null;type:char(30);comment:订单号"`
	CouponID          int64  `json:"coupon_id" gorm:"column:coupon_id;not null;type:int;default:0;comment:优惠券id"`
	Username          string `json:"username" gorm:"column:username;not null;type:varchar(60);comment:用户名"`
	TotalAmount       int    `json:"total_amount" gorm:"column:total_amount;not null;type:int;comment:订单总额/分"`
	FreightAmount     int    `json:"freight_amount" gorm:"column:freight_amount;not null;default:0;type:int;comment:运费/分"`
	PromotionAmount   int    `json:"promotion_amount" gorm:"column:promotion_amount;not null;type:int;default:0;comment:促销优惠金额（促销价、满减、阶梯价）"`
	IntegrationAmount int    `json:"integration_amount" gorm:"column:integration_amount;not null;type:int;default:0;comment:积分抵扣金额"`
	CouponAmount      int    `json:"coupon_amount" gorm:"column:coupon_amount;not null;type:int;default:0;comment:优惠券折扣金额"`
	DiscountAmount    int    `json:"discount_amount" gorm:"column:discount_amount;not null;type:int;default:0;comment:后台调整订单使用的折扣金额"`
	Amount            int    `json:"amount" gorm:"column:amount;not null;type:int;comment:优惠后的真实金额"`
	PayAmount         int    `json:"pay_amount" gorm:"column:pay_amount;not null;type:int;default:0;comment:支付金额/分"`
	PayType           int8   `json:"pay_type" gorm:"column:pay_type;not null;default:0;comment:支付方式"`
	SourceType        int8   `json:"source_type" gorm:"column:source_type;not null;default:0;comment:订单来源[0->PC订单；1->app订单]"`
	Status            int8   `json:"status" gorm:"column:status;not null;default:0;comment:订单状态"`
	DeliveryCompany   string `json:"delivery_company" gorm:"column:delivery_company;not null;type:varchar(64);default:'';comment:物流公司(配送方式)"`
	DeliveryNo        string `json:"delivery_no" gorm:"column:delivery_no;not null;type:varchar(20);default:'';comment:物流单号"`
	AutoConfirmDay    int    `json:"auto_confirm_day" gorm:"column:auto_confirm_day;not null;type:int;default:15;comment:自动确认天数"`
	Integration       int    `json:"integration" gorm:"column:integration;not null;type:int;default:0;comment:可以获得的积分"`
	Growth            int    `json:"growth" gorm:"column:growth;not null;type:int;default:0;comment:可以获得的成长值"`
	AddressName       string `json:"address_name" gorm:"column:address_name;not null;type:varchar(64);comment:收货人姓名"`
	AddressPhone      string `json:"address_phone" gorm:"column:address_phone;not null;type:char(11);comment:收货人手机"`
	AddressProvince   string `json:"address_province" gorm:"column:address_province;not null;type:varchar(30);comment:省"`
	AddressCity       string `json:"address_city" gorm:"column:address_city;not null;type:varchar(30);comment:市"`
	AddressCounty     string `json:"address_county" gorm:"column:address_county;not null;type:varchar(30);comment:县/区"`
	AddressDetail     string `json:"address_detail" gorm:"column:address_detail;not null;type:varchar(255);comment:详情"`
	Note              string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	IsConfirm         int8   `json:"is_confirm" gorm:"column:is_confirm;not null;default:0;comment:是否确认收货"`
	UseIntegration    int    `json:"use_integration" gorm:"column:use_integration;not null;type:int;default:0;comment:下单时使用的积分"`
	TradeNo           string `json:"trade_no" gorm:"column:trade_no;not null;type:varchar(100);comment:交易号"`
	TransHash         string `json:"trans_hash" gorm:"column:trans_hash;not null;type:varchar(100);comment:交易hash"`
	PayAt             int64  `json:"pay_at" gorm:"column:pay_at;not null;type:int;default:0;comment:支付时间"`
	DeliveryAt        int64  `json:"delivery_at" gorm:"column:delivery_at;not null;type:int;default:0;comment:发货时间"`
	ReceiveAt         int64  `json:"receive_at" gorm:"column:receive_at;not null;type:int;default:0;comment:确认收货时间"`
	CommentAt         int64  `json:"comment_at" gorm:"column:comment_at;not null;type:int;default:0;comment:评价时间"`
	mysql.CUT
	mysql.DT
	Items []*OrderItemModel `json:"items" gorm:"foreignkey:order_id;references:id"`
}

// TableName 表名
func (u *OrderModel) TableName() string {
	return "oms_order"
}
