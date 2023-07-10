package app

import (
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"github.com/pkg/errors"
	"go-micro.dev/v4/config/source/env"
	"time"

	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"go-micro.dev/v4/logger"
)

var DefDfsConfig = DFSConfig{
	Endpoint: "http://127.0.0.1:8080",
}

var DefAMQPConfig = AMQPConfig{
	Addr: "amqp://guest:guest@localhost:5672",
}

var DefRedisConfig = redis.Config{
	Addr:         "127.0.0.1:6379",
	Password:     "",
	DB:           0,
	MinIdleConn:  1,
	DialTimeout:  5 * time.Second,
	ReadTimeout:  500 * time.Millisecond,
	WriteTimeout: 500 * time.Millisecond,
	PoolSize:     10,
	PoolTimeout:  5 * time.Minute,
}

func DefOrmConfig(database string) orm.Config {
	return orm.Config{
		Host:            "127.0.0.1",
		Port:            3306,
		User:            "root",
		Password:        "root",
		Database:        database,
		Debug:           true,
		MaxIdleConn:     1,
		MaxOpenConn:     10,
		ConnMaxLifeTime: 100 * time.Second,
	}
}

// DFSConfig 图片资源配置
type DFSConfig struct {
	Endpoint string `json:"endpoint"`
}

// AMQPConfig amqp配置
type AMQPConfig struct {
	Addr string `json:"addr"`
}

// LoadEnv 加载系统 env 配置
func LoadEnv(val any) error {
	conf, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "conf.New")
	}

	return load(conf, "", val)
}

// LoadFile 加载文件配置
func LoadFile(filepath string, path string, val any) error {
	// new config
	conf, err := config.NewConfig(
		config.WithSource(file.NewSource(
			file.WithPath(filepath),
			source.WithEncoder(yaml.NewEncoder()),
		)),
	)
	if err != nil {
		return errors.Wrap(err, "conf.New")
	}

	return load(conf, path, val)
}

// LoadConsul 加载consul配置
func LoadConsul(address string, path string, val any) error {
	// new config
	conf, err := config.NewConfig(
		config.WithSource(consul.NewSource(
			//设置配置中心地址
			consul.WithAddress(address),
			//设置前缀，不设置默认为 /micro/config
			consul.WithPrefix("/mall/config"),
			//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
			consul.StripPrefix(true),
			source.WithEncoder(yaml.NewEncoder()),
		)),
	)
	if err != nil {
		return errors.Wrap(err, "conf.New")
	}

	return load(conf, path, val)
}

func load(conf config.Config, path string, val any) error {
	// load the config from a file source
	if err := conf.Load(); err != nil {
		return errors.Wrap(err, "conf.Load")
	}

	if path != "" {
		go watch(conf, path, val)
		conf.Get(path)
	}

	if err := conf.Scan(val); err != nil {
		return errors.Wrap(err, "conf.Scan")
	}
	logger.Infof("[config] load config: %v", val)

	return nil
}

// watch 配置变化监听
func watch(c config.Config, path string, val any) {
	for {
		w, err := c.Watch(path)
		if err != nil {
			logger.Warnf("[config] watch err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}

		v, err := w.Next()
		if err != nil {
			logger.Warnf("[config] next err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}

		if err = v.Scan(val); err != nil {
			logger.Warnf("[config] scan err: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}

		logger.Info("[config] watch update config: %v", val)
	}
}
