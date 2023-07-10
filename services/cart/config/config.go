package config

import (
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	Redis: app.DefRedisConfig,
	DFS:   app.DefDfsConfig,
}

type Config struct {
	Redis redis.Config
	DFS   app.DFSConfig
}
