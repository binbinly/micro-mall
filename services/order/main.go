package main

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/server"

	"order/cmd"
	"order/config"
	"order/event"
	"order/handler"
	"order/logic"
	"pkg/app"
	"pkg/constvar"
	pb "pkg/proto/order"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init broker
	b, err := app.NewRabbitBroker(config.Cfg.AMQP.Addr, constvar.ExchangeOrder)
	if err != nil {
		log.Fatal(err)
	}
	// init delay queue
	initDelay(b)

	// init logic
	a := app.New(
		app.WithName(constvar.ServiceOrder),
		app.WithVersion(version),
		app.WithAuthFunc(handler.Auth),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))
	a.Init(micro.Broker(b))

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	l := logic.New(db, rdb, a.Service().Client())

	// register handler
	if err = pb.RegisterOrderHandler(a.Service().Server(), handler.New(l)); err != nil {
		log.Fatal(err)
	}

	// register subscriber 订阅订单秒杀事件
	if err = micro.RegisterSubscriber(constvar.TopicOrderSeckill, a.Service().Server(), event.NewKill(l)); err != nil {
		log.Fatal(err)
	}

	// register subscriber 监听死信队列，处理订单超时
	if err = micro.RegisterSubscriber(constvar.KeyOrderRelease, a.Service().Server(), event.NewCancel(l),
		server.SubscriberQueue(constvar.QueueReleaseOrder),
	); err != nil {
		log.Fatal(err)
	}

	// run
	a.Run()
}

func initDelay(b broker.Broker) {
	//if _, err := b.Subscribe(constvar.KeyOrderRelease, func(p broker.Event) error {
	//	fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
	//	return nil
	//}, broker.Queue(constvar.QueueReleaseOrder)); err != nil {
	//	log.Fatal(err)
	//}

	s, err := b.Subscribe(constvar.KeyOrderCreate, nil,
		broker.Queue(constvar.QueueDelayOrder),
		rabbitmq.DurableQueue(),
		rabbitmq.QueueArguments(map[string]interface{}{
			"x-dead-letter-exchange":    constvar.ExchangeOrder,   // 设置死信交换器
			"x-dead-letter-routing-key": constvar.KeyOrderRelease, // 设置死信路由键
			"x-message-ttl":             60000,                    // 设置过期时间
		}))
	if err != nil {
		log.Fatal(err)
	}
	// 死信队列不可以订阅消息
	if err = s.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}
