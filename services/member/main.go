package main

import (
	"github.com/binbinly/pkg/storage/redis"
	"log"
	"member/cmd"
	"member/config"
	"member/handler"
	"member/logic"
	"pkg/app"
	"pkg/constvar"
	"pkg/mysql"
	pb "pkg/proto/member"

	"go-micro.dev/v4/logger"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceMember, config.Cfg); err != nil {
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
		app.WithName(constvar.ServiceMember),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}),
		app.WithAuthFunc(handler.Auth))
	a.Init()

	// register handler
	if err = pb.RegisterMemberHandler(a.Service().Server(), handler.New(logic.New(db, rdb), a.Service().Client())); err != nil {
		logger.Fatal(err)
	}

	a.Run()
}
