package main

import (
	"log"
	"market/cmd"
	"market/config"
	"market/handler"
	"market/logic"
	"pkg/app"
	"pkg/constvar"
	"pkg/mysql"
	pb "pkg/proto/market"

	"github.com/binbinly/pkg/storage/redis"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceMarket, config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init mysql
	db := mysql.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// init app
	a := app.New(
		app.WithName(constvar.ServiceMarket),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}),
		app.WithAuthFunc(handler.Auth))
	a.Init()

	// register handler
	if err = pb.RegisterMarketHandler(a.Service().Server(), handler.New(logic.New(db, rdb))); err != nil {
		log.Fatal(err)
	}

	a.Run()
}
