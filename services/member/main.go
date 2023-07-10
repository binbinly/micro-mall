package main

import (
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"member/cmd"
	"member/config"
	"member/handler"
	"member/logic"
	"pkg/app"
	"pkg/constvar"
	pb "pkg/proto/member"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init app
	a := app.New(
		app.WithName(constvar.ServiceMember),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}),
		app.WithAuthFunc(handler.Auth))
	a.Init()

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// register handler
	if err = pb.RegisterMemberHandler(a.Service().Server(), handler.New(logic.New(db, rdb), a.Service().Client())); err != nil {
		log.Fatal(err)
	}

	a.Run()
}
