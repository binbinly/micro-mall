package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
	"log"
)

const (
	queue        = "order.release.queue"
	delayQueue   = "order.delay.queue"
	exchangeName = "order.event.exchange"
)

const (
	//KeyOrderCreate 订单创建路由键
	KeyOrderCreate = "order.create.order"
	//KeyOrderRelease 订单过期路由键
	KeyOrderRelease = "order.release.order"
)

// Example of a shared subscription which receives a subset of messages
func delaySub(b broker.Broker) {
	_, err := b.Subscribe(KeyOrderRelease, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue(queue))
	if err != nil {
		fmt.Println(err)
	}

	s, err := b.Subscribe(KeyOrderCreate, nil, broker.Queue(delayQueue),
		rabbitmq.DurableQueue(),
		rabbitmq.QueueArguments(map[string]any{
			"x-dead-letter-exchange":    exchangeName,    // 设置死信交换器
			"x-dead-letter-routing-key": KeyOrderRelease, // 设置死信路由键
			"x-message-ttl":             5000,            //设置过期时间
		}))
	if err != nil {
		fmt.Println(err)
	}
	// 死信队列不可以订阅消息
	s.Unsubscribe()
}

func main() {
	cmd.Init()

	b := rabbitmq.NewBroker(broker.Addrs("amqp://guest:guest@localhost:5672"),
		rabbitmq.ExchangeName(exchangeName), // 交换机的名称
		rabbitmq.DurableExchange(),          // 消息在Exchange中时会进行持久化处理
	)

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	delaySub(b)

	fmt.Println("start success")
	select {}
}
