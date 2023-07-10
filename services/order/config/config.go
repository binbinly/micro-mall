package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	MySQL: app.DefOrmConfig("mall_oms"),
	Redis: app.DefRedisConfig,
	DFS:   app.DefDfsConfig,
	AMQP:  app.DefAMQPConfig,
}

type Config struct {
	MySQL orm.Config
	Redis redis.Config
	DFS   app.DFSConfig
	AMQP  app.AMQPConfig
}
