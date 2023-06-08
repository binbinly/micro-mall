package config

import (
	"github.com/binbinly/pkg/storage/redis"
)

var Cfg = &Config{
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
	DFS: DFSConfig{
		Endpoint: "http://127.0.0.1:8080",
	},
}

type Config struct {
	Redis redis.Config
	DFS   DFSConfig
}

// DFSConfig 图片资源配置
type DFSConfig struct {
	Endpoint string
}
