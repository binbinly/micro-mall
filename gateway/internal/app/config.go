package app

import (
	"context"
	"google.golang.org/grpc"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

type Service struct {
	Name    string
	Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}
