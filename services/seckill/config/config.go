package config

import (
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
	DFS: app.DFSConfig{
		Endpoint: "http://127.0.0.1:8080",
	},
	AMQP: app.AMQPConfig{
		Addr: "amqp://guest:guest@localhost:5672",
	},
}

type Config struct {
	Redis redis.Config
	DFS   app.DFSConfig
	AMQP  app.AMQPConfig
}
