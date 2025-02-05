package entity

import (
	"errors"
)

var (
	ErrNoResult      = errors.New("no result")
	ErrAlreadyClosed = errors.New("already closed")

	ErrOverflow     = errors.New("overflow")
	ErrInvalidValue = errors.New("invalid value")

	ErrTracerServiceNameRequired = errors.New("tracer: service name required")
	ErrTracerEndpointRequired    = errors.New("tracer: endpoint required")

	ErrRecursive = errors.New("recursive_error")
)
