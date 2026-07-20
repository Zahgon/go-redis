package redisotel

import (
	"sync"

	"go.opentelemetry.io/otel/metric"
)

const (
	MetricOperationDuration = "db.client.operation.duration"

	MetricConnectionCount          = "db.client.connection.count"
	MetricConnectionCreateTime     = "db.client.connection.create_time"
	MetricConnectionWaitTime       = "db.client.connection.wait_time"
	MetricConnectionPendingReqs    = "db.client.connection.pending_requests"
	MetricConnectionRelaxedTimeout = "redis.client.connection.relaxed_timeout"
	MetricConnectionHandoff        = "redis.client.connection.handoff"
	MetricConnectionClosed         = "redis.client.connection.closed"

	MetricClientErrors             = "redis.client.errors"
	MetricMaintenanceNotifications = "redis.client.maintenance.notifications"

	MetricPubSubMessages = "redis.client.pubsub.messages"

	MetricStreamLag = "redis.client.stream.lag"

	PoolNameMain   = "main"
	PoolNamePubSub = "pubsub"
)

var (
	observabilityInstance     *ObservabilityInstance
	observabilityInstanceOnce sync.Once
)

type ObservabilityInstance struct {
	mu          sync.RWMutex
	config      *Config
	recorder    *metricsRecorder
	initialized bool
}

func GetObservabilityInstance() *ObservabilityInstance { _ = "STUB: not implemented"; return nil }

func (o *ObservabilityInstance) Init(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func (o *ObservabilityInstance) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (o *ObservabilityInstance) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (o *ObservabilityInstance) shutdownLocked() error { _ = "STUB: not implemented"; return nil }

func (o *ObservabilityInstance) configToInternal(cfg *Config) config {
	_ = "STUB: not implemented"
	return *new(config)
}

func (o *ObservabilityInstance) createRecorder(meter metric.Meter, cfg config) (*metricsRecorder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePoolName(poolName string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}
