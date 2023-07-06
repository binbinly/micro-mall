// Copyright 2017 David Ackroyd. All Rights Reserved.
// See LICENSE for licensing terms.

package recovery

import "context"

var (
	defaultOptions = &options{
		recoveryHandlerFunc: nil,
	}
)

type options struct {
	recoveryHandlerFunc HandlerFuncContext
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithHandler customizes the function for recovering from a panic.
func WithHandler(f HandlerFunc) Option {
	return func(o *options) {
		o.recoveryHandlerFunc = HandlerFuncContext(func(ctx context.Context, p any) error {
			return f(p)
		})
	}
}

// WithHandlerContext customizes the function for recovering from a panic.
func WithHandlerContext(f HandlerFuncContext) Option {
	return func(o *options) {
		o.recoveryHandlerFunc = f
	}
}
