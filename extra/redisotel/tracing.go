package redisotel

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/redis/go-redis/v9"
)

const (
	instrumName = "github.com/redis/go-redis/extra/redisotel"
)

func InstrumentTracing(rdb redis.UniversalClient, opts ...TracingOption) error {
	_ = "STUB: not implemented"
	return nil
}

type tracingHook struct {
	conf *config

	spanOpts []trace.SpanStartOption
}

var _ redis.Hook = (*tracingHook)(nil)

func newTracingHook(connString string, opts ...TracingOption) *tracingHook {
	_ = "STUB: not implemented"
	return nil
}

func (th *tracingHook) DialHook(hook redis.DialHook) redis.DialHook {
	_ = "STUB: not implemented"
	return *new(redis.DialHook)
}

func (th *tracingHook) ProcessHook(hook redis.ProcessHook) redis.ProcessHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessHook)
}

func (th *tracingHook) ProcessPipelineHook(
	hook redis.ProcessPipelineHook,
) redis.ProcessPipelineHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessPipelineHook)
}

func recordError(span trace.Span, err error) { _ = "STUB: not implemented"; return }

func formatDBConnString(network, addr string) string { _ = "STUB: not implemented"; return "" }

func funcFileLine(pkg string) (string, string, int) { _ = "STUB: not implemented"; return "", "", 0 }

func addServerAttributes(opts []TracingOption, addr string) []TracingOption {
	_ = "STUB: not implemented"
	return nil
}
