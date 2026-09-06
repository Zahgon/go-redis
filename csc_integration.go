package redis

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/push"
)

func cscRegisterCleanups(c *Client) { _ = "STUB: not implemented"; return }

type ClientSideCacheConfig = CacheConfig

const (
	invalidatePushName = "invalidate"

	cscNamespaceSep = "\x00"
)

func cscNamespacePrefix(db int, username string) string { _ = "STUB: not implemented"; return "" }

func cscNamespacedKey(prefix, key string) string { _ = "STUB: not implemented"; return "" }

type invalidateHandler struct {
	mu        sync.RWMutex
	cache     Cache
	keyPrefix string
	users     int
}

func (h *invalidateHandler) HandlePushNotification(
	_ context.Context, _ push.NotificationHandlerContext, notification []interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *invalidateHandler) release() { _ = "STUB: not implemented"; return }

func (h *invalidateHandler) releaseLocked() { _ = "STUB: not implemented"; return }

func sameCache(a, b Cache) bool { _ = "STUB: not implemented"; return false }

func isNilCache(cache Cache) bool { _ = "STUB: not implemented"; return false }

var errInvalidateHandlerBound = errors.New(`csc: a different "invalidate" push handler is already registered`)

func (h *invalidateHandler) bindTo(cache Cache, keyPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func lookupInvalidateHandler(p push.NotificationProcessor) *invalidateHandler {
	_ = "STUB: not implemented"
	return nil
}

func registerInvalidateHandler(p push.NotificationProcessor, cache Cache, keyPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) attachCSC(ctx context.Context, cache Cache) { _ = "STUB: not implemented"; return }

func (c *baseClient) attachSharedTrackingCSC(ctx context.Context, cache Cache) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) cscHook() *cscEvictOnRemoveHook { _ = "STUB: not implemented"; return nil }

func (c *baseClient) cscInstallConnCloseHook(cn *pool.Conn) { _ = "STUB: not implemented"; return }

func (c *baseClient) cscInstallConnReinitHook(cn *pool.Conn) { _ = "STUB: not implemented"; return }

func (c *baseClient) cscOnConnClose(connID uint64) { _ = "STUB: not implemented"; return }

type poolHookSupport interface {
	AddPoolHook(hook pool.PoolHook)
	RemovePoolHook(hook pool.PoolHook)
	SupportsPoolHooks() bool
}

type cscEvictOnRemoveHook struct {
	evictor Cache

	mu sync.Mutex

	initGen map[uint64]uint64
}

func (h *cscEvictOnRemoveHook) OnGet(_ context.Context, _ *pool.Conn, _ bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (h *cscEvictOnRemoveHook) OnPut(_ context.Context, _ *pool.Conn) (shouldPool, shouldRemove bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (h *cscEvictOnRemoveHook) OnRemove(_ context.Context, cn *pool.Conn, _ error) {
	_ = "STUB: not implemented"
	return
}

func (h *cscEvictOnRemoveHook) markRemoved(connID uint64) { _ = "STUB: not implemented"; return }

func (h *cscEvictOnRemoveHook) bumpInitGen(connID uint64) { _ = "STUB: not implemented"; return }

func (h *cscEvictOnRemoveHook) invalidateConnCoverage(connID uint64) {
	_ = "STUB: not implemented"
	return
}

func (h *cscEvictOnRemoveHook) initGenOf(connID uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (h *cscEvictOnRemoveHook) forgetConn(connID uint64) { _ = "STUB: not implemented"; return }

func (h *cscEvictOnRemoveHook) fulfillOwnedIfCovered(
	cacheKey string,
	token, ownerConnID, capturedGen uint64,
	value []byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *cscEvictOnRemoveHook) invalidateAllCoverage() { _ = "STUB: not implemented"; return }

func (c *baseClient) registerConnEvictHook(cache Cache, reg poolHookSupport) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) cscEvictOwnedEntries(connID uint64) { _ = "STUB: not implemented"; return }

func (c *baseClient) newStickyConnPool() *pool.StickyConnPool {
	_ = "STUB: not implemented"
	return nil
}

type cscFetchCapture struct {
	raw     []byte
	connID  uint64
	initGen uint64
}

func (c *baseClient) cscConnInitGen(connID uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (c *baseClient) cscForgetConn(connID uint64) { _ = "STUB: not implemented"; return }

var errClientTrackingWithCSC = errors.New(
	"redis: CLIENT TRACKING is not allowed when client-side caching is enabled")

var errSelectWithCSC = errors.New(
	"redis: SELECT is not allowed when client-side caching is enabled")

var errAuthWithCSC = errors.New(
	"redis: AUTH is not allowed when client-side caching is enabled")

var errHelloWithCSC = errors.New(
	"redis: HELLO with arguments is not allowed when client-side caching is enabled")

var errResetWithCSC = errors.New(
	"redis: RESET is not allowed when client-side caching is enabled")

var errSubscribeWithCSC = errors.New(
	"redis: SUBSCRIBE is not allowed on pooled connections when client-side caching is enabled")

func (c *baseClient) cscCommandError(cmd Cmder) error { _ = "STUB: not implemented"; return nil }

type cscDrainHandle struct {
	stop              chan struct{}
	done              chan struct{}
	stopOnce          sync.Once
	teardownOnce      sync.Once
	handlerCloseOnce  sync.Once
	closeOnce         sync.Once
	closeErr          error
	invalidateHandler *invalidateHandler
}

func (h *cscDrainHandle) signalStop() { _ = "STUB: not implemented"; return }

type cscHandlerClient struct {
	*baseClient
}

func (c cscHandlerClient) Close() error { _ = "STUB: not implemented"; return nil }

const cscMinDrainInterval = time.Millisecond

func (c *baseClient) cscDrainInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type idleConnDrainer interface {
	DrainIdleConns(ctx context.Context, st *pool.DrainState, fn func(cn *pool.Conn) error)
}

func (c *baseClient) startBackgroundDrainer() { _ = "STUB: not implemented"; return }

func (c *baseClient) disableCSCServing(ctx context.Context, reason string) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) stopBackgroundDrainer() { _ = "STUB: not implemented"; return }

func applyCachedReply(cmd Cmder, raw []byte) error { _ = "STUB: not implemented"; return nil }

func isCacheableReplyResult(err error) bool { _ = "STUB: not implemented"; return false }

const cscDrainSkipWindow = 5 * time.Millisecond

var cscDrainHardReadCap = 50 * time.Millisecond

const cscDrainProbeReadCap = 50 * time.Microsecond

const cscDrainCustomErrCap = 8

func (c *baseClient) processCached(ctx context.Context, cmd Cmder, state *processState) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) fulfillCached(key string, token uint64, fc *cscFetchCapture) bool {
	_ = "STUB: not implemented"
	return false
}
