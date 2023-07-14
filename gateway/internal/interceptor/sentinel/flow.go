package sentinel

import (
	"context"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FlowClientInterceptor 流量控制
func FlowClientInterceptor(opts ...Flow) grpc.UnaryClientInterceptor {
	rule := defaultFlowRule()
	for _, opt := range opts {
		opt(rule)
	}
	if _, err := flow.LoadRules([]*flow.Rule{rule}); err != nil {
		log.Fatalf("[Sentinel] load rules error: %+v", err)
	}
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		e, b := sentinel.Entry(rule.ResourceName(), sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			return status.Error(codes.PermissionDenied, "Server is busy")
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		e.Exit()
		return err
	}
}

// defaultFlowRule 默认流控规则
func defaultFlowRule() *flow.Rule {
	return &flow.Rule{
		Resource:               "flow",
		TokenCalculateStrategy: flow.Direct, //表示直接使用规则中的 Threshold 表示当前统计周期内的最大Token数量
		ControlBehavior:        flow.Reject, //表示如果当前统计周期内，统计结构统计的请求数超过了阈值，就直接拒绝。
		Threshold:              500,         //表示流控阈值
		StatIntervalInMs:       1000,        //规则对应的流量控制器的独立统计结构的统计周期。如果StatIntervalInMs是1000，也就是统计QPS
	}
}
