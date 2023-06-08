package wrap

import (
	"context"

	"go-micro.dev/v4/errors"
	"go-micro.dev/v4/server"
)

type validator interface {
	Validate() error
}

// Validator 参数验证
func Validator() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if v, ok := req.Body().(validator); ok {
				if err := v.Validate(); err != nil {
					return errors.BadRequest(req.Service(), "%v", err)
				}
			}
			return fn(ctx, req, rsp)
		}
	}
}
