package main

import (
	"context"
	"log"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
	"pkg/constvar"
	"pkg/elasticsearch"
	pb "pkg/proto/product"
	"product/cmd"
	"product/config"
	"product/es"
	"product/handler"
	"product/logic"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceProduct, config.Cfg); err != nil {
		log.Fatal(err)
	}

	// init es
	elastic := es.New(elasticsearch.NewClient(&config.Cfg.Elastic))
	if err := elastic.Init(context.Background()); err != nil {
		log.Fatalf("es init err: %v", err)
	}

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// init broker
	a := app.New(
		app.WithName(constvar.ServiceProduct),
		app.WithVersion(version),
		app.WithMigrate(func() {
			cmd.Migrate()
		}))
	a.Init()

	// register handler
	if err = pb.RegisterProductHandler(a.Service().Server(), handler.New(logic.New(elastic, db, rdb), a.Service().Client())); err != nil {
		log.Fatal(err)
	}

	// run
	a.Run()
}
