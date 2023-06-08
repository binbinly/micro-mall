package router

import (
	"encoding/json"
	"github.com/binbinly/pkg/util"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/protobuf/encoding/protojson"

	"net/http"
)

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// NewServeMux 实例化网关路由
// 这里直接跨域设置比较繁琐，请求失败，成功，请求方式option 缺一不可，dev下使用，其他环境nginx反向代理实现跨域
func NewServeMux() *runtime.ServeMux {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customMatcher),
		runtime.WithRoutingErrorHandler(handleRoutingError),
		runtime.WithErrorHandler(handleError),
		runtime.WithForwardResponseOption(filter),
		runtime.WithUnescapingMode(runtime.UnescapingModeAllExceptReserved),
		runtime.WithMetadata(addMetadata),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &CustomMarshaler{
				m: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true, // 使用 proto 中字段，而不是驼峰
						EmitUnpopulated: true, // 填充空默认字段
					},
					UnmarshalOptions: protojson.UnmarshalOptions{},
				}},
		}),
		//runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		//	MarshalOptions: protojson.MarshalOptions{
		//		UseProtoNames:   true, // 使用 proto 中字段，而不是驼峰
		//		EmitUnpopulated: true, //填充空默认字段
		//	},
		//}),
	)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	if err := mux.HandlePath("GET", "/metrics", metrics(promhttp.Handler())); err != nil {
		log.Fatalf("[Mux] handle metrics err: %v", err)
	}
	// health 健康检查路由
	if err := mux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&healthCheckResponse{Status: "UP", Hostname: util.Hostname()})
		w.Write(data)
	}); err != nil {
		log.Fatalf("[Mux] handle health err: %v", err)
	}
	// swagger 文档
	if err := mux.HandlePath("GET", "/swagger/*", HandlerSwagger()); err != nil {
		log.Fatalf("[Mux] handle swagger err: %v", err)
	}
	return mux
}

// metrics is a helper function for wrapping http.Handler
func metrics(h http.Handler) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		h.ServeHTTP(w, r)
	}
}
