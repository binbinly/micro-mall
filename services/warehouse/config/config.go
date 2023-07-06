package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

type Config struct {
	MySQL orm.Config
	Redis redis.Config
	AMQP  app.AMQPConfig
}

var Cfg = &Config{
	MySQL: orm.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root",
		Database: "mall_wms",
	},
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
	AMQP: app.AMQPConfig{
		Addr: "amqp://guest:guest@localhost:5672",
	},
}
