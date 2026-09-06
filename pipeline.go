package redis

import (
	"context"
)

type pipelineExecer func(context.Context, []Cmder) error

type Pipeliner interface {
	StatefulCmdable

	Len() int

	Do(ctx context.Context, args ...interface{}) *Cmd

	Process(ctx context.Context, cmd Cmder) error

	BatchProcess(ctx context.Context, cmd ...Cmder) error

	Discard()

	Exec(ctx context.Context) ([]Cmder, error)

	Cmds() []Cmder
}

var _ Pipeliner = (*Pipeline)(nil)

type Pipeline struct {
	cmdable
	statefulCmdable

	exec pipelineExecer
	cmds []Cmder
}

func (c *Pipeline) init() {
	c.cmdable = c.Process
	c.statefulCmdable = c.Process
}

func (c *Pipeline) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Pipeline) Do(ctx context.Context, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Process(ctx context.Context, cmd Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BatchProcess(ctx context.Context, cmd ...Cmder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Discard() { _ = "STUB: not implemented"; return }

func (c *Pipeline) Exec(ctx context.Context) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Pipeline) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Pipeline) Cmds() []Cmder { _ = "STUB: not implemented"; return nil }
