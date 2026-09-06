package maintnotifications

import (
	"context"

	"github.com/redis/go-redis/v9/push"
)

type LoggingHook struct {
	LogLevel int
}

func (lh *LoggingHook) PreHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}) ([]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (lh *LoggingHook) PostHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}, result error) {
	_ = "STUB: not implemented"
	return
}

func NewLoggingHook(logLevel int) *LoggingHook { _ = "STUB: not implemented"; return nil }
