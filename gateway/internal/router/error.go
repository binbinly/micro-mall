package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/binbinly/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// handleRoutingError 自定义路由错误
func handleRoutingError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	// 放行所有OPTIONS方法，因为有的模板是要请求两次的
	if r.Method == "OPTIONS" {
		corsHeader(w)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if httpStatus != http.StatusMethodNotAllowed {
		runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
		return
	}

	// Use HTTPStatusError to customize the DefaultHTTPErrorHandler status code
	err := &runtime.HTTPStatusError{
		HTTPStatus: httpStatus,
		Err:        status.Error(codes.Unimplemented, http.StatusText(httpStatus)),
	}

	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}

// handleError 错误处理
func handleError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	code := codes.Code(500)
	msg := "Internal Server Error"

	if s, ok := status.FromError(err); ok {
		logger.Debugf("method: %v, code: %v, msg: %v", r.Method, s.Code(), s.Message())
		if s.Code() == codes.InvalidArgument {
			code = codes.Code(400)
			msg = "请求参数非法"
		} else {
			e, errr := Parse(s.Message())
			if errr != nil {
				code = s.Code()
				msg = s.Message()
			} else {
				code = codes.Code(e.Code)
				msg = e.Detail
			}
		}
	} else {
		logger.Warnf("[gateway] err: %v", err)
	}

	err = &runtime.HTTPStatusError{
		HTTPStatus: http.StatusOK,
		Err:        status.Error(code, msg),
	}
	corsHeader(w)
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}

// Parse go-micro rpc err parse
func Parse(err string) (*Error, error) {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		return nil, errr
	}
	return e, nil
}

// Error go-micro rpc err
type Error struct {
	Id     string `json:"id,omitempty"`
	Code   int32  `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
	Status string `json:"status,omitempty"`
}
