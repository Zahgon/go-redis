package redisotel

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type timeoutError interface {
	Timeout() bool
}

const (
	libraryName = "go-redis"
)

func getLibraryVersionAttr() attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func addServerPortIfNonDefault(attrs []attribute.KeyValue, serverPort string) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

type metricsRecorder struct {
	operationDuration        metric.Float64Histogram
	connectionCount          metric.Int64UpDownCounter
	connectionCreateTime     metric.Float64Histogram
	connectionRelaxedTimeout metric.Int64UpDownCounter
	connectionHandoff        metric.Int64Counter
	clientErrors             metric.Int64Counter
	maintenanceNotifications metric.Int64Counter

	connectionWaitTime    metric.Float64Histogram
	connectionClosed      metric.Int64Counter
	connectionPendingReqs metric.Int64UpDownCounter

	pubsubMessages metric.Int64Counter

	streamLag metric.Float64Histogram

	cfg *config
}

func (r *metricsRecorder) RecordOperationDuration(
	ctx context.Context,
	duration time.Duration,
	cmd redis.Cmder,
	attempts int,
	err error,
	cn redis.ConnInfo,
	dbIndex int,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordPipelineOperationDuration(
	ctx context.Context,
	duration time.Duration,
	operationName string,
	cmdCount int,
	attempts int,
	err error,
	cn redis.ConnInfo,
	dbIndex int,
) {
	_ = "STUB: not implemented"
	return
}

func classifyError(err error) string { _ = "STUB: not implemented"; return "" }

func normalizeNetworkError(err error) string { _ = "STUB: not implemented"; return "" }

func normalizeGenericError(errStr string) string { _ = "STUB: not implemented"; return "" }

func extractRedisErrorPrefix(err error) string { _ = "STUB: not implemented"; return "" }

func isNetworkError(err error) bool { _ = "STUB: not implemented"; return false }

func isTimeoutError(err error) bool { _ = "STUB: not implemented"; return false }

func getErrorCategory(err error) string { _ = "STUB: not implemented"; return "" }

func getErrorCategoryFromString(errStr string) string { _ = "STUB: not implemented"; return "" }

func splitHostPort(addr string) (host, port string) { _ = "STUB: not implemented"; return "", "" }

func parseAddr(addr string) (host, port string) { _ = "STUB: not implemented"; return "", "" }

func extractServerInfo(cn redis.ConnInfo) (addr, port string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (r *metricsRecorder) RecordConnectionCreateTime(
	ctx context.Context,
	duration time.Duration,
	cn redis.ConnInfo,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordConnectionRelaxedTimeout(
	ctx context.Context,
	delta int,
	cn redis.ConnInfo,
	poolName, notificationType string,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordConnectionHandoff(
	ctx context.Context,
	cn redis.ConnInfo,
	poolName string,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordError(
	ctx context.Context,
	errorType string,
	cn redis.ConnInfo,
	statusCode string,
	isInternal bool,
	retryAttempts int,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordMaintenanceNotification(
	ctx context.Context,
	cn redis.ConnInfo,
	notificationType string,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordConnectionWaitTime(
	ctx context.Context,
	duration time.Duration,
	cn redis.ConnInfo,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordConnectionClosed(
	ctx context.Context,
	cn redis.ConnInfo,
	reason string,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordPubSubMessage(
	ctx context.Context,
	cn redis.ConnInfo,
	direction string,
	channel string,
	sharded bool,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordStreamLag(
	ctx context.Context,
	lag time.Duration,
	cn redis.ConnInfo,
	streamName string,
	consumerGroup string,
	consumerName string,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordConnectionCount(
	ctx context.Context,
	delta int,
	cn redis.ConnInfo,
	state string,
	isPubSub bool,
) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsRecorder) RecordPendingRequests(
	ctx context.Context,
	delta int,
	cn redis.ConnInfo,
	poolName string,
) {
	_ = "STUB: not implemented"
	return
}
