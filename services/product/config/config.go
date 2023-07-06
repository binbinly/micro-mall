package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"pkg/app"
	"pkg/elasticsearch"
)

type Config struct {
	MySQL   orm.Config
	Redis   redis.Config
	DFS     app.DFSConfig
	Elastic elasticsearch.Config
}

var Cfg = &Config{
	MySQL: orm.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root",
		Database: "mall_pms",
	},
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
	DFS: app.DFSConfig{
		Endpoint: "http://127.0.0.1:8080",
	},
	Elastic: elasticsearch.Config{
		Host: "http://127.0.0.1:9200",
	},
}
