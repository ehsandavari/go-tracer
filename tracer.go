package tracer

import (
	"context"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type sTracer struct {
	config         *SConfig
	tracer         trace.Tracer
	span           trace.Span
	tracerProvider *traceSdk.TracerProvider
	jaegerExporter traceSdk.SpanExporter
	stdExporter    traceSdk.SpanExporter
}

func NewTracer(isEnabled bool, host string, port string, serviceId int, serviceName string, serviceVersion string, serviceMode string, sampler bool, useStdout bool) ITracer {
	structTracer := &sTracer{
		config: &SConfig{
			IsEnabled:      isEnabled,
			Host:           host,
			Port:           port,
			serviceId:      serviceId,
			serviceName:    serviceName,
			serviceVersion: serviceVersion,
			serviceMode:    serviceMode,
			Sampler:        sampler,
			UseStdout:      useStdout,
		},
	}

	if err := structTracer.configExporters(); err != nil {
		log.Fatalln("error in config exporters : ", err)
	}

	if err := structTracer.configTracerProviders(); err != nil {
		log.Fatalln("error in config tracer providers : ", err)
	}

	propagators := []propagation.TextMapPropagator{
		ot.OT{},
		propagation.TraceContext{},
		propagation.Baggage{},
	}
	otel.SetTracerProvider(structTracer.tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagators...))
	return structTracer
}

func (r *sTracer) configExporters() error {
	jaegerExporter, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(r.config.Host),
		jaeger.WithAgentPort(r.config.Port),
	))
	if err != nil {
		return err
	}
	r.jaegerExporter = jaegerExporter
	if r.config.UseStdout {
		stdExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return err
		}
		r.stdExporter = stdExporter
	}

	return nil
}

func (r *sTracer) configTracerProviders() error {
	sampler := traceSdk.NeverSample()
	if r.config.Sampler {
		sampler = traceSdk.AlwaysSample()
	}

	newResource, err := resource.New(
		context.Background(),
		resource.WithOS(),
		resource.WithProcess(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithHostID(),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceName(r.config.serviceName),
			semconv.ServiceVersion(r.config.serviceVersion),
			attribute.Int("service.id", r.config.serviceId),
			attribute.String("service.mode", r.config.serviceMode),
		),
	)
	if err != nil {
		return err
	}

	resourceMerged, err := resource.Merge(
		resource.Default(),
		newResource,
	)
	if err != nil {
		return err
	}

	tracerProvider := traceSdk.NewTracerProvider(
		traceSdk.WithBatcher(r.jaegerExporter),
		traceSdk.WithBatcher(r.stdExporter),
		traceSdk.WithSampler(sampler),
		traceSdk.WithResource(resourceMerged),
	)
	r.tracerProvider = tracerProvider
	return nil
}

func (r *sTracer) Shutdown() error {
	return r.tracerProvider.Shutdown(context.Background())
}

func (r *sTracer) Tracer(name string) ISpan {
	r.tracer = otel.Tracer(name)
	return r
}

func (r *sTracer) Start(ctx context.Context, spanName string) context.Context {
	ctx, r.span = r.tracer.Start(ctx, spanName)
	return ctx
}

func (r *sTracer) End() {
	r.span.End()
}
