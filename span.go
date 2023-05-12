package tracer

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type ISpan interface {
	Start(ctx context.Context, spanName string) context.Context
	AddEvent(name string)
	IsRecording() bool
	RecordError(err error)
	TraceId() [16]byte
	SetStatus(code codes.Code, description string)
	IAttribute
	End()
}

func (r *sTracer) Start(ctx context.Context, spanName string) context.Context {
	ctx, r.span = r.tracer.Start(ctx, spanName, trace.WithStackTrace(status))
	return ctx
}

func (r *sTracer) AddEvent(name string) {
	r.span.AddEvent(name, trace.WithStackTrace(status))
}

func (r *sTracer) IsRecording() bool {
	return r.span.IsRecording()
}

func (r *sTracer) RecordError(err error) {
	r.span.RecordError(err, trace.WithStackTrace(status))
}

func (r *sTracer) TraceId() [16]byte {
	return r.span.SpanContext().TraceID()
}

func (r *sTracer) SetStatus(code codes.Code, description string) {
	r.span.SetStatus(code, description)
}

func (r *sTracer) End() {
	r.span.End()
}
