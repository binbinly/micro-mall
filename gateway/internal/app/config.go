package app

import (
	"time"
)

// Conf 全局配置
var Conf = &Config{}

// Config global config
type Config struct {
	App      AppConfig
	HTTP     ServerConfig
	Trace    TraceConfig
	Registry RegistryConfig
}

// AppConfig app config
type AppConfig struct {
	Name      string
	Version   string
	JwtSecret string
	Debug     bool
}

// RegistryConfig 注册中心
type RegistryConfig struct {
	Addr string
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// TraceConfig 链路追踪配置
type TraceConfig struct {
	Endpoint string
}
