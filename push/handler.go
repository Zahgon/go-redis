package push

import (
	"context"
)

type NotificationHandler interface {
	HandlePushNotification(ctx context.Context, handlerCtx NotificationHandlerContext, notification []interface{}) error
}
