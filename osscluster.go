package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"math"
	"net"
	"net/url"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9/auth"
	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
	"github.com/redis/go-redis/v9/internal/routing"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/redis/go-redis/v9/push"
)

const (
	minLatencyMeasurementInterval = 10 * time.Second
)

var (
	errClusterNoNodes = errors.New("redis: cluster has no nodes")
	errNoWatchKeys    = errors.New("redis: Watch requires at least one key")
	errWatchCrosslot  = errors.New("redis: Watch requires all keys to be in the same slot")
)

type ClusterOptions struct {
	Addrs []string

	ClientName string

	NewClient func(opt *Options) *Client

	MaxRedirects int

	ReadOnly bool

	RouteByLatency bool

	RouteByLatencyTolerance time.Duration

	RouteRandomly bool

	ClusterSlots func(context.Context) ([]ClusterSlot, error)

	Dialer func(ctx context.Context, network, addr string) (net.Conn, error)

	OnConnect func(ctx context.Context, cn *Conn) error

	Protocol                     int
	Username                     string
	Password                     string
	CredentialsProvider          func() (username string, password string)
	CredentialsProviderContext   func(ctx context.Context) (username string, password string, err error)
	StreamingCredentialsProvider auth.StreamingCredentialsProvider

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout time.Duration

	DialerRetries int

	DialerRetryTimeout time.Duration

	DialerRetryBackoff func(attempt int) time.Duration

	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool

	MaxConcurrentDials int

	PoolFIFO              bool
	PoolSize              int
	PoolTimeout           time.Duration
	MinIdleConns          int
	MaxIdleConns          int
	MaxActiveConns        int
	ConnMaxIdleTime       time.Duration
	ConnMaxLifetime       time.Duration
	ConnMaxLifetimeJitter time.Duration

	ReadBufferSize int

	WriteBufferSize int

	PipelineReadBufferSize  int
	PipelineWriteBufferSize int
	PipelinePoolSize        int

	AutoPipelineOptions *AutoPipelineOptions

	TLSConfig *tls.Config

	DisableRoutingPolicies bool

	DisableIndentity bool

	DisableIdentity bool

	IdentitySuffix string

	UnstableResp3 bool

	PushNotificationProcessor push.NotificationProcessor

	FailingTimeoutSeconds int

	MaintNotificationsConfig *maintnotifications.Config

	ShardPicker routing.ShardPicker

	ClusterStateReloadInterval time.Duration
}

func (opt *ClusterOptions) init() {
	switch opt.MaxRedirects {
	case -1:
		opt.MaxRedirects = 0
	case 0:
		opt.MaxRedirects = 3
	}

	if opt.RouteByLatency || opt.RouteRandomly {
		opt.ReadOnly = true
	}

	if opt.DialTimeout == 0 {
		opt.DialTimeout = 5 * time.Second
	}
	if opt.DialerRetries == 0 {
		opt.DialerRetries = 5
	}
	if opt.DialerRetryTimeout == 0 {
		opt.DialerRetryTimeout = 100 * time.Millisecond
	}

	if opt.PoolSize == 0 {
		opt.PoolSize = 5 * runtime.GOMAXPROCS(0)
	}
	if opt.MaxConcurrentDials <= 0 {
		opt.MaxConcurrentDials = opt.PoolSize
	} else if opt.MaxConcurrentDials > opt.PoolSize {
		opt.MaxConcurrentDials = opt.PoolSize
	}
	if opt.ReadBufferSize == 0 {
		opt.ReadBufferSize = proto.DefaultBufferSize
	}
	if opt.WriteBufferSize == 0 {
		opt.WriteBufferSize = proto.DefaultBufferSize
	}

	switch opt.ReadTimeout {
	case -1:
		opt.ReadTimeout = 0
	case 0:
		opt.ReadTimeout = 5 * time.Second
	}
	switch opt.WriteTimeout {
	case -1:
		opt.WriteTimeout = 0
	case 0:
		opt.WriteTimeout = opt.ReadTimeout
	}

	if opt.MaxRetries == 0 {
		opt.MaxRetries = -1
	}
	switch opt.MinRetryBackoff {
	case -1:
		opt.MinRetryBackoff = 0
	case 0:
		opt.MinRetryBackoff = 10 * time.Millisecond
	}
	switch opt.MaxRetryBackoff {
	case -1:
		opt.MaxRetryBackoff = 0
	case 0:
		opt.MaxRetryBackoff = time.Second
	}

	if opt.NewClient == nil {
		opt.NewClient = NewClient
	}

	if opt.FailingTimeoutSeconds == 0 {
		opt.FailingTimeoutSeconds = 15
	}

	if opt.ShardPicker == nil {
		opt.ShardPicker = &routing.RoundRobinPicker{}
	}

	if opt.ClusterStateReloadInterval == 0 {
		opt.ClusterStateReloadInterval = 60 * time.Second
	}
}

func ParseClusterURL(redisURL string) (*ClusterOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupClusterConn(u *url.URL, host string, o *ClusterOptions) (*ClusterOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupClusterQueryParams(u *url.URL, o *ClusterOptions) (*ClusterOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (opt *ClusterOptions) clientOptions() *Options { _ = "STUB: not implemented"; return nil }

type clusterNode struct {
	Client *Client

	latency    atomic.Uint32
	generation atomic.Uint32
	failing    atomic.Uint32
	loaded     atomic.Uint32

	lastLatencyMeasurement atomic.Int64
}

func newClusterNodeWithNodeAddress(clOpt *ClusterOptions, addr, nodeAddress string) *clusterNode {
	_ = "STUB: not implemented"
	return nil
}

func (n *clusterNode) String() string { _ = "STUB: not implemented"; return "" }

func (n *clusterNode) Close() error { _ = "STUB: not implemented"; return nil }

const maximumNodeLatency = 1 * time.Minute

const (
	unmeasuredNodeLatencyMicros uint32 = math.MaxUint32
	unmeasuredNodeLatency              = time.Duration(unmeasuredNodeLatencyMicros) * time.Microsecond
)

func (n *clusterNode) updateLatency() { _ = "STUB: not implemented"; return }

func (n *clusterNode) Latency() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (n *clusterNode) MarkAsFailing() { _ = "STUB: not implemented"; return }

func (n *clusterNode) Failing() bool { _ = "STUB: not implemented"; return false }

func (n *clusterNode) Generation() uint32 { _ = "STUB: not implemented"; return 0 }

func (n *clusterNode) LastLatencyMeasurement() int64 { _ = "STUB: not implemented"; return 0 }

func (n *clusterNode) SetGeneration(gen uint32) { _ = "STUB: not implemented"; return }

func (n *clusterNode) SetLastLatencyMeasurement(t time.Time) { _ = "STUB: not implemented"; return }

func (n *clusterNode) Loading() bool { _ = "STUB: not implemented"; return false }

type clusterNodes struct {
	opt *ClusterOptions

	mu          sync.RWMutex
	addrs       []string
	nodes       map[string]*clusterNode
	activeAddrs []string
	closed      bool
	onNewNode   []func(rdb *Client)

	generation atomic.Uint32
}

func newClusterNodes(opt *ClusterOptions) *clusterNodes { _ = "STUB: not implemented"; return nil }

func (c *clusterNodes) Close() error { _ = "STUB: not implemented"; return nil }

func (c *clusterNodes) OnNewNode(fn func(rdb *Client)) { _ = "STUB: not implemented"; return }

func (c *clusterNodes) Addrs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:ifshort

func (c *clusterNodes) NextGeneration() uint32 { _ = "STUB: not implemented"; return 0 }

func (c *clusterNodes) GC(generation uint32) { _ = "STUB: not implemented"; return }

func (c *clusterNodes) GetOrCreate(addr string) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterNodes) GetOrCreateWithNodeAddress(addr, nodeAddress string) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterNodes) get(addr string) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterNodes) All() ([]*clusterNode, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *clusterNodes) Random() (*clusterNode, error) { _ = "STUB: not implemented"; return nil, nil }

type clusterSlot struct {
	start int
	end   int
	nodes []*clusterNode

	latencyBandNodeCursor atomic.Uint32
}

type clusterState struct {
	nodes   *clusterNodes
	Masters []*clusterNode
	Slaves  []*clusterNode

	slots []*clusterSlot

	generation uint32
	createdAt  time.Time
}

func newClusterState(
	nodes *clusterNodes, slots []ClusterSlot, origin string,
) (*clusterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func replaceLoopbackHost(nodeAddr, originHost string) string { _ = "STUB: not implemented"; return "" }

func replaceZeroPort(nodeAddr, originPort string) string { _ = "STUB: not implemented"; return "" }

func isLoopback(host string) bool { _ = "STUB: not implemented"; return false }

func (c *clusterState) slotMasterNode(slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotSlaveNode(slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotClosestNode(slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotNodeWithinLatency(slot int, tolerance time.Duration) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotRandomNode(slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotShardPickerSlaveNode(slot int, shardPicker routing.ShardPicker) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterState) slotEntry(slot int) *clusterSlot { _ = "STUB: not implemented"; return nil }

func (c *clusterState) slotNodes(slot int) []*clusterNode { _ = "STUB: not implemented"; return nil }

type clusterStateHolder struct {
	load func(ctx context.Context) (*clusterState, error)

	reloadInterval time.Duration
	state          atomic.Value
	reloading      atomic.Uint32
	reloadPending  atomic.Uint32
}

func newClusterStateHolder(load func(ctx context.Context) (*clusterState, error), reloadInterval time.Duration) *clusterStateHolder {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterStateHolder) Reload(ctx context.Context) (*clusterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterStateHolder) LazyReload() { _ = "STUB: not implemented"; return }

func (c *clusterStateHolder) Get(ctx context.Context) (*clusterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterStateHolder) ReloadOrGet(ctx context.Context) (*clusterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClusterClient struct {
	opt             *ClusterOptions
	nodes           *clusterNodes
	state           *clusterStateHolder
	cmdsInfoCache   *cmdsInfoCache
	cmdInfoResolver *commandInfoResolver
	cmdable
	hooksMixin

	himport *himportRegistry

	autopipelinerMu     *sync.Mutex
	autopipeliner       *AutoPipeliner
	asyncAutopipeliner  *AutoPipeliner
	autopipelinerClosed bool
}

func NewClusterClient(opt *ClusterOptions) *ClusterClient { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) Options() *ClusterOptions { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) ReloadState(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *ClusterClient) Close() error { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) Process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) OnNewNode(fn func(rdb *Client)) { _ = "STUB: not implemented"; return }

func (c *ClusterClient) ForEachMaster(
	ctx context.Context,
	fn func(ctx context.Context, client *Client) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) ForEachSlave(
	ctx context.Context,
	fn func(ctx context.Context, client *Client) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) ForEachShard(
	ctx context.Context,
	fn func(ctx context.Context, client *Client) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) PoolStats() *PoolStats { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) loadState(ctx context.Context) (*clusterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func clusterAutoPipelineOptions(cfg *AutoPipelineOptions) *AutoPipelineOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) AutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) AutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) installAutoPipelineSharding(ap *AutoPipeliner) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterClient) AsyncAutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) AsyncAutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) processPipeline(ctx context.Context, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) mapCmdsByNode(ctx context.Context, cmdsMap *cmdsMap, cmds []Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) cmdsAreReadOnly(ctx context.Context, cmds []Cmder) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ClusterClient) processPipelineNode(
	ctx context.Context, node *clusterNode, cmds []Cmder, failedCmds *cmdsMap,
) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterClient) processPipelineNodeConn(
	ctx context.Context, node *clusterNode, cn *pool.Conn, cmds []Cmder, failedCmds *cmdsMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) pipelineReadCmds(
	ctx context.Context,
	node *clusterNode,
	cn *pool.Conn,
	rd *proto.Reader,
	cmds []Cmder,
	failedCmds *cmdsMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) checkMovedErr(
	ctx context.Context, cmd Cmder, err error, failedCmds *cmdsMap,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ClusterClient) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *ClusterClient) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type txOutcomeKind int

const (
	txSuccess txOutcomeKind = iota
	txRetryMoved
	txRetryAsk
	txRetryTryAgain
	txRetryConn
	txFatal
)

type txOutcome struct {
	kind          txOutcomeKind
	err           error
	addr          string
	execErr       error
	unreadReplies bool
}

type txRedirect struct {
	moved    bool
	ask      bool
	tryAgain bool
	addr     string
	err      error
}

var errTxDirtyConn = errors.New("redis: connection has unread transaction replies")

func (c *ClusterClient) processTxPipeline(ctx context.Context, cmds []Cmder) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) slottedKeyedCommands(_ context.Context, cmds []Cmder) map[int][]Cmder {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) processTxPipelineNode(
	ctx context.Context, node *clusterNode, cmds []Cmder, asking bool,
) *txOutcome {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) processTxPipelineNodeConn(
	ctx context.Context, node *clusterNode, cn *pool.Conn, wire []Cmder, cmds []Cmder, asking bool,
) *txOutcome {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) readTxPipelineReplies(
	ctx context.Context, node *clusterNode, cn *pool.Conn, rd *proto.Reader, cmds []Cmder, asking bool,
) *txOutcome {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) txProcessPush(ctx context.Context, node *clusterNode, cn *pool.Conn, rd *proto.Reader) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterClient) txReadFatal(err error) *txOutcome { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) txPreQueueErrorOutcome(err error, cmds []Cmder) *txOutcome {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) classifyExecError(execErr error, firstRedirect *txRedirect, firstFatal error) *txOutcome {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) Watch(ctx context.Context, fn func(*Tx) error, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) pubSub() *PubSub { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) Subscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) PSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) SSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) retryBackoff(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *ClusterClient) cmdsInfo(ctx context.Context) (map[string]*CommandInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) cmdInfo(ctx context.Context, name string) *CommandInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) cmdInfoPeek(name string) *CommandInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) cmdSlot(cmd Cmder, prefferedSlot int) int {
	_ = "STUB: not implemented"
	return 0
}

func (c *ClusterClient) cmdSlotWithPos(cmd Cmder, pos int, prefferedSlot int) int {
	_ = "STUB: not implemented"
	return 0
}

func cmdSlot(cmd Cmder, pos int, prefferedRandomSlot int) int { _ = "STUB: not implemented"; return 0 }

func (c *ClusterClient) cmdNode(
	ctx context.Context,
	cmdName string,
	slot int,
) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) cmdNodeWithShardPicker(
	ctx context.Context,
	cmdName string,
	slot int,
	shardPicker routing.ShardPicker,
) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) slotReadOnlyNode(state *clusterState, slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) slotMasterNode(ctx context.Context, slot int) (*clusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) SlaveForKey(ctx context.Context, key string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) MasterForKey(ctx context.Context, key string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterClient) context(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *ClusterClient) GetResolver() *commandInfoResolver { _ = "STUB: not implemented"; return nil }

func (c *ClusterClient) SetCommandInfoResolver(cmdInfoResolver *commandInfoResolver) {
	_ = "STUB: not implemented"
	return
}

func (c *ClusterClient) extractCommandInfo(ctx context.Context, cmd Cmder) *routing.CommandPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterClient) NewDynamicResolver() *commandInfoResolver {
	_ = "STUB: not implemented"
	return nil
}

func appendIfNotExist[T comparable](vals []T, newVal T) []T { _ = "STUB: not implemented"; return nil }

type cmdsMap struct {
	mu sync.Mutex
	m  map[*clusterNode][]Cmder
}

func newCmdsMap() *cmdsMap { _ = "STUB: not implemented"; return nil }

func (m *cmdsMap) Add(node *clusterNode, cmds ...Cmder) { _ = "STUB: not implemented"; return }
