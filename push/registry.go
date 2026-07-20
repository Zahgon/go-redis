package push

import (
	"sync"
)

type Registry struct {
	mu        sync.RWMutex
	handlers  map[string]NotificationHandler
	protected map[string]bool
}

func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

func (r *Registry) RegisterHandler(pushNotificationName string, handler NotificationHandler, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) GetHandler(pushNotificationName string) NotificationHandler {
	_ = "STUB: not implemented"
	return *new(NotificationHandler)
}

func (r *Registry) UnregisterHandler(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}
