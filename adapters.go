package redis

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/redis/go-redis/v9/internal/interfaces"
	"github.com/redis/go-redis/v9/push"
)

var ErrInvalidCommand = errors.New("invalid command type")

var ErrInvalidPool = errors.New("invalid pool type")

func newClientAdapter(client *baseClient) interfaces.ClientInterface {
	_ = "STUB: not implemented"
	return *new(interfaces.ClientInterface)
}

type clientAdapter struct {
	client *baseClient
}

func (ca *clientAdapter) GetOptions() interfaces.OptionsInterface {
	_ = "STUB: not implemented"
	return *new(interfaces.OptionsInterface)
}

func (ca *clientAdapter) GetPushProcessor() interfaces.NotificationProcessor {
	_ = "STUB: not implemented"
	return *new(interfaces.NotificationProcessor)
}

type optionsAdapter struct {
	options *Options
}

func (oa *optionsAdapter) GetReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (oa *optionsAdapter) GetWriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (oa *optionsAdapter) GetNetwork() string { _ = "STUB: not implemented"; return "" }

func (oa *optionsAdapter) GetAddr() string { _ = "STUB: not implemented"; return "" }

func (oa *optionsAdapter) GetNodeAddress() string { _ = "STUB: not implemented"; return "" }

func (oa *optionsAdapter) IsTLSEnabled() bool { _ = "STUB: not implemented"; return false }

func (oa *optionsAdapter) GetProtocol() int { _ = "STUB: not implemented"; return 0 }

func (oa *optionsAdapter) GetPoolSize() int { _ = "STUB: not implemented"; return 0 }

func (oa *optionsAdapter) NewDialer() func(context.Context) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

type pushProcessorAdapter struct {
	processor push.NotificationProcessor
}

func (ppa *pushProcessorAdapter) RegisterHandler(pushNotificationName string, handler interface{}, protected bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ppa *pushProcessorAdapter) UnregisterHandler(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ppa *pushProcessorAdapter) GetHandler(pushNotificationName string) interface{} {
	_ = "STUB: not implemented"
	return nil
}
