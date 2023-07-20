package tracer

import (
	contextplus "github.com/ehsandavari/go-context-plus"
	"go.opentelemetry.io/otel/codes"
)

type ISpan interface {
	Start(ctx contextplus.Context, spanName string) contextplus.Context
	AddEvent(name string)
	IsRecording() bool
	RecordError(err error)
	TraceId() string
	SetStatus(code codes.Code, description string)
	IAttribute
	End()
}

func (r *sTracer) Start(ctx contextplus.Context, spanName string) contextplus.Context {
	ctx.Context, r.span = r.tracer.Start(ctx.Context, spanName)
	requestId := ctx.RequestId()
	if len(requestId) != 0 {
		r.SetString("RequestId", requestId)
	}
	ctx.SetTraceId(r.TraceId())
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

func (r *sTracer) TraceId() string {
	return r.span.SpanContext().TraceID().String()
}

func (r *sTracer) SetStatus(code codes.Code, description string) {
	r.span.SetStatus(code, description)
}

func (r *sTracer) End() {
	r.span.End()
}
