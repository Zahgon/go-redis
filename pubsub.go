package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9/internal/pool"
	"github.com/redis/go-redis/v9/internal/proto"
	"github.com/redis/go-redis/v9/push"
)

type PubSub struct {
	opt *Options

	newConn   func(ctx context.Context, addr string, channels []string) (*pool.Conn, error)
	closeConn func(*pool.Conn) error

	mu        sync.Mutex
	cn        *pool.Conn
	channels  map[string]struct{}
	patterns  map[string]struct{}
	schannels map[string]struct{}

	closed bool
	exit   chan struct{}

	cmd *Cmd

	chOnce sync.Once
	msgCh  *channel
	allCh  *channel

	pushProcessor push.NotificationProcessor

	onClose func()
}

func (c *PubSub) init() {
	c.exit = make(chan struct{})
}

func (c *PubSub) String() string { _ = "STUB: not implemented"; return "" }

func (c *PubSub) connWithLock(ctx context.Context) (*pool.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) conn(ctx context.Context, newChannels []string) (*pool.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) writeCmd(ctx context.Context, cn *pool.Conn, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) resubscribe(ctx context.Context, cn *pool.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) _subscribe(
	ctx context.Context, cn *pool.Conn, redisCmd string, channels []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) releaseConnWithLock(
	ctx context.Context,
	cn *pool.Conn,
	err error,
	allowTimeout bool,
) {
	_ = "STUB: not implemented"
	return
}

func (c *PubSub) releaseConn(ctx context.Context, cn *pool.Conn, err error, allowTimeout bool) {
	_ = "STUB: not implemented"
	return
}

func (c *PubSub) reconnect(ctx context.Context, reason error) { _ = "STUB: not implemented"; return }

func (c *PubSub) closeTheCn(reason error) error { _ = "STUB: not implemented"; return nil }

func (c *PubSub) Close() error { _ = "STUB: not implemented"; return nil }

func (c *PubSub) Subscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) PSubscribe(ctx context.Context, patterns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) SSubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) Unsubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) PUnsubscribe(ctx context.Context, patterns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) SUnsubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) subscribe(ctx context.Context, redisCmd string, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) Ping(ctx context.Context, payload ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) ClientSetName(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

type Subscription struct {
	Kind string

	Channel string

	Count int
}

func (m *Subscription) String() string { _ = "STUB: not implemented"; return "" }

type Message struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }

type Pong struct {
	Payload string
}

func (p *Pong) String() string { _ = "STUB: not implemented"; return "" }

func (c *PubSub) newMessage(ctx context.Context, cn *pool.Conn, reply interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) ReceiveTimeout(ctx context.Context, timeout time.Duration) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) Receive(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) ReceiveMessage(ctx context.Context) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PubSub) getContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *PubSub) Channel(opts ...ChannelOption) <-chan *Message {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) ChannelSize(size int) <-chan *Message { _ = "STUB: not implemented"; return nil }

func (c *PubSub) ChannelWithSubscriptions(opts ...ChannelOption) <-chan interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) processPendingPushNotificationWithReader(ctx context.Context, cn *pool.Conn, rd *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PubSub) pushNotificationHandlerContext(cn *pool.Conn) push.NotificationHandlerContext {
	_ = "STUB: not implemented"
	return *new(push.NotificationHandlerContext)
}

type ChannelOption func(c *channel)

func WithChannelSize(size int) ChannelOption { _ = "STUB: not implemented"; return *new(ChannelOption) }

func WithChannelHealthCheckInterval(d time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *new(ChannelOption)
}

func WithChannelSendTimeout(d time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *new(ChannelOption)
}

func WithChannelPingTimeout(d time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *new(ChannelOption)
}

func WithChannelReconnectTimeout(d time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *new(ChannelOption)
}

type channel struct {
	pubSub *PubSub

	msgCh chan *Message
	allCh chan interface{}
	ping  chan struct{}

	chanSize         int
	chanSendTimeout  time.Duration
	checkInterval    time.Duration
	pingTimeout      time.Duration
	reconnectTimeout time.Duration
}

func newChannel(pubSub *PubSub, opts ...ChannelOption) *channel {
	_ = "STUB: not implemented"
	return nil
}

func (c *channel) initHealthCheck() { _ = "STUB: not implemented"; return }

func (c *channel) initMsgChan() { _ = "STUB: not implemented"; return }

func (c *channel) initAllChan() { _ = "STUB: not implemented"; return }
