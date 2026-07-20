package maintnotifications

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/redis/go-redis/v9/internal/pool"
)

type OperationsManagerInterface interface {
	TrackMovingOperationWithConnID(ctx context.Context, newEndpoint string, deadline time.Time, seqID int64, connID uint64) error
	UntrackOperationWithConnID(seqID int64, connID uint64)
}

type maintNotificationsConnTracker interface {
	UntrackMaintNotificationsConn(connID uint64)
}

type HandoffRequest struct {
	Conn     *pool.Conn
	ConnID   uint64
	Endpoint string
	SeqID    int64
	Pool     pool.Pooler
}

type PoolHook struct {
	baseDialer func(context.Context, string, string) (net.Conn, error)

	network string

	workerManager *handoffWorkerManager

	config *Config

	operationsManager OperationsManagerInterface

	pool pool.Pooler
}

func NewPoolHook(baseDialer func(context.Context, string, string) (net.Conn, error), network string, config *Config, operationsManager OperationsManagerInterface) *PoolHook {
	_ = "STUB: not implemented"
	return nil
}

func NewPoolHookWithPoolSize(baseDialer func(context.Context, string, string) (net.Conn, error), network string, config *Config, operationsManager OperationsManagerInterface, poolSize int) *PoolHook {
	_ = "STUB: not implemented"
	return nil
}

func (ph *PoolHook) SetPool(pooler pool.Pooler) { _ = "STUB: not implemented"; return }

func (ph *PoolHook) GetCurrentWorkers() int { _ = "STUB: not implemented"; return 0 }

func (ph *PoolHook) IsHandoffPending(conn *pool.Conn) bool { _ = "STUB: not implemented"; return false }

func (ph *PoolHook) GetPendingMap() *sync.Map { _ = "STUB: not implemented"; return nil }

func (ph *PoolHook) GetMaxWorkers() int { _ = "STUB: not implemented"; return 0 }

func (ph *PoolHook) GetHandoffQueue() chan HandoffRequest { _ = "STUB: not implemented"; return nil }

func (ph *PoolHook) GetCircuitBreakerStats() []CircuitBreakerStats {
	_ = "STUB: not implemented"
	return nil
}

func (ph *PoolHook) ResetCircuitBreakers() { _ = "STUB: not implemented"; return }

func (ph *PoolHook) OnGet(_ context.Context, conn *pool.Conn, _ bool) (accept bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ph *PoolHook) OnPut(ctx context.Context, conn *pool.Conn) (shouldPool bool, shouldRemove bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (ph *PoolHook) OnRemove(_ context.Context, conn *pool.Conn, _ error) {
	_ = "STUB: not implemented"
	return
}

func (ph *PoolHook) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
