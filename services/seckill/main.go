package main

import (
	"log"
	"pkg/app"
	"pkg/constvar"
	pb "pkg/proto/seckill"
	"seckill/config"
	"seckill/handler"
	"seckill/logic"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/binbinly/pkg/storage/redis"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceSeckill, config.Cfg); err != nil {
		logger.Fatal(err)
	}

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
		app.WithName(constvar.ServiceSeckill),
		app.WithVersion(version),
		app.WithAuthFunc(handler.Auth))
	a.Init(micro.Broker(b))

	// register handler
	if err = pb.RegisterSeckillHandler(a.Service().Server(), handler.New(logic.New(rdb), a.Service().Client())); err != nil {
		logger.Fatal(err)
	}

	// run
	a.Run()
}
