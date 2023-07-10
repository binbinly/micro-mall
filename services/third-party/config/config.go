package config

import (
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
)

var Cfg = &Config{
	Redis: app.DefRedisConfig,
	Sms:   SmsConfig{IsReal: false},
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
