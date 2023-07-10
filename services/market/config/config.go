package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	MySQL: app.DefOrmConfig("mall_sms"),
	Redis: app.DefRedisConfig,
	DFS:   app.DefDfsConfig,
}

type Config struct {
	MySQL orm.Config
	Redis redis.Config
	DFS   app.DFSConfig
}
