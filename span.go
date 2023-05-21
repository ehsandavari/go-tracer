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
	value, ok := ctx.Value(RequestId).(string)
	if ok {
		r.SetString(RequestId, value)
	}
	return context.WithValue(ctx, TraceId, r.TraceId())
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

func (r *sTracer) TraceId() string {
	return r.span.SpanContext().TraceID().String()
}

func (r *sTracer) SetStatus(code codes.Code, description string) {
	r.span.SetStatus(code, description)
}

func (r *sTracer) End() {
	r.span.End()
}
