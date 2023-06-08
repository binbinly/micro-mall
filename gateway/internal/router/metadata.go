package router

import (
	"context"
	"gateway/internal/app"
	"github.com/binbinly/pkg/auth"
	"net/http"
	"strconv"

	"github.com/binbinly/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

// customMatcher 从 HTTP 请求头到 gRPC 客户端元数据的映射
func customMatcher(key string) (string, bool) {
	switch key {
	case "X-Forwarded-For":
		return key, true
	case "X-Real-Ip":
		return key, true
	case "Proxy-Forwarded-For":
		return key, true
	case "Token":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

// addMetadata 添加 gRPC 元数据
func addMetadata(ctx context.Context, r *http.Request) metadata.MD {
	md := make(map[string]string)
	if method, ok := runtime.RPCMethod(ctx); ok {
		md["method"] = method
	}
	if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
		md["pattern"] = pattern
	}
	if token := r.Header.Get("Token"); token != "" { // 解析令牌
		payload, err := auth.Parse(token, app.Conf.App.JwtSecret)
		if err != nil {
			md["User-Id"] = ""
			logger.Infof("JWT Parse err: %v", err)
			return metadata.New(md)
		}
		logger.Infof("Context UserID: %v", payload.UserID)
		md["User-Id"] = strconv.Itoa(payload.UserID)
	}
	return metadata.New(md)
}
