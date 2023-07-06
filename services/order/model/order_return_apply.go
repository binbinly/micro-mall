package model

import (
	"pkg/dbs"
)

// OrderReturnApplyModel 订单退货申请
type OrderReturnApplyModel struct {
	dbs.PriID
	OID
	ONo
	dbs.Sku
	Username      string `json:"username" gorm:"column:username;not null;type:varchar(60);comment:用户名"`
	Amount        int    `json:"amount" gorm:"column:amount;not null;type:int;comment:退款金额"`
	ReturnName    string `json:"return_name" gorm:"column:return_name;not null;type:varchar(100);comment:退款人姓名"`
	ReturnPhone   string `json:"return_phone" gorm:"column:return_phone;not null;type:char(11);comment:退款人手机号"`
	Status        int8   `json:"status" gorm:"column:status;not null;default:0;comment:订单状态"`
	SkuName       string `json:"sku_name" gorm:"column:sku_name;not null;type:varchar(60);comment:商品sku名"`
	SkuImg        string `json:"sku_img" gorm:"column:sku_img;not null;type:varchar(128);comment:商品sku图片"`
	SkuPrice      int    `json:"sku_price" gorm:"column:sku_price;not null;type:int;comment:商品单价/分"`
	SkuAttrs      string `json:"sku_attrs" gorm:"column:sku_attrs;not null;type:varchar(255);comment:商品销售属性组合（JSON）"`
	PayAmount     int    `json:"pay_amount" gorm:"column:pay_amount;not null;type:int;comment:支付金额"`
	Num           int    `json:"sku_num" gorm:"column:sku_num;not null;type:int;comment:退货数量"`
	SpuBrand      string `json:"spu_brand" gorm:"column:spu_brand;not null;type:varchar(60);comment:品牌"`
	Reason        string `json:"reason" gorm:"column:reason;not null;type:varchar(255);comment:原因"`
	Desc          string `json:"desc" gorm:"column:desc;not null;type:varchar(255);comment:描述"`
	DescPics      string `json:"desc_pics" gorm:"column:desc_pics;not null;type:varchar(2000);default:'';comment:凭证图片多个逗号连接"`
	HandleAt      int64  `json:"handle_at" gorm:"column:handle_at;not null;type:int;default:0;comment:处理时间"`
	HandleNote    string `json:"handle_note" gorm:"column:handle_note;not null;type:varchar(255);default:'';comment:处理备注"`
	HandleBy      string `json:"handle_by" gorm:"column:handle_by;not null;type:varchar(50);default:'';comment:处理人"`
	ReceiveName   string `json:"receive_name" gorm:"column:receive_name;not null;type:varchar(64);comment:收货人姓名"`
	ReceivePhone  string `json:"receive_phone" gorm:"column:receive_phone;not null;type:char(11);comment:收货人手机"`
	ReceiveNote   string `json:"receive_note" gorm:"column:receive_note;not null;type:varchar(255);default:'';comment:收货备注"`
	ReceiveAt     int64  `json:"receive_at" gorm:"column:receive_at;not null;type:int;default:0;comment:收货时间"`
	AddressDetail string `json:"address_detail" gorm:"column:address_detail;not null;type:varchar(255);comment:收货地址"`
	dbs.CUT
}

// TableName 表名
func (u *OrderReturnApplyModel) TableName() string {
	return "oms_order_return_apply"
}
