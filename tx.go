package redis

import (
	"context"

	"github.com/redis/go-redis/v9/internal/proto"
)

const TxFailedErr = proto.RedisError("redis: transaction failed")

type Tx struct {
	baseClient
	cmdable
	statefulCmdable

	watchArmed bool
}

func (c *Client) newTx() *Tx { _ = "STUB: not implemented"; return nil }

func (c *Tx) init() {
	c.cmdable = c.Process
	c.statefulCmdable = c.Process

	c.initHooks(hooks{
		dial:       c.baseClient.dial,
		process:    c.baseClient.process,
		pipeline:   c.baseClient.processPipeline,
		txPipeline: c.baseClient.processTxPipeline,
	})
}

func (c *Tx) Process(ctx context.Context, cmd Cmder) error { _ = "STUB: not implemented"; return nil }

func (c *Client) Watch(ctx context.Context, fn func(*Tx) error, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Tx) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Tx) Watch(ctx context.Context, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Tx) Unwatch(ctx context.Context, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Tx) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Tx) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Tx) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Tx) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func wrapMultiExec(ctx context.Context, cmds []Cmder) []Cmder {
	_ = "STUB: not implemented"
	return nil
}
