package sentinel

import (
	"context"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitBreakerClientInterceptor 熔断降级
func CircuitBreakerClientInterceptor(opts ...Breaker) grpc.UnaryClientInterceptor {
	rule := defaultBreaker()
	for _, opt := range opts {
		opt(rule)
	}
	if _, err := circuitbreaker.LoadRules([]*circuitbreaker.Rule{rule}); err != nil {
		log.Fatalf("[Sentinel] load rules error: %+v", err)
	}
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		e, b := sentinel.Entry(rule.ResourceName(), sentinel.WithTrafficType(base.Inbound), sentinel.WithResourceType(base.ResTypeAPIGateway))
		if b != nil {
			return status.Error(codes.DeadlineExceeded, "Server fusing")
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			e.SetError(err)
		}
		e.Exit()
		return err
	}
}

// defaultBreaker 默认错误比例规则
func defaultBreaker() *circuitbreaker.Rule {
	return &circuitbreaker.Rule{
		Resource:                     "breaker",
		Strategy:                     circuitbreaker.ErrorRatio, //错误比例
		RetryTimeoutMs:               5000,                      //即熔断触发后持续的时间（单位为 ms）
		MinRequestAmount:             10,                        //触发熔断的最小请求数目
		StatIntervalMs:               10000,                     //统计的时间窗口长度（单位为 ms）
		StatSlidingWindowBucketCount: 1,                         //统计滑动窗口的个数。越大越精确，内存消耗也更大,默认 1
		Threshold:                    0.1,                       //熔断触发阈值
	}
}
