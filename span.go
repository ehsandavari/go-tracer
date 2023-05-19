package tracer

import (
	"context"
	"go.opentelemetry.io/otel/codes"
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
	ctx, r.span = r.tracer.Start(ctx, spanName)
	return ctx
}

func (r *sTracer) AddEvent(name string) {
	r.span.AddEvent(name)
}

func (r *sTracer) IsRecording() bool {
	return r.span.IsRecording()
}

func (r *sTracer) RecordError(err error) {
	r.span.RecordError(err)
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
