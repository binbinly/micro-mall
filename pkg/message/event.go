package message

// Event 异步消息结构定义

// OrderMessage 订单完成消息
type OrderMessage struct {
	OrderId int64
	Finish  bool
}

// SeckillMessage 秒杀服务秒杀消息结构
type SeckillMessage struct {
	MemberID  int64  // 会员id
	SkuID     int64  // 商品id
	AddressID int64  // 收货地址id
	Num       int64  // 秒杀数量
	Price     int    // 秒杀价格
	OrderNo   string // 订单号
}

// OrderCancelMessage 订单取消消息
type OrderCancelMessage struct {
	OrderID  int64 //订单id
	MemberID int64 //会员id
}
