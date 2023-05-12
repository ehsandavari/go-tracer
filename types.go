package tracer

import (
	"go.opentelemetry.io/otel/codes"
)

const (
	// Unset is the default status code.
	Unset codes.Code = 0

	// Error indicates the operation contains an error.
	//
	// NOTE: The error code in OTLP is 2.
	// The value of this enum is only relevant to the internals
	// of the Go SDK.
	Error codes.Code = 1

	// Ok indicates operation has been validated by an Application developers
	// or Operator to have completed successfully, or contain no error.
	//
	// NOTE: The Ok code in OTLP is 1.
	// The value of this enum is only relevant to the internals
	// of the Go SDK.
	Ok codes.Code = 2
)
