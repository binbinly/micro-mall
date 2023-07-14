package app

import (
	"context"
	"log"

	"gateway/internal/interceptor/sentinel"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"gateway/internal/interceptor"
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
)

func newRPCClientConn(ctx context.Context, service Service) *grpc.ClientConn {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	// Setup metrics.
	reg := prometheus.NewRegistry()
	clMetrics := grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg.MustRegister(clMetrics)
	fromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	target := "consul:///" + service.Name
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			interceptor.TimeoutClientInterceptor(service.Timeout),
			interceptor.ValidatorClientInterceptor(),
			otelgrpc.UnaryClientInterceptor(),
			clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(fromContext)),
			sentinel.FlowClientInterceptor(
				sentinel.WithResource(service.Name),
				sentinel.WithThreshold(service.QPSLimit)),
			sentinel.CircuitBreakerClientInterceptor(sentinel.WithBreakerResource(service.Name)),
			logging.UnaryClientInterceptor(interceptor.Logger(), logging.WithFieldsFromContext(logTraceID))),
		grpc.WithInitialWindowSize(grpcInitialWindowSize),
		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize), grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			PermitWithoutStream: true, // 即使没有活跃的流也发送ping
		}),
	)
	if err != nil {
		log.Fatalf("failed new grpc client err: %v by target: %v", err, target)
	}

	return conn
}
