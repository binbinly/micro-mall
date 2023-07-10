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
	MySQL: app.DefOrmConfig("mall_wms"),
	Redis: app.DefRedisConfig,
	AMQP:  app.DefAMQPConfig,
}
