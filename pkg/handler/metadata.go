package handler

import (
	"context"
	"go-micro.dev/v4/metadata"
	"strconv"
)

// GetUserID 获取用户id
func GetUserID(ctx context.Context) int64 {
	md, _ := metadata.FromContext(ctx)
	if uid, ok := md["User-Id"]; ok {
		id, _ := strconv.ParseInt(uid, 10, 64)
		return id
	}
	return 0
}

// CHeckTokenEmpty 验证是否携带令牌
func CHeckTokenEmpty(ctx context.Context) bool {
	md, _ := metadata.FromContext(ctx)
	if token, ok := md["Token"]; ok {
		if token != "" {
			return true
		}
	}
	return false
}
