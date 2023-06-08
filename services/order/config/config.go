package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"pkg/app"
)

var Cfg = &Config{
	MySQL: orm.Config{
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "root",
		Database: "mall_oms",
	},
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
	DFS: app.DFSConfig{
		Endpoint: "http://127.0.0.1:8080",
	},
	AMQP: app.AMQPConfig{
		Addr: "amqp://guest:guest@localhost:5672/",
	},
}

type Config struct {
	MySQL orm.Config
	Redis redis.Config
	DFS   app.DFSConfig
	AMQP  app.AMQPConfig
}
