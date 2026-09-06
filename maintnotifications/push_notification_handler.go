package maintnotifications

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/push"
)

type NotificationHandler struct {
	manager           *Manager
	operationsManager OperationsManagerInterface
}

func (snh *NotificationHandler) HandlePushNotification(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleMoving(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) markConnForHandoff(conn *pool.Conn, newEndpoint string, seqID int64, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleMigrating(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleMigrated(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleFailingOver(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleFailedOver(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleSMigrating(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (snh *NotificationHandler) handleSMigrated(ctx context.Context, handlerCtx push.NotificationHandlerContext, notification []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
