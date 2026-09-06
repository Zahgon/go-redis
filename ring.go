package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9/auth"
	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
)

var errRingShardsDown = errors.New("redis: all ring shards are down")

var defaultHeartbeatFn = func(ctx context.Context, client *Client) bool {
	err := client.Ping(ctx).Err()
	return err == nil || err == pool.ErrPoolTimeout
}

type ConsistentHash interface {
	Get(string) string
}

func newRendezvous(shards []string) ConsistentHash {
	_ = "STUB: not implemented"
	return *new(ConsistentHash)
}

type RingOptions struct {
	Addrs map[string]string

	NewClient func(opt *Options) *Client

	himport *himportRegistry

	ClientName string

	HeartbeatFrequency time.Duration

	HeartbeatFn func(ctx context.Context, client *Client) bool

	NewConsistentHash func(shards []string) ConsistentHash

	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *Conn) error

	Protocol int
	Username string
	Password string

	CredentialsProvider func() (username string, password string)

	CredentialsProviderContext func(ctx context.Context) (username string, password string, err error)

	StreamingCredentialsProvider auth.StreamingCredentialsProvider
	DB                           int

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

	PoolFIFO bool

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

	TLSConfig *tls.Config
	Limiter   Limiter

	DisableIndentity bool

	DisableIdentity bool
	IdentitySuffix  string

	UnstableResp3 bool
}

func (opt *RingOptions) init() {
	if opt.NewClient == nil {
		opt.NewClient = func(opt *Options) *Client {
			return NewClient(opt)
		}
	}

	if opt.HeartbeatFrequency == 0 {
		opt.HeartbeatFrequency = 500 * time.Millisecond
	}

	if opt.HeartbeatFn == nil {
		opt.HeartbeatFn = defaultHeartbeatFn
	}

	if opt.NewConsistentHash == nil {
		opt.NewConsistentHash = newRendezvous
	}

	switch opt.MaxRetries {
	case -1:
		opt.MaxRetries = 0
	case 0:
		opt.MaxRetries = 3
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

	if opt.ReadBufferSize == 0 {
		opt.ReadBufferSize = proto.DefaultBufferSize
	}
	if opt.WriteBufferSize == 0 {
		opt.WriteBufferSize = proto.DefaultBufferSize
	}
}

func (opt *RingOptions) clientOptions() *Options { _ = "STUB: not implemented"; return nil }

type ringShard struct {
	Client *Client
	down   atomic.Int32
	addr   string
}

func newRingShard(opt *RingOptions, addr string) *ringShard { _ = "STUB: not implemented"; return nil }

func (shard *ringShard) String() string { _ = "STUB: not implemented"; return "" }

func (shard *ringShard) IsDown() bool { _ = "STUB: not implemented"; return false }

func (shard *ringShard) IsUp() bool { _ = "STUB: not implemented"; return false }

func (shard *ringShard) Vote(up bool) bool { _ = "STUB: not implemented"; return false }

type ringSharding struct {
	opt *RingOptions

	mu        sync.RWMutex
	shards    *ringShards
	closed    bool
	hash      ConsistentHash
	numShard  int
	onNewNode []func(rdb *Client)

	setAddrsMu sync.Mutex
}

type ringShards struct {
	m    map[string]*ringShard
	list []*ringShard
}

func newRingSharding(opt *RingOptions) *ringSharding { _ = "STUB: not implemented"; return nil }

func (c *ringSharding) OnNewNode(fn func(rdb *Client)) { _ = "STUB: not implemented"; return }

func (c *ringSharding) SetAddrs(addrs map[string]string) { _ = "STUB: not implemented"; return }

func (c *ringSharding) newRingShards(
	addrs map[string]string, existing *ringShards, onNewNode []func(rdb *Client),
) (shards *ringShards, created, unused map[string]*ringShard) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *ringSharding) List() []*ringShard { _ = "STUB: not implemented"; return nil }

func (c *ringSharding) Hash(key string) string { _ = "STUB: not implemented"; return "" }

func (c *ringSharding) GetByKey(key string) (*ringShard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ringSharding) GetByName(shardName string) (*ringShard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ringSharding) Random() (*ringShard, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *ringSharding) Heartbeat(ctx context.Context, frequency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *ringSharding) rebalanceLocked() { _ = "STUB: not implemented"; return }

func (c *ringSharding) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *ringSharding) Close() error { _ = "STUB: not implemented"; return nil }

type Ring struct {
	cmdable
	hooksMixin

	opt               *RingOptions
	sharding          *ringSharding
	cmdsInfoCache     *cmdsInfoCache
	heartbeatCancelFn context.CancelFunc
}

func NewRing(opt *RingOptions) *Ring { _ = "STUB: not implemented"; return nil }

func (c *Ring) SetAddrs(addrs map[string]string) { _ = "STUB: not implemented"; return }

func (c *Ring) Process(ctx context.Context, cmd Cmder) error { _ = "STUB: not implemented"; return nil }

func (c *Ring) Options() *RingOptions { _ = "STUB: not implemented"; return nil }

func (c *Ring) retryBackoff(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Ring) PoolStats() *PoolStats { _ = "STUB: not implemented"; return nil }

func (c *Ring) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Ring) Subscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) PSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) SSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) Publish(ctx context.Context, channel string, message interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) OnNewNode(fn func(rdb *Client)) { _ = "STUB: not implemented"; return }

func (c *Ring) ForEachShard(
	ctx context.Context,
	fn func(ctx context.Context, client *Client) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) cmdsInfo(ctx context.Context) (map[string]*CommandInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) cmdShard(cmd Cmder) (*ringShard, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Ring) process(ctx context.Context, cmd Cmder) error { _ = "STUB: not implemented"; return nil }

func (c *Ring) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

var ErrRingAutoPipelineUnsupported = errors.New("redis: AutoPipeline is not supported by Ring")

func (c *Ring) AutoPipeline() (*AutoPipeliner, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Ring) AutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) AsyncAutoPipeline() (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) AsyncAutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Ring) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Ring) generalProcessPipeline(
	ctx context.Context, cmds []Cmder, tx bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) Watch(ctx context.Context, fn func(*Tx) error, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Ring) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Ring) GetShardClients() []*Client { _ = "STUB: not implemented"; return nil }

func (c *Ring) GetShardClientForKey(key string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
