package config

import (
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	Redis: app.DefRedisConfig,
	DFS:   app.DefDfsConfig,
	AMQP:  app.DefAMQPConfig,
}

type Config struct {
	Redis redis.Config
	DFS   app.DFSConfig
	AMQP  app.AMQPConfig
}
