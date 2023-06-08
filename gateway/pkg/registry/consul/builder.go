package consul

import (
	"google.golang.org/grpc/resolver"

	"gateway/pkg/registry"
)

var _ resolver.Builder = &builder{}

// schemeName for the urls
// All target URLs like 'consul://.../...' will be resolved by this resolver
const schemeName = "consul"

// builder implements resolver.Builder and use for constructing all consul resolvers
type builder struct {
	rs registry.Registry
}

// NewBuilder 实例化 consul resolver
func NewBuilder(rs registry.Registry) resolver.Builder {
	return &builder{rs: rs}
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cr := &consulResolver{
		name:                 target.Endpoint(),
		cc:                   cc,
		disableServiceConfig: opts.DisableServiceConfig,
	}

	cr.wg.Add(1)
	go cr.watcher(b.rs)
	return cr, nil
}

// Scheme returns the scheme supported by this resolver.
// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
func (b *builder) Scheme() string {
	return schemeName
}
