package config

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
)

type Config struct {
	Secret Secret
	MySQL  orm.Config
	Redis  redis.Config
}

var Cfg = &Config{
	Secret: Secret{
		JwtSecret:  "u33A7ss43e89bbNtjtfkARnpi96zKWVv11MHOiHHdl89rLjKsIe15oNPpOuRkslG",
		JwtTimeout: 86400,
	},
	MySQL: orm.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root",
		Database: "center",
	},
	Redis: redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	},
}

type Secret struct {
	JwtSecret  string
	JwtTimeout int64
}
