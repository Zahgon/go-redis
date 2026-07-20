package redisotel

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/redis/go-redis/v9"
)

type metricsState struct {
	registrations []metric.Registration
	closed        bool
	mutex         sync.Mutex
}

func InstrumentMetrics(rdb redis.UniversalClient, opts ...MetricsOption) error {
	_ = "STUB: not implemented"
	return nil
}

func registerClient(rdb *redis.Client, conf *config, state *metricsState) error {
	_ = "STUB: not implemented"
	return nil
}

func poolStatsAttrs(conf *config) (poolAttrs, idleAttrs, usedAttrs attribute.Set) {
	_ = "STUB: not implemented"
	return *new(attribute.Set), *new(attribute.Set), *new(attribute.Set)
}

func reportPoolStats(rdb *redis.Client, conf *config) (metric.Registration, error) {
	_ = "STUB: not implemented"
	return *new(metric.Registration), nil
}

func addMetricsHook(rdb *redis.Client, conf *config) error { _ = "STUB: not implemented"; return nil }

type metricsHook struct {
	createTime metric.Float64Histogram
	useTime    metric.Float64Histogram
	attrs      []attribute.KeyValue
}

var _ redis.Hook = (*metricsHook)(nil)

func (mh *metricsHook) DialHook(hook redis.DialHook) redis.DialHook {
	_ = "STUB: not implemented"
	return *new(redis.DialHook)
}

func (mh *metricsHook) ProcessHook(hook redis.ProcessHook) redis.ProcessHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessHook)
}

func (mh *metricsHook) ProcessPipelineHook(
	hook redis.ProcessPipelineHook,
) redis.ProcessPipelineHook {
	_ = "STUB: not implemented"
	return *new(redis.ProcessPipelineHook)
}

func milliseconds(d time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

func statusAttr(err error) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func errorTypeAttribute(err error) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
