package constvar

const (
	// ExchangeOrder 订单服务rabbitmq交换器名
	ExchangeOrder = "order.event.exchange"
)

const (
	// TopicOrderRelease 订单服务中订单状态完成
	TopicOrderRelease = "mall.order.release"
	// TopicOrderSeckill 秒杀服务秒杀
	TopicOrderSeckill = "mall.order.seckill"
)

const (
	// KeyOrderCreate 订单创建路由键
	KeyOrderCreate = "order.create.order"
	// KeyOrderRelease 订单过期路由键
	KeyOrderRelease = "order.release.order"
)

const (
	// QueueReleaseOrder 订单服务中普通队列
	QueueReleaseOrder = "order.release.queue"
	// QueueDelayOrder     订单服务中死信队列
	QueueDelayOrder = "order.delay.queue"
)
