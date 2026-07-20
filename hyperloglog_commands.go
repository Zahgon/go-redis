package redis

import "context"

type HyperLogLogCmdable interface {
	PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
	PFCount(ctx context.Context, keys ...string) *IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
}

func (c cmdable) PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PFCount(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}
