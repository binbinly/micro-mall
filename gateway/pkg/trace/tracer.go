package trace

import (
	"strings"

	jaegerprop "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

//see: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go

// InitTracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func InitTracerProvider(options ...Option) (*tracesdk.TracerProvider, error) {
	var endpointOption jaeger.EndpointOption
	opts := applyOptions(options...)
	if strings.HasPrefix(opts.endpoint, "http") {
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opts.endpoint))
	} else {
		endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(opts.endpoint))
	}

	// Create the Jaeger exporter
	exporter, err := jaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// set sample
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(opts.samplingRatio)),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.serviceName),
			attribute.String("environment", opts.environment),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(jaegerprop.Jaeger{})
	//otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))
	return tp, nil
}
