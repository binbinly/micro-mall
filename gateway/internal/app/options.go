package app

import (
	"context"
	"time"

	"gateway/pkg/registry"
	"gateway/proto/cart"
	"gateway/proto/market"
	"gateway/proto/member"
	"gateway/proto/order"
	"gateway/proto/product"
	"gateway/proto/seckill"

	"github.com/binbinly/pkg/transport"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
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
	{
		Name: "mall.cart",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return cart.RegisterCartHandler(ctx, mux, conn)
		},
	},
	{
		Name: "mall.member",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return member.RegisterMemberHandler(ctx, mux, conn)
		},
	},
	{
		Name: "mall.order",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return order.RegisterOrderHandler(ctx, mux, conn)
		},
	},
	{
		Name: "mall.seckill",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return seckill.RegisterSeckillHandler(ctx, mux, conn)
		},
	},
	{
		Name: "mall.product",
		Handler: func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return product.RegisterProductHandler(ctx, mux, conn)
		},
	},
}

type Service struct {
	Name    string
	Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
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
