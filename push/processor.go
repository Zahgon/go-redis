package push

import (
	"context"

	"github.com/redis/go-redis/v9/internal/proto"
)

type NotificationProcessor interface {
	GetHandler(pushNotificationName string) NotificationHandler

	ProcessPendingNotifications(ctx context.Context, handlerCtx NotificationHandlerContext, rd *proto.Reader) error

	RegisterHandler(pushNotificationName string, handler NotificationHandler, protected bool) error

	UnregisterHandler(pushNotificationName string) error
}

type Processor struct {
	registry *Registry
}

func NewProcessor() *Processor { _ = "STUB: not implemented"; return nil }

func (p *Processor) GetHandler(pushNotificationName string) NotificationHandler {
	_ = "STUB: not implemented"
	return *new(NotificationHandler)
}

func (p *Processor) RegisterHandler(pushNotificationName string, handler NotificationHandler, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Processor) UnregisterHandler(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Processor) ProcessPendingNotifications(ctx context.Context, handlerCtx NotificationHandlerContext, rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

type VoidProcessor struct{}

func NewVoidProcessor() *VoidProcessor { _ = "STUB: not implemented"; return nil }

func (v *VoidProcessor) GetHandler(_ string) NotificationHandler {
	_ = "STUB: not implemented"
	return *new(NotificationHandler)
}

func (v *VoidProcessor) RegisterHandler(pushNotificationName string, _ NotificationHandler, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VoidProcessor) UnregisterHandler(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *VoidProcessor) ProcessPendingNotifications(_ context.Context, handlerCtx NotificationHandlerContext, rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func willHandleNotificationInClient(notificationType string) bool {
	_ = "STUB: not implemented"
	return false
}
