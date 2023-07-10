package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	Secret: Secret{
		JwtSecret:  "u33A7ss43e89bbNtjtfkARnpi96zKWVv11MHOiHHdl89rLjKsIe15oNPpOuRkslG",
		JwtTimeout: 86400,
	},
	MySQL: app.DefOrmConfig("center"),
	Redis: app.DefRedisConfig,
}

type Config struct {
	Secret Secret
	MySQL  orm.Config
	Redis  redis.Config
}

type Secret struct {
	JwtSecret  string
	JwtTimeout int64
}
