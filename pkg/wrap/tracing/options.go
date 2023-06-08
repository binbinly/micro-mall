package tracing

import (
	"github.com/binbinly/pkg/util"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	localIP = util.LocalIP()
)

type options struct {
	ServiceName    string
	TracerProvider trace.TracerProvider
	Propagators    propagation.TextMapPropagator
}

// Option specifies instrumentation configuration options.
type Option func(*options)

// WithServiceName set service name
func WithServiceName(name string) Option {
	return func(o *options) {
		o.ServiceName = name
	}
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.Propagators = propagators
	}
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(o *options) {
		o.TracerProvider = provider
	}
}
