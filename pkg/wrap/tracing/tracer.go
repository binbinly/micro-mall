package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Tracer is otel span tracer
type Tracer struct {
	tracer trace.Tracer
	kind   trace.SpanKind
	opt    *options
}

// NewTracer create tracer instance
func NewTracer(kind trace.SpanKind, opts ...Option) *Tracer {
	o := options{
		ServiceName:    "mall",
		TracerProvider: otel.GetTracerProvider(),
		Propagators:    otel.GetTextMapPropagator(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	return &Tracer{tracer: otel.Tracer(o.ServiceName), kind: kind, opt: &o}
}

// Start start tracing span
func (t *Tracer) Start(ctx context.Context, operation string) (context.Context, trace.Span) {
	if t.kind == trace.SpanKindServer {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = make(metadata.MD)
		}
		ctx = t.opt.Propagators.Extract(ctx, MetadataCarrier(md))
	}
	ctx, span := t.tracer.Start(ctx,
		operation,
		trace.WithSpanKind(t.kind),
		trace.WithAttributes(
			semconv.RPCSystemKey.String("grpc"),
		),
	)
	if t.kind == trace.SpanKindClient {
		md, _ := metadata.FromOutgoingContext(ctx)
		mdCopy := md.Copy()
		t.opt.Propagators.Inject(ctx, MetadataCarrier(mdCopy))
		//ctx = metadata.NewOutgoingContext(ctx, mdCopy)
	}
	return ctx, span
}

// End finish tracing span
func (t *Tracer) End(ctx context.Context, span trace.Span, m interface{}, err error) {
	if err != nil {
		s, _ := status.FromError(err)
		span.SetStatus(codes.Error, s.Message())
		span.SetAttributes(semconv.RPCGRPCStatusCodeKey.String(s.Code().String()))
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	if p, ok := m.(proto.Message); ok {
		if t.kind == trace.SpanKindServer {
			span.SetAttributes(attribute.Key("send_msg.size").Int(proto.Size(p)))
		} else {
			span.SetAttributes(attribute.Key("recv_msg.size").Int(proto.Size(p)))
		}
	}
}
