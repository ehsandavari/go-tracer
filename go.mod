module github.com/ehsandavari/go-tracer

go 1.20

require (
	github.com/ehsandavari/go-context-plus v0.0.3
	github.com/google/uuid v1.3.0
	go.opentelemetry.io/contrib/propagators/ot v1.16.1
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/jaeger v1.15.1
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.15.1
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.15.1
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)
