package sentinel

import (
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
)

// Flow 限流
type Flow func(o *flow.Rule)

// Breaker 熔断降级
type Breaker func(o *circuitbreaker.Rule)

// WithResource 资源名，即规则的作用目标。
func WithResource(serviceName string) Flow {
	return func(o *flow.Rule) {
		o.Resource = serviceName
	}
}

// WithTokenWarmUp WarmUp表示通过预热的方式计算当前统计周期内的最大Token数量，
// 预热的计算方式会根据规则中的字段 WarmUpPeriodSec 和 WarmUpColdFactor 来决定预热的曲线
func WithTokenWarmUp() Flow {
	return func(o *flow.Rule) {
		o.TokenCalculateStrategy = flow.WarmUp
	}
}

// WithControlThrottling 表示匀速排队的统计策略。它的中心思想是，以固定的间隔时间让请求通过。当请求到来的时候，
// 如果当前请求距离上个通过的请求通过的时间间隔不小于预设值，则让当前请求通过；
// 否则，计算当前请求的预期通过时间，如果该请求的预期通过时间小于规则预设的 timeout 时间，
// 则该请求会等待直到预设时间到来通过（排队等待处理）；若预期的通过时间超出最大排队时长，则直接拒接这个请求。
func WithControlThrottling() Flow {
	return func(o *flow.Rule) {
		o.ControlBehavior = flow.Throttling
	}
}

// WithThreshold 表示流控阈值；如果字段 StatIntervalInMs 是1000(也就是1秒)，
// 那么Threshold就表示QPS，流量控制器也就会依据资源的QPS来做流控
func WithThreshold(threshold float64) Flow {
	return func(o *flow.Rule) {
		if threshold > 0 {
			o.Threshold = threshold
		}
	}
}

// WithStatIntervalInMs 规则对应的流量控制器的独立统计结构的统计周期。如果StatIntervalInMs是1000，也就是统计QPS。
func WithStatIntervalInMs(ms uint32) Flow {
	return func(o *flow.Rule) {
		o.StatIntervalInMs = ms
	}
}

// WithWarmUpPeriodSec 预热的时间长度，该字段仅仅对 WarmUp 的TokenCalculateStrategy生效
func WithWarmUpPeriodSec(sec uint32) Flow {
	return func(o *flow.Rule) {
		o.WarmUpPeriodSec = sec
	}
}

// WithWarmUpColdFactor 预热的因子，默认是3，该值的设置会影响预热的速度
func WithWarmUpColdFactor(factor uint32) Flow {
	return func(o *flow.Rule) {
		o.WarmUpColdFactor = factor
	}
}

// WithBreakerResource 资源名，即规则的作用目标。
func WithBreakerResource(serviceName string) Breaker {
	return func(o *circuitbreaker.Rule) {
		o.Resource = serviceName
	}
}

// WithBreakerRetryTimeoutMs 熔断触发后持续的时间（单位为 ms）
// 资源进入熔断状态后，在配置的熔断时长内，请求都会快速失败。熔断结束后进入探测恢复模式（HALF-OPEN）
func WithBreakerRetryTimeoutMs(ms uint32) Breaker {
	return func(o *circuitbreaker.Rule) {
		o.RetryTimeoutMs = ms
	}
}

// WithBreakerMinRequestAmount 静默数量，如果当前统计周期内对资源的访问数量小于静默数量，那么熔断器就处于静默期。
// 换言之，也就是触发熔断的最小请求数目，若当前统计周期内的请求数小于此值，即使达到熔断条件规则也不会触发
func WithBreakerMinRequestAmount(amount uint64) Breaker {
	return func(o *circuitbreaker.Rule) {
		o.MinRequestAmount = amount
	}
}

// WithBreakerStatIntervalMs 统计的时间窗口长度（单位为 ms）
func WithBreakerStatIntervalMs(ms uint32) Breaker {
	return func(o *circuitbreaker.Rule) {
		o.StatIntervalMs = ms
	}
}

// WithBreakerThreshold 如果当前资源的错误比例如果高于Threshold，那么熔断器就会断开,
// Threshold表示的是错误比例的阈值(小数表示，比如0.1表示10%)
func WithBreakerThreshold(threshold float64) Breaker {
	return func(o *circuitbreaker.Rule) {
		if threshold > 0 {
			o.Threshold = threshold
		}
	}
}
