package recovery

import (
	"context"

	"go-micro.dev/v4/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type HandlerFunc func(p any) (err error)

// HandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type HandlerFuncContext func(ctx context.Context, p any) (err error)

// NewHandlerWrapper 错误捕获
func NewHandlerWrapper(opts ...Option) server.HandlerWrapper {
	o := evaluateOptions(opts)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp any) (err error) {
			panicked := true

			defer func() {
				if r := recover(); r != nil || panicked {
					err = recoverFrom(ctx, r, o.recoveryHandlerFunc)
				}
			}()
			err = fn(ctx, req, rsp)
			panicked = false
			return err
		}
	}
}

func recoverFrom(ctx context.Context, p any, r HandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
