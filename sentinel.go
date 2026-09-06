package redis

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/redis/go-redis/v9/auth"
	"github.com/redis/go-redis/v9/push"
)

type FailoverOptions struct {
	MasterName string

	SentinelAddrs []string

	ClientName string

	SentinelUsername string

	SentinelPassword string

	RouteByLatency bool

	RouteByLatencyTolerance time.Duration

	RouteRandomly bool

	ReplicaOnly bool

	UseDisconnectedReplicas bool

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

	ReadBufferSize int

	WriteBufferSize int

	PipelineReadBufferSize  int
	PipelineWriteBufferSize int
	PipelinePoolSize        int

	AutoPipelineOptions *AutoPipelineOptions

	PoolFIFO bool

	PoolSize int

	MaxConcurrentDials int

	PoolTimeout           time.Duration
	MinIdleConns          int
	MaxIdleConns          int
	MaxActiveConns        int
	ConnMaxIdleTime       time.Duration
	ConnMaxLifetime       time.Duration
	ConnMaxLifetimeJitter time.Duration

	TLSConfig *tls.Config

	DisableIndentity bool

	DisableIdentity bool

	IdentitySuffix string

	FailingTimeoutSeconds int

	UnstableResp3 bool

	PushNotificationProcessor push.NotificationProcessor
}

func (opt *FailoverOptions) clientOptions() *Options { _ = "STUB: not implemented"; return nil }

func (opt *FailoverOptions) sentinelOptions(addr string) *Options {
	_ = "STUB: not implemented"
	return nil
}

func (opt *FailoverOptions) clusterOptions() *ClusterOptions { _ = "STUB: not implemented"; return nil }

func ParseFailoverURL(redisURL string) (*FailoverOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupFailoverConn(u *url.URL) (*FailoverOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupFailoverConnParams(u *url.URL, o *FailoverOptions) (*FailoverOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFailoverClient(failoverOpt *FailoverOptions) *Client { _ = "STUB: not implemented"; return nil }

func masterReplicaDialer(
	failover *sentinelFailover,
) func(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

type SentinelClient struct {
	*baseClient
}

func NewSentinelClient(opt *Options) *SentinelClient { _ = "STUB: not implemented"; return nil }

func (c *SentinelClient) GetPushNotificationHandler(pushNotificationName string) push.NotificationHandler {
	_ = "STUB: not implemented"
	return *new(push.NotificationHandler)
}

func (c *SentinelClient) RegisterPushNotificationHandler(pushNotificationName string, handler push.NotificationHandler, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) pubSub() *PubSub { _ = "STUB: not implemented"; return nil }

func (c *SentinelClient) Ping(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Subscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) PSubscribe(ctx context.Context, channels ...string) *PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) GetMasterAddrByName(ctx context.Context, name string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Sentinels(ctx context.Context, name string) *MapStringStringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Failover(ctx context.Context, name string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Reset(ctx context.Context, pattern string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) FlushConfig(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Master(ctx context.Context, name string) *MapStringStringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Masters(ctx context.Context) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Replicas(ctx context.Context, name string) *MapStringStringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) CkQuorum(ctx context.Context, name string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Monitor(ctx context.Context, name, ip, port, quorum string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Set(ctx context.Context, name, option, value string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *SentinelClient) Remove(ctx context.Context, name string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

type sentinelFailover struct {
	opt *FailoverOptions

	sentinelAddrs []string

	onFailover func(ctx context.Context, addr string)
	onUpdate   func(ctx context.Context)

	mu         sync.RWMutex
	masterAddr string
	sentinel   *SentinelClient
	pubsub     *PubSub
}

func (c *sentinelFailover) Close() error { _ = "STUB: not implemented"; return nil }

func (c *sentinelFailover) closeSentinel() error { _ = "STUB: not implemented"; return nil }

func (c *sentinelFailover) RandomReplicaAddr(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *sentinelFailover) MasterAddr(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *sentinelFailover) replicaAddrs(ctx context.Context, useDisconnected bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sentinelFailover) getMasterAddr(ctx context.Context, sentinel *SentinelClient) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *sentinelFailover) getReplicaAddrs(ctx context.Context, sentinel *SentinelClient) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseReplicaAddrs(addrs []map[string]string, keepDisconnected bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *sentinelFailover) trySwitchMaster(ctx context.Context, addr string) {
	_ = "STUB: not implemented"
	return
}

//nolint:ifshort

func (c *sentinelFailover) setSentinel(ctx context.Context, sentinel *SentinelClient) {
	_ = "STUB: not implemented"
	return
}

func (c *sentinelFailover) discoverSentinels(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (c *sentinelFailover) listen(pubsub *PubSub) { _ = "STUB: not implemented"; return }

func NewFailoverClusterClient(failoverOpt *FailoverOptions) *ClusterClient {
	_ = "STUB: not implemented"
	return nil
}
