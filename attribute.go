package tracer

import "go.opentelemetry.io/otel/attribute"

type IAttribute interface {
	SetString(key string, value string)
	SetStringSlice(key string, value []string)
	SetInt(key string, value int)
	SetIntSlice(key string, value []int)
	SetInt64(key string, value int64)
	SetInt64Slice(key string, value []int64)
	SetBool(key string, value bool)
	SetBoolSlice(key string, value []bool)
	SetFloat64(key string, value float64)
	SetFloat64Slice(key string, value []float64)
}

func (r *sTracer) SetString(key string, value string) {
	r.span.SetAttributes(attribute.String(key, value))
}

func (r *sTracer) SetStringSlice(key string, value []string) {
	r.span.SetAttributes(attribute.StringSlice(key, value))
}

func (r *sTracer) SetInt(key string, value int) {
	r.span.SetAttributes(attribute.Int(key, value))
}

func (r *sTracer) SetIntSlice(key string, value []int) {
	r.span.SetAttributes(attribute.IntSlice(key, value))
}

func (r *sTracer) SetInt64(key string, value int64) {
	r.span.SetAttributes(attribute.Int64(key, value))
}

func (r *sTracer) SetInt64Slice(key string, value []int64) {
	r.span.SetAttributes(attribute.Int64Slice(key, value))
}

func (r *sTracer) SetBool(key string, value bool) {
	r.span.SetAttributes(attribute.Bool(key, value))
}

func (r *sTracer) SetBoolSlice(key string, value []bool) {
	r.span.SetAttributes(attribute.BoolSlice(key, value))
}

func (r *sTracer) SetFloat64(key string, value float64) {
	r.span.SetAttributes(attribute.Float64(key, value))
}

func (r *sTracer) SetFloat64Slice(key string, value []float64) {
	r.span.SetAttributes(attribute.Float64Slice(key, value))
}
