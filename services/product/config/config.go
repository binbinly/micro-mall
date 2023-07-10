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
	MySQL: app.DefOrmConfig("mall_pms"),
	Redis: app.DefRedisConfig,
	DFS:   app.DefDfsConfig,
	Elastic: elasticsearch.Config{
		Host: "http://127.0.0.1:9200",
	},
}
