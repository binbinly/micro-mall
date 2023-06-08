package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
)

type Config struct {
	MySQL orm.Config
	Redis redis.Config
}

var Cfg = &Config{
	MySQL: orm.Config{
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "root",
		Database: "mall_ums",
	},
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
}
