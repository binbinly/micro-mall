package app

import (
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4/broker"
)

// NewRabbitBroker 创建rabbitmq broker
func NewRabbitBroker(addr, exchange string) (broker.Broker, error) {
	b := rabbitmq.NewBroker(broker.Addrs(addr),
		rabbitmq.DurableExchange(), // 消息在Exchange中时会进行持久化处理
		rabbitmq.ExchangeName(exchange))
	if err := b.Init(); err != nil {
		return nil, err
	}
	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}
