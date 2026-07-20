package maintnotifications

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9/push"
)

type contextKey string

const (
	startTimeKey contextKey = "maint_notif_start_time"
)

type MetricsHook struct {
	NotificationCounts map[string]int64
	ProcessingTimes    map[string]time.Duration
	ErrorCounts        map[string]int64
	HandoffCounts      int64
	HandoffSuccesses   int64
	HandoffFailures    int64
}

func NewMetricsHook() *MetricsHook { _ = "STUB: not implemented"; return nil }

func (mh *MetricsHook) PreHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}) ([]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mh *MetricsHook) PostHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}, result error) {
	_ = "STUB: not implemented"
	return
}

func (mh *MetricsHook) GetMetrics() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func ExampleCircuitBreakerMonitor(poolHook *PoolHook) { _ = "STUB: not implemented"; return }
