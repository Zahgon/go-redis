package maintnotifications

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9/internal/interfaces"
	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/push"
)

const (
	NotificationMoving      = "MOVING"
	NotificationMigrating   = "MIGRATING"
	NotificationMigrated    = "MIGRATED"
	NotificationFailingOver = "FAILING_OVER"
	NotificationFailedOver  = "FAILED_OVER"
	NotificationSMigrating  = "SMIGRATING"
	NotificationSMigrated   = "SMIGRATED"
)

var maintenanceNotificationTypes = []string{
	NotificationMoving,
	NotificationMigrating,
	NotificationMigrated,
	NotificationFailingOver,
	NotificationFailedOver,
	NotificationSMigrating,
	NotificationSMigrated,
}

type NotificationHook interface {
	PreHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}) ([]interface{}, bool)
	PostHook(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}, result error)
}

type MovingOperationKey struct {
	SeqID  int64
	ConnID uint64
}

func (k MovingOperationKey) String() string { _ = "STUB: not implemented"; return "" }

type Manager struct {
	client  interfaces.ClientInterface
	config  *Config
	options interfaces.OptionsInterface
	pool    pool.Pooler

	activeMovingOps sync.Map

	processedSMigratedSeqIDs sync.Map

	activeOperationCount atomic.Int64
	closed               atomic.Bool

	shutdownTimeout time.Duration

	hooks        []NotificationHook
	hooksMu      sync.RWMutex
	poolHooksRef *PoolHook

	additionalPoolHooks []additionalPoolHook

	maintNotificationsConns sync.Map

	clusterStateReloadCallback atomic.Pointer[ClusterStateReloadCallback]
}

type MovingOperation struct {
	SeqID       int64
	NewEndpoint string
	StartTime   time.Time
	Deadline    time.Time
}

type ClusterStateReloadCallback func(ctx context.Context, hostPort string, slotRanges []string)

func NewManager(client interfaces.ClientInterface, pool pool.Pooler, config *Config) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hm *Manager) InitPoolHook(baseDialer func(context.Context, string, string) (net.Conn, error)) {
	_ = "STUB: not implemented"
	return
}

type additionalPoolHook struct {
	pool pool.Pooler
	hook *PoolHook
}

func (hm *Manager) InitPoolHookForPool(p pool.Pooler, baseDialer func(context.Context, string, string) (net.Conn, error)) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) hookForConn(cn *pool.Conn) *PoolHook { _ = "STUB: not implemented"; return nil }

func (hm *Manager) setupPushNotifications() error { _ = "STUB: not implemented"; return nil }

func (hm *Manager) TrackMovingOperationWithConnID(ctx context.Context, newEndpoint string, deadline time.Time, seqID int64, connID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (hm *Manager) UntrackOperationWithConnID(seqID int64, connID uint64) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) GetActiveMovingOperations() map[MovingOperationKey]*MovingOperation {
	_ = "STUB: not implemented"
	return nil
}

func (hm *Manager) IsHandoffInProgress() bool { _ = "STUB: not implemented"; return false }

func (hm *Manager) GetActiveOperationCount() int64 { _ = "STUB: not implemented"; return 0 }

func (hm *Manager) MarkSMigratedSeqIDProcessed(seqID int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (hm *Manager) TrackMaintNotificationsConn(cn *pool.Conn) { _ = "STUB: not implemented"; return }

func (hm *Manager) UntrackMaintNotificationsConn(connID uint64) { _ = "STUB: not implemented"; return }

func (hm *Manager) maintNotificationsConnSnapshot() []*pool.Conn {
	_ = "STUB: not implemented"
	return nil
}

func (hm *Manager) retireMaintNotificationsConns(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) Close() error { _ = "STUB: not implemented"; return nil }

func (hm *Manager) GetState() State { _ = "STUB: not implemented"; return *new(State) }

func (hm *Manager) processPreHooks(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}) ([]interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (hm *Manager) processPostHooks(ctx context.Context, notificationCtx push.NotificationHandlerContext, notificationType string, notification []interface{}, result error) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) createPoolHook(baseDialer func(context.Context, string, string) (net.Conn, error)) *PoolHook {
	_ = "STUB: not implemented"
	return nil
}

func (hm *Manager) AddNotificationHook(notificationHook NotificationHook) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) SetClusterStateReloadCallback(callback ClusterStateReloadCallback) {
	_ = "STUB: not implemented"
	return
}

func (hm *Manager) TriggerClusterStateReload(ctx context.Context, hostPort string, slotRanges []string) {
	_ = "STUB: not implemented"
	return
}
