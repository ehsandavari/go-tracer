package tracer

import "context"

type (
	ITracer interface {
		Shutdown() error
		Tracer(name string) ISpan
	}
	ISpan interface {
		Start(ctx context.Context, spanName string) context.Context
		End()
	}
)
