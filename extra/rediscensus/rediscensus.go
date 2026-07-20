package rediscensus

import (
	"context"

	"go.opencensus.io/trace"

	"github.com/redis/go-redis/v9"
)

type TracingHook struct{}

var _ redis.Hook = (*TracingHook)(nil)

func NewTracingHook() *TracingHook { _ = "STUB: not implemented"; return nil }

func (TracingHook) DialHook(next redis.DialHook) redis.DialHook {
	_ = "STUB: not implemented"
	return *new(redis.DialHook)
}

func (TracingHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessHook)
}

func (TracingHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessPipelineHook)
}

func recordErrorOnOCSpan(ctx context.Context, span *trace.Span, err error) {
	_ = "STUB: not implemented"
	return
}
