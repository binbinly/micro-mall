package recovery

import (
	"context"

	"go-micro.dev/v4/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerFunc func(p interface{}) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// NewHandlerWrapper 错误捕获
func NewHandlerWrapper(opts ...Option) server.HandlerWrapper {
	o := evaluateOptions(opts)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
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

func recoverFrom(ctx context.Context, p interface{}, r RecoveryHandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
