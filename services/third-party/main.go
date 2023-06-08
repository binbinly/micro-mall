package main

import (
	"github.com/binbinly/pkg/storage/redis"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"pkg/app"
	"pkg/constvar"
	"third-party/config"
	"third-party/logic"

	pb "pkg/proto/core"
	"third-party/handler"

	"go-micro.dev/v4/logger"
)

var (
	version = "latest"
)

func main() {
	// load config
	if err := app.LoadEnv(constvar.ServiceThird, config.Cfg); err != nil {
		log.Fatal(err)
	}

	// connect to eth
	c, err := ethclient.Dial(config.Cfg.Eth.NetworkUrl)
	if err != nil {
		log.Fatalf("contract connect err:%v", err)
	}

	// init redis
	rdb, err := redis.NewClient(&config.Cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// init app
	a := app.New(
		app.WithName(constvar.ServiceThird),
		app.WithVersion(version))
	a.Init()

	// register handler
	if err = pb.RegisterThirdPartyHandler(a.Service().Server(), handler.New(logic.New(rdb, c))); err != nil {
		logger.Fatal(err)
	}

	// run
	a.Run()
}
