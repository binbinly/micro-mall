package main

import (
	"center/cmd"
	"center/config"
	"center/handler"
	"center/logic"
	"log"
	"pkg/app"
	"pkg/constvar"
	"pkg/mysql"
	pb "pkg/proto/core"

	"github.com/binbinly/pkg/storage/redis"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceCenter, config.Cfg); err != nil {
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
		app.WithName(constvar.ServiceCenter),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))
	a.Init()

	// register handler
	if err = pb.RegisterCenterHandler(a.Service().Server(), handler.New(logic.New(db, rdb))); err != nil {
		log.Fatal(err)
	}

	a.Run()
}
