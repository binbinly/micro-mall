package app

import (
	"time"

	"github.com/spf13/viper"
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

// SetDefaultConf 设置默认配置
func SetDefaultConf(v *viper.Viper) {
	v.SetDefault("app", map[string]any{
		"name":      "gateway",
		"version":   "latest",
		"jwtSecret": "u33A7ss43e89bbNtjtfkARnpi96zKWVv11MHOiHHdl89rLjKsIe15oNPpOuRkslG",
		"debug":     true,
	})
	v.SetDefault("http", map[string]any{
		"addr":         ":9520",
		"readTimeout":  5 * time.Second,
		"writeTimeout": 5 * time.Second,
	})
	v.SetDefault("registry", map[string]any{
		"name": "consul",
		"addr": "127.0.0.1:8500",
	})
	v.SetDefault("trace", map[string]any{
		"endpoint": "http://127.0.0.1:14268/api/traces",
	})
	v.BindEnv("registry.addr", "MALL_REGISTRY_ADDR")
	v.BindEnv("trace.endpoint", "MALL_TRACE_ENDPOINT")
}
