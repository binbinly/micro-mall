package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

var (
	defTimeout = time.Second * 5
)

// TimeoutClientInterceptor returns a new unary client interceptor that sets a timeout on the request context.
func TimeoutClientInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	if timeout == time.Duration(0) {
		timeout = defTimeout
	}
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		timedCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return invoker(timedCtx, method, req, reply, cc, opts...)
	}
}