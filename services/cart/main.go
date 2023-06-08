package main

import (
	"cart/config"
	"cart/handler"
	"cart/logic"
	"github.com/binbinly/pkg/storage/redis"
	"log"
	"pkg/app"
	"pkg/constvar"
	pb "pkg/proto/cart"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceCart, config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// init app
	a := app.New(
		app.WithName(constvar.ServiceCart),
		app.WithVersion(version))
	a.Init()

	// register handler
	if err = pb.RegisterCartHandler(a.Service().Server(), handler.New(logic.New(rdb), a.Service().Client())); err != nil {
		log.Fatal(err)
	}

	a.Run()
}
