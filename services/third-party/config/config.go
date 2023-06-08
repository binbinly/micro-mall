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
	Sms: SmsConfig{IsReal: false},
	Eth: EthConfig{
		NetworkID:  15,
		NetworkUrl: "http://127.0.0.1:8545",
	},
}

type Config struct {
	Redis redis.Config
	Sms   SmsConfig
	Eth   EthConfig
}

// SmsConfig 短信配置
type SmsConfig struct {
	IsReal bool
}

// EthConfig 以太坊配置
type EthConfig struct {
	NetworkID  int
	NetworkUrl string
}
