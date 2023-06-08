package router

import (
	"context"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// filter 过滤器，改变响应消息或设置响应头
func filter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	corsHeader(w)

	//fmt.Printf("resp:%#v\n", resp)

	return nil
}

// corsHeader 跨域响应头设置
func corsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
