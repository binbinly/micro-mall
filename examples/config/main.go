package main

import (
	"log"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

// export PORT=9000
// export DATABASE_SERVER_HOST=localhost
type Config struct {
	Port     int
	Database Database
}

type Database struct {
	Server Server
}

type Server struct {
	Host string
}

var cfg = &Config{}

func main() {
	err := LoadEnv(cfg)
	if err != nil {
		panic(err)
	}
	log.Printf("cfg: %+v", cfg)
}

// LoadEnv 加载系统 env 配置
func LoadEnv(val any) error {
	conf, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		log.Panic(err)
	}

	return load(conf, val)
}

func load(conf config.Config, val any) error {
	// load the config from a file source
	if err := conf.Load(); err != nil {
		log.Panic(err)
	}

	if err := conf.Scan(val); err != nil {
		log.Panic(err)
	}

	return nil
}
