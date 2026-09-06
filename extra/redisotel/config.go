package redisotel

import (
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type config struct {
	dbSystem string
	attrs    []attribute.KeyValue

	tp     trace.TracerProvider
	tracer trace.Tracer

	dbStmtEnabled         bool
	callerEnabled         bool
	filterDial            bool
	filterProcessPipeline func(cmds []redis.Cmder) bool
	filterProcess         func(cmd redis.Cmder) bool

	mp    metric.MeterProvider
	meter metric.Meter

	semconvCompliantMetrics bool

	poolName string

	closeChan chan struct{}
}

type baseOption interface {
	apply(conf *config)
}

type Option interface {
	baseOption
	tracing()
	metrics()
}

type option func(conf *config)

func (fn option) apply(conf *config) { _ = "STUB: not implemented"; return }

func (fn option) tracing() { _ = "STUB: not implemented"; return }

func (fn option) metrics() { _ = "STUB: not implemented"; return }

func newConfig(opts ...baseOption) *config { _ = "STUB: not implemented"; return nil }

func WithDBSystem(dbSystem string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAttributes(attrs ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPoolName(poolName string) Option { _ = "STUB: not implemented"; return *new(Option) }

type TracingOption interface {
	baseOption
	tracing()
}

type tracingOption func(conf *config)

var _ TracingOption = (*tracingOption)(nil)

func (fn tracingOption) apply(conf *config) { _ = "STUB: not implemented"; return }

func (fn tracingOption) tracing() { _ = "STUB: not implemented"; return }

func WithTracerProvider(provider trace.TracerProvider) TracingOption {
	_ = "STUB: not implemented"
	return *new(TracingOption)
}

func WithDBStatement(on bool) TracingOption { _ = "STUB: not implemented"; return *new(TracingOption) }

func WithCallerEnabled(on bool) TracingOption {
	_ = "STUB: not implemented"
	return *new(TracingOption)
}

func WithCommandFilter(filter func(cmd redis.Cmder) bool) TracingOption {
	_ = "STUB: not implemented"
	return *new(TracingOption)
}

func WithCommandsFilter(filter func(cmds []redis.Cmder) bool) TracingOption {
	_ = "STUB: not implemented"
	return *new(TracingOption)
}

func WithDialFilter(on bool) TracingOption { _ = "STUB: not implemented"; return *new(TracingOption) }

func DefaultCommandFilter(cmd redis.Cmder) bool { _ = "STUB: not implemented"; return false }

func BasicCommandFilter(cmd redis.Cmder) bool { _ = "STUB: not implemented"; return false }

type MetricsOption interface {
	baseOption
	metrics()
}

type metricsOption func(conf *config)

var _ MetricsOption = (*metricsOption)(nil)

func (fn metricsOption) apply(conf *config) { _ = "STUB: not implemented"; return }

func (fn metricsOption) metrics() { _ = "STUB: not implemented"; return }

func WithMeterProvider(mp metric.MeterProvider) MetricsOption {
	_ = "STUB: not implemented"
	return *new(MetricsOption)
}

func WithCloseChan(closeChan chan struct{}) MetricsOption {
	_ = "STUB: not implemented"
	return *new(MetricsOption)
}

func WithSemConvCompliantMetrics(on bool) MetricsOption {
	_ = "STUB: not implemented"
	return *new(MetricsOption)
}
