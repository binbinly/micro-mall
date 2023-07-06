package main

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"go-micro.dev/v4"

	"pkg/app"
	"pkg/constvar"
	pb "pkg/proto/warehouse"
	"warehouse/cmd"
	"warehouse/config"
	"warehouse/event"
	"warehouse/handler"
	"warehouse/logic"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceWarehouse, config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init logic
	a := app.New(
		app.WithName(constvar.ServiceWarehouse),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))

	// init broker
	b, err := app.NewRabbitBroker(config.Cfg.AMQP.Addr, constvar.ExchangeOrder)
	if err != nil {
		log.Fatal(err)
	}
	a.Init(micro.Broker(b))

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// new logic
	l := logic.New(db, rdb)

	// register handler
	if err = pb.RegisterWarehouseHandler(a.Service().Server(), handler.New(l)); err != nil {
		log.Fatal(err)
	}

	// register subscriber 订阅订单完成事件，扣减，解锁库存处理
	if err = micro.RegisterSubscriber(constvar.TopicOrderRelease, a.Service().Server(), event.New(l)); err != nil {
		log.Fatal(err)
	}

	// run
	a.Run()
}
