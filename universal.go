package redis

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/redis/go-redis/v9/auth"
	"github.com/redis/go-redis/v9/maintnotifications"
	"github.com/redis/go-redis/v9/push"
)

type UniversalOptions struct {
	Addrs []string

	ClientName string

	DB int

	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *Conn) error

	Protocol int
	Username string
	Password string

	CredentialsProvider func() (username string, password string)

	CredentialsProviderContext func(ctx context.Context) (username string, password string, err error)

	StreamingCredentialsProvider auth.StreamingCredentialsProvider

	SentinelUsername string
	SentinelPassword string

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout time.Duration

	DialerRetries int

	DialerRetryTimeout time.Duration

	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool

	ReadBufferSize int

	WriteBufferSize int

	PipelineReadBufferSize  int
	PipelineWriteBufferSize int
	PipelinePoolSize        int

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

	MaxRedirects   int
	ReadOnly       bool
	RouteByLatency bool

	RouteByLatencyTolerance time.Duration
	RouteRandomly           bool

	MasterName string

	DisableIndentity bool

	DisableIdentity bool

	IdentitySuffix string

	FailingTimeoutSeconds int

	UnstableResp3 bool

	PushNotificationProcessor push.NotificationProcessor

	IsClusterMode bool

	AutoPipelineOptions *AutoPipelineOptions

	MaintNotificationsConfig *maintnotifications.Config

	ClientSideCacheConfig *ClientSideCacheConfig

	ClientSideCache Cache

	ClientSideCacheStrategy CSCStrategy
}

func (o *UniversalOptions) Cluster() *ClusterOptions { _ = "STUB: not implemented"; return nil }

func (o *UniversalOptions) Failover() *FailoverOptions { _ = "STUB: not implemented"; return nil }

func (o *UniversalOptions) Simple() *Options { _ = "STUB: not implemented"; return nil }

type UniversalClient interface {
	Cmdable
	AddHook(Hook)
	Watch(ctx context.Context, fn func(*Tx) error, keys ...string) error
	Do(ctx context.Context, args ...interface{}) *Cmd
	Process(ctx context.Context, cmd Cmder) error

	AutoPipeline() (*AutoPipeliner, error)
	AutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error)
	AsyncAutoPipeline() (*AutoPipeliner, error)
	AsyncAutoPipelineWithOptions(config *AutoPipelineOptions) (*AutoPipeliner, error)
	Subscribe(ctx context.Context, channels ...string) *PubSub
	PSubscribe(ctx context.Context, channels ...string) *PubSub
	SSubscribe(ctx context.Context, channels ...string) *PubSub
	Close() error
	PoolStats() *PoolStats
}

var (
	_ UniversalClient = (*Client)(nil)
	_ UniversalClient = (*ClusterClient)(nil)
	_ UniversalClient = (*Ring)(nil)

	_ UniversalClient = (*AutoPipeliner)(nil)
)

func NewUniversalClient(opts *UniversalOptions) UniversalClient {
	_ = "STUB: not implemented"
	return *new(UniversalClient)
}
