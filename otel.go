package redis

import (
	"context"
	"net"
	"time"

	"github.com/redis/go-redis/v9/internal/otel"
	"github.com/redis/go-redis/v9/internal/pool"
)

type ConnInfo interface {
	RemoteAddr() net.Addr
	PoolName() string
}

type Pooler interface {
	PoolStats() *pool.Stats
}

type PubSubPooler interface {
	Stats() *pool.PubSubStats
}

type OTelRecorder interface {
	RecordOperationDuration(ctx context.Context, duration time.Duration, cmd Cmder, attempts int, err error, cn ConnInfo, dbIndex int)

	RecordPipelineOperationDuration(ctx context.Context, duration time.Duration, operationName string, cmdCount int, attempts int, err error, cn ConnInfo, dbIndex int)

	RecordConnectionCreateTime(ctx context.Context, duration time.Duration, cn ConnInfo)

	RecordConnectionRelaxedTimeout(ctx context.Context, delta int, cn ConnInfo, poolName, notificationType string)

	RecordConnectionHandoff(ctx context.Context, cn ConnInfo, poolName string)

	RecordError(ctx context.Context, errorType string, cn ConnInfo, statusCode string, isInternal bool, retryAttempts int)

	RecordMaintenanceNotification(ctx context.Context, cn ConnInfo, notificationType string)

	RecordConnectionWaitTime(ctx context.Context, duration time.Duration, cn ConnInfo)

	RecordConnectionClosed(ctx context.Context, cn ConnInfo, reason string, err error)

	RecordPubSubMessage(ctx context.Context, cn ConnInfo, direction, channel string, sharded bool)

	RecordStreamLag(ctx context.Context, lag time.Duration, cn ConnInfo, streamName, consumerGroup, consumerName string)
}

type OTelConnectionCounter interface {
	RecordConnectionCount(ctx context.Context, delta int, cn ConnInfo, state string, isPubSub bool)

	RecordPendingRequests(ctx context.Context, delta int, cn ConnInfo, poolName string)
}

type OTelPoolRegistrar interface {
	RegisterPool(poolName string, pool Pooler)

	UnregisterPool(pool Pooler)

	RegisterPubSubPool(poolName string, pool PubSubPooler)

	UnregisterPubSubPool(pool PubSubPooler)
}

func SetOTelRecorder(r OTelRecorder) { _ = "STUB: not implemented"; return }

type otelRecorderAdapter struct {
	recorder OTelRecorder
}

func toConnInfo(cn *pool.Conn) ConnInfo { _ = "STUB: not implemented"; return *new(ConnInfo) }

func (a *otelRecorderAdapter) RecordOperationDuration(ctx context.Context, duration time.Duration, cmd otel.Cmder, attempts int, err error, cn *pool.Conn, dbIndex int) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordPipelineOperationDuration(ctx context.Context, duration time.Duration, operationName string, cmdCount int, attempts int, err error, cn *pool.Conn, dbIndex int) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionCreateTime(ctx context.Context, duration time.Duration, cn *pool.Conn) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionRelaxedTimeout(ctx context.Context, delta int, cn *pool.Conn, poolName, notificationType string) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionHandoff(ctx context.Context, cn *pool.Conn, poolName string) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordError(ctx context.Context, errorType string, cn *pool.Conn, statusCode string, isInternal bool, retryAttempts int) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordMaintenanceNotification(ctx context.Context, cn *pool.Conn, notificationType string) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionWaitTime(ctx context.Context, duration time.Duration, cn *pool.Conn) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionClosed(ctx context.Context, cn *pool.Conn, reason string, err error) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordPubSubMessage(ctx context.Context, cn *pool.Conn, direction, channel string, sharded bool) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordStreamLag(ctx context.Context, lag time.Duration, cn *pool.Conn, streamName, consumerGroup, consumerName string) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordConnectionCount(ctx context.Context, delta int, cn *pool.Conn, state string, isPubSub bool) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RecordPendingRequests(ctx context.Context, delta int, cn *pool.Conn, poolName string) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) RegisterPool(poolName string, p pool.Pooler) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) UnregisterPool(p pool.Pooler) { _ = "STUB: not implemented"; return }

func (a *otelRecorderAdapter) RegisterPubSubPool(poolName string, p otel.PubSubPooler) {
	_ = "STUB: not implemented"
	return
}

func (a *otelRecorderAdapter) UnregisterPubSubPool(p otel.PubSubPooler) {
	_ = "STUB: not implemented"
	return
}

type poolerAdapter struct {
	p pool.Pooler
}

func (a *poolerAdapter) PoolStats() *pool.Stats { _ = "STUB: not implemented"; return nil }

type pubSubPoolerAdapter struct {
	p otel.PubSubPooler
}

func (a *pubSubPoolerAdapter) Stats() *pool.PubSubStats { _ = "STUB: not implemented"; return nil }
