package app

import (
	"context"

	"gateway/pkg/transport"
	"google.golang.org/grpc"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"gateway/pkg/registry"
	market "gateway/proto/market"
)

var (
	registerTTL      = time.Minute
	registerInterval = time.Second * 30
)

// services 网关代理的所有服务
var services = []Service{
	{
		Name: "mall.market",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return market.RegisterMarketHandler(ctx, mux, conn)
		},
	},
}

// Option is func for application
type Option func(o *options)

// options is an application options
type options struct {
	id               string
	name             string
	version          string
	metadata         map[string]string
	registry         registry.Registry
	registerTTL      time.Duration
	registerInterval time.Duration
	services         []Service
	server           transport.Server
	mux              *runtime.ServeMux
}

// WithID with app id
func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

// WithName .
func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// WithVersion with a version
func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// WithServer with a server
func WithServer(srv transport.Server) Option {
	return func(o *options) {
		o.server = srv
	}
}

// WithMux with mux
func WithMux(mux *runtime.ServeMux) Option {
	return func(o *options) {
		o.mux = mux
	}
}

// WithRegistry with service registry.
func WithRegistry(r registry.Registry) Option {
	return func(o *options) {
		o.registry = r
	}
}

// WithRegistryTTL with service registryTTL.
func WithRegistryTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.registerTTL = ttl
	}
}

// WithRegistryInterval with service registryInterval.
func WithRegistryInterval(interval time.Duration) Option {
	return func(o *options) {
		o.registerInterval = interval
	}
}
