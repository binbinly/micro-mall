package tracing

import (
	"context"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type traceWrapper struct {
	tracer *Tracer
	client.Client
}

func (l *traceWrapper) Call(ctx context.Context, req client.Request, rsp any, opts ...client.CallOption) error {
	var span trace.Span
	ctx, span = l.tracer.Start(ctx, req.Method())
	defer span.End()
	setClientSpan(ctx, span, req.Body())

	span.SetAttributes(
		semconv.RPCServiceKey.String(req.Service()),
		semconv.RPCMethodKey.String(req.Method()))

	err := l.Client.Call(ctx, req, rsp)
	l.tracer.End(ctx, span, rsp, err)
	return err
}

// NewClientWrapper accepts an open tracing Trace and returns a Client Wrapper
func NewClientWrapper(opts ...Option) client.Wrapper {
	tracer := NewTracer(trace.SpanKindClient, opts...)
	return func(c client.Client) client.Client {
		return &traceWrapper{
			tracer: tracer,
			Client: c,
		}
	}
}

// NewHandlerWrapper accepts an opentracing Tracer and returns a Handler Wrapper
func NewHandlerWrapper(opts ...Option) server.HandlerWrapper {
	tracer := NewTracer(trace.SpanKindServer, opts...)
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp any) error {
			var span trace.Span
			ctx, span = tracer.Start(ctx, req.Endpoint())
			defer span.End()
			setServerSpan(ctx, span, req.Body())

			span.SetAttributes(
				semconv.RPCServiceKey.String(req.Service()),
				semconv.RPCMethodKey.String(req.Method()))

			err := fn(ctx, req, rsp)
			tracer.End(ctx, span, rsp, err)
			return err
		}
	}
}

// Valuer is returns a log value.
type Valuer func(ctx context.Context) any

// TraceID returns a traceid valuer.
func TraceID() Valuer {
	return func(ctx context.Context) any {
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			return span.TraceID().String()
		}
		return ""
	}
}

// SpanID returns a spanid valuer.
func SpanID() Valuer {
	return func(ctx context.Context) any {
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			return span.SpanID().String()
		}
		return ""
	}
}
