package main

import (
	"log"
	"warehouse/event"

	"github.com/binbinly/pkg/storage/redis"
	"go-micro.dev/v4"
	"pkg/app"
	"pkg/constvar"
	"pkg/mysql"
	"warehouse/cmd"
	"warehouse/config"
	"warehouse/logic"

	pb "pkg/proto/warehouse"
	"warehouse/handler"

	"go-micro.dev/v4/logger"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceWarehouse, config.Cfg); err != nil {
		logger.Fatal(err)
	}

	// init mysql
	db := mysql.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// init broker
	b, err := app.NewRabbitBroker(config.Cfg.AMQP.Addr, constvar.ExchangeOrder)
	if err != nil {
		log.Fatal(err)
	}

	// init logic
	a := app.New(
		app.WithName(constvar.ServiceWarehouse),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))
	a.Init(micro.Broker(b))

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
