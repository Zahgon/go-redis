package redis

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9/auth"
	"github.com/redis/go-redis/v9/internal"
	"github.com/redis/go-redis/v9/internal/auth/streaming"
	"github.com/redis/go-redis/v9/internal/hscan"
	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/redis/go-redis/v9/push"
)

type Scanner = hscan.Scanner

const Nil = proto.Nil

const (
	NaN  = internal.NaN
	Inf  = internal.Inf
	NInf = internal.NInf
)

func SetLogger(logger internal.Logging) { _ = "STUB: not implemented"; return }

func SetLogLevel(logLevel internal.LogLevelT) { _ = "STUB: not implemented"; return }

type Hook interface {
	DialHook(next DialHook) DialHook
	ProcessHook(next ProcessHook) ProcessHook
	ProcessPipelineHook(next ProcessPipelineHook) ProcessPipelineHook
}

type (
	DialHook            func(ctx context.Context, network, addr string) (net.Conn, error)
	ProcessHook         func(ctx context.Context, cmd Cmder) error
	ProcessPipelineHook func(ctx context.Context, cmds []Cmder) error
)

type hooksMixin struct {
	hooksMu *sync.Mutex

	state *atomic.Pointer[hooksState]
}

type hooksState struct {
	slice   []Hook
	initial hooks
	current hooks
}

func (s *hooksState) rebuild() { _ = "STUB: not implemented"; return }

func (hs *hooksMixin) initHooks(hooks hooks) { _ = "STUB: not implemented"; return }

type hooks struct {
	dial       DialHook
	process    ProcessHook
	pipeline   ProcessPipelineHook
	txPipeline ProcessPipelineHook
}

func (h *hooks) setDefaults() { _ = "STUB: not implemented"; return }

func (hs *hooksMixin) AddHook(hook Hook) { _ = "STUB: not implemented"; return }

func (hs *hooksMixin) clone() hooksMixin { _ = "STUB: not implemented"; return *new(hooksMixin) }

func (hs *hooksMixin) withProcessHook(ctx context.Context, cmd Cmder, hook ProcessHook) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *hooksMixin) withProcessPipelineHook(
	ctx context.Context, cmds []Cmder, hook ProcessPipelineHook,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *hooksMixin) dialHook(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (hs *hooksMixin) hookCount() int { _ = "STUB: not implemented"; return 0 }

func (hs *hooksMixin) processHook(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *hooksMixin) processPipelineHook(ctx context.Context, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (hs *hooksMixin) processTxPipelineHook(ctx context.Context, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	onCloseHookIDSentinelFailover = "sentinel-failover"
)

type onCloseHooks struct {
	mu    sync.Mutex
	order []string
	hooks map[string]func() error
}

func (h *onCloseHooks) register(id string, fn func() error) { _ = "STUB: not implemented"; return }

//nolint:unused // kept for API symmetry with register; see comment above.
func (h *onCloseHooks) unregister(id string) { _ = "STUB: not implemented"; return }

func (h *onCloseHooks) run() error { _ = "STUB: not implemented"; return nil }

type pipelinePoolRef struct {
	pool *pool.ConnPool
	name string
}

type baseClient struct {
	apClosed *atomic.Bool

	opt        *Options
	optLock    sync.RWMutex
	connPool   pool.Pooler
	pubSubPool *pool.PubSubPool

	pipelinePool *pipelinePoolRef
	hooksMixin

	onClose *onCloseHooks

	pushProcessor push.NotificationProcessor

	maintNotificationsManager     *maintnotifications.Manager
	maintNotificationsManagerLock sync.RWMutex

	streamingCredentialsManager *streaming.Manager

	himport *himportRegistry

	csc Cache

	cscKeyPrefix string

	allowClientTracking bool

	cscOwnsCache bool

	cscDrainHandle *cscDrainHandle

	cscPoolHook pool.PoolHook

	cscActive *atomic.Bool
}

func (c *baseClient) clone() *baseClient { _ = "STUB: not implemented"; return nil }

func (c *baseClient) cloneOpt() *Options { _ = "STUB: not implemented"; return nil }

func (c *baseClient) withTimeout(timeout time.Duration) *baseClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) String() string { _ = "STUB: not implemented"; return "" }

func (c *baseClient) getConn(ctx context.Context) (*pool.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseClient) _getConn(ctx context.Context) (*pool.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseClient) initPooledConn(ctx context.Context, p pool.Pooler, cn *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) loadPipelinePool() *pipelinePoolRef { _ = "STUB: not implemented"; return nil }

func (c *baseClient) getPipelinePool() pool.Pooler {
	_ = "STUB: not implemented"
	return *new(pool.Pooler)
}

func (c *baseClient) isPipelinePoolConn(cn *pool.Conn) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *baseClient) poolForConn(cn *pool.Conn) pool.Pooler {
	_ = "STUB: not implemented"
	return *new(pool.Pooler)
}

func pipelinePoolOptions(opt *Options) *Options { _ = "STUB: not implemented"; return nil }

func (c *baseClient) buildPipelinePool(poolName string) (*pipelinePoolRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *baseClient) reAuthConnection() func(poolCn *pool.Conn, credentials auth.Credentials) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) onAuthenticationErr() func(poolCn *pool.Conn, err error) {
	_ = "STUB: not implemented"
	return nil
}

func (opt *Options) resolveCredentials(ctx context.Context) (username, password string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (c *baseClient) initConn(ctx context.Context, cn *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) releaseConn(ctx context.Context, cn *pool.Conn, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) releaseConnToPool(ctx context.Context, p pool.Pooler, cn *pool.Conn, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *baseClient) withConn(
	ctx context.Context, fn func(context.Context, *pool.Conn) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) withPipelineConn(
	ctx context.Context, fn func(context.Context, *pool.Conn) error,
) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (c *baseClient) cscTrackingRequested() bool { _ = "STUB: not implemented"; return false }

func (c *baseClient) autopipelineCSCActive() bool { _ = "STUB: not implemented"; return false }

func (c *baseClient) process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

type processState struct {
	attempts int
	lastConn *pool.Conn
}

func (c *baseClient) processCommand(ctx context.Context, cmd Cmder, state *processState) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) processWithRetry(
	ctx context.Context, cmd Cmder, capture *cscFetchCapture, state *processState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func classifyCommandError(err error) (errorType, statusCode string, isInternal bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (c *baseClient) _process(ctx context.Context, cmd Cmder, attempt int, capture *cscFetchCapture) (bool, *pool.Conn, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (c *baseClient) retryBackoff(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *baseClient) cmdTimeout(cmd Cmder) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *baseClient) context(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *baseClient) createInitConnFunc() func(context.Context, *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) enableMaintNotificationsUpgrades() error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) disableMaintNotificationsUpgrades() error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) Close() error { _ = "STUB: not implemented"; return nil }

func (c *baseClient) closeResources() error { _ = "STUB: not implemented"; return nil }

func (c *baseClient) getAddr() string { _ = "STUB: not implemented"; return "" }

func (c *baseClient) processPipeline(ctx context.Context, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) processTxPipeline(ctx context.Context, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

type pipelineProcessor func(context.Context, *pool.Conn, []Cmder) (bool, error)

func (c *baseClient) generalProcessPipeline(
	ctx context.Context, cmds []Cmder, p pipelineProcessor, operationName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) pipelineProcessCmds(
	ctx context.Context, cn *pool.Conn, cmds []Cmder,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *baseClient) pipelineReadCmds(ctx context.Context, cn *pool.Conn, rd *proto.Reader, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) txPipelineProcessCmds(
	ctx context.Context, cn *pool.Conn, cmds []Cmder,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *baseClient) txPipelineReadQueued(ctx context.Context, cn *pool.Conn, rd *proto.Reader, statusCmd *StatusCmd, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

type Client struct {
	*baseClient
	cmdable

	cscLifecycleOwner *Client

	autopipelinerMu     *sync.Mutex
	autopipeliner       *AutoPipeliner
	asyncAutopipeliner  *AutoPipeliner
	autopipelinerClosed bool
}

func NewClient(opt *Options) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) init() {

	c.autopipelinerMu = &sync.Mutex{}
	c.autopipeliner = nil
	c.asyncAutopipeliner = nil
	c.cmdable = c.Process
	c.initHooks(hooks{
		dial:       c.baseClient.dial,
		process:    c.baseClient.process,
		pipeline:   c.baseClient.processPipeline,
		txPipeline: c.baseClient.processTxPipeline,
	})
}

func (c *Client) WithTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) Conn() *Conn { _ = "STUB: not implemented"; return nil }

func (c *Client) Process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Options() *Options { _ = "STUB: not implemented"; return nil }

func (c *Client) NodeAddress() string { _ = "STUB: not implemented"; return "" }

func (c *Client) GetMaintNotificationsManager() *maintnotifications.Manager {
	_ = "STUB: not implemented"
	return nil
}

func initializePushProcessor(opt *Options) push.NotificationProcessor {
	_ = "STUB: not implemented"
	return *new(push.NotificationProcessor)
}

func (c *Client) RegisterPushNotificationHandler(pushNotificationName string, handler push.NotificationHandler, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) GetPushNotificationHandler(pushNotificationName string) push.NotificationHandler {
	_ = "STUB: not implemented"
	return *new(push.NotificationHandler)
}

type PoolStats pool.Stats

func (c *Client) PoolStats() *PoolStats { _ = "STUB: not implemented"; return nil }

func (c *Client) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Client) AutoPipeline() (*AutoPipeliner, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) AutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) AsyncAutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) AsyncAutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Client) pubSub() *PubSub { _ = "STUB: not implemented"; return nil }

func (c *Client) Subscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) PSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) SSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

type Conn struct {
	baseClient
	cmdable
	statefulCmdable
}

func newConn(opt *Options, connPool pool.Pooler, parentHooks *hooksMixin, himport *himportRegistry) *Conn {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) Process(ctx context.Context, cmd Cmder) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) RegisterPushNotificationHandler(pushNotificationName string, handler push.NotificationHandler, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Conn) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *baseClient) processPushNotifications(ctx context.Context, cn *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) peekAndProcessPushNotifications(ctx context.Context, cn *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

const cscFallbackProbeInterval = 100 * time.Millisecond

func (c *baseClient) drainPushNotifications(cn *pool.Conn) (processorSucceeded bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *baseClient) processPendingPushNotificationWithReader(ctx context.Context, cn *pool.Conn, rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *baseClient) pushNotificationHandlerContext(cn *pool.Conn) push.NotificationHandlerContext {
	_ = "STUB: not implemented"
	return *new(push.NotificationHandlerContext)
}
