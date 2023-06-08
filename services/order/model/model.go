package model

// OID 订单id
type OID struct {
	OrderID int64 `json:"order_id" gorm:"column:order_id;index;not null;type:int;comment:订单id"`
}

// ONo 订单号
type ONo struct {
	OrderNo string `json:"order_no" gorm:"column:order_no;index;not null;type:char(30);comment:订单号"`
}
