package tracing

import (
	"context"
	"github.com/binbinly/pkg/util"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var commonAttrs = []attribute.KeyValue{
	attribute.String("hostname", util.Hostname()),
	attribute.String("local-ip", localIP),
}

func setClientSpan(ctx context.Context, span trace.Span, m any) {
	attrs := commonAttrs
	if p, ok := m.(proto.Message); ok {
		attrs = append(attrs, attribute.Key("send_msg.size").Int(proto.Size(p)))
	}

	span.SetAttributes(attrs...)
}

func setServerSpan(ctx context.Context, span trace.Span, m any) {
	attrs := commonAttrs
	if p, ok := m.(proto.Message); ok {
		attrs = append(attrs, attribute.Key("recv_msg.size").Int(proto.Size(p)))
	}

	span.SetAttributes(attrs...)
}
