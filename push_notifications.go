package redis

import (
	"github.com/redis/go-redis/v9/push"
)

func NewPushNotificationProcessor() push.NotificationProcessor {
	_ = "STUB: not implemented"
	return *new(push.NotificationProcessor)
}

func NewVoidPushNotificationProcessor() push.NotificationProcessor {
	_ = "STUB: not implemented"
	return *new(push.NotificationProcessor)
}
