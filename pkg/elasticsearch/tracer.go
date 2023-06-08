package elasticsearch

import (
	"bytes"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// Transport for tracing Elastic operations.
type Transport struct {
	rt      http.RoundTripper
	tracer  trace.Tracer
	logBody bool // show req.body content
}

// Option signature for specifying options, e.g. WithRoundTripper.
type Option func(t *Transport)

// WithRoundTripper specifies the http.RoundTripper to call
// next after this transport. If it is nil (default), the
// transport will use http.DefaultTransport.
func WithRoundTripper(rt http.RoundTripper) Option {
	return func(t *Transport) {
		t.rt = rt
	}
}

// WithLogBody with logBody true
func WithLogBody() Option {
	return func(t *Transport) {
		t.logBody = true
	}
}

// NewTransport specifies a transport that will trace Elastic
// and report back via OpenTracing.
func NewTransport(opts ...Option) *Transport {
	t := &Transport{tracer: otel.Tracer("elastic")}
	for _, o := range opts {
		o(t)
	}
	return t
}

// RoundTrip captures the request and starts an OpenTracing span
// for Elastic PerformRequest operation.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, span := t.tracer.Start(req.Context(), "elastic",
		trace.WithSpanKind(trace.SpanKindClient))
	req = req.WithContext(ctx)
	defer span.End()

	span.SetAttributes(
		attribute.String("component", "github.com/olivere/elastic/v7"),
		semconv.HTTPURLKey.String(req.URL.String()),
		semconv.HTTPMethodKey.String(req.Method),
		semconv.NetPeerNameKey.String(req.URL.Hostname()),
		semconv.NetPeerPortKey.String(req.URL.Port()))

	if t.logBody && req.Body != nil {
		b, _ := io.ReadAll(req.Body)
		span.AddEvent("body", trace.WithAttributes(
			attribute.String("req", string(b))))
		//重新写入
		req.Body = io.NopCloser(bytes.NewBuffer(b))
	}

	var (
		resp *http.Response
		err  error
	)
	if t.rt != nil {
		resp, err = t.rt.RoundTrip(req)
	} else {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}
	if err != nil {
		span.RecordError(err)
	}
	if resp != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(resp.StatusCode))
	}

	return resp, err
}
