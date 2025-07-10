package logger

import "strings"

type environment string
type contextKey string
type traceKey string

const (
	production        environment = "production"
	development       environment = "development"
	staging           environment = "staging"
	test              environment = "test"
	defaultContextKey contextKey  = "x-context-id"
	defaultTraceKey   traceKey    = "x-trace-id"
)

type loggerOptions struct {
	env        string
	contextKey string
	traceKey   string
}

var defaultLoggerOptions = loggerOptions{
	env:        string(development),
	contextKey: string(defaultContextKey),
	traceKey:   string(defaultTraceKey),
}

type LoggerOption interface {
	apply(*loggerOptions)
}

type funcLoggerOption struct {
	f func(*loggerOptions)
}

func newFuncLoggerOption(f func(*loggerOptions)) *funcLoggerOption {
	return &funcLoggerOption{
		f: f,
	}
}

func (fnc *funcLoggerOption) apply(do *loggerOptions) {
	fnc.f(do)
}

// Environment returns a LoggerOption that sets the environment for
// the logging level.
//
// If this is not set, the default value is "development".
//
// # Possible Values
//   - production (info)
//   - staging (info)
//   - development (debug)
//   - test (discard)
func Environment(s string) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.env = strings.ToLower(s)
	})
}

// TraceKey is the key that is used for printing the trace ID in the log request.
//
// If this is not set, the default value is "x-trace-id".
func TraceKey(s string) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.traceKey = s
	})
}

// ContextKey is the key that is used for storing the trace ID of the log request in the context.
//
// If this is not set, the default value is "x-context-id".
func ContextKey(s string) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.contextKey = s
	})
}
