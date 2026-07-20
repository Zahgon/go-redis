package redis

import (
	"context"
)

type PubSubCmdable interface {
	Publish(ctx context.Context, channel string, message interface{}) *IntCmd
	SPublish(ctx context.Context, channel string, message interface{}) *IntCmd
	PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *MapStringIntCmd
	PubSubNumPat(ctx context.Context) *IntCmd
	PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd
	PubSubShardNumSub(ctx context.Context, channels ...string) *MapStringIntCmd
}

func (c cmdable) Publish(ctx context.Context, channel string, message interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) SPublish(ctx context.Context, channel string, message interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PubSubNumSub(ctx context.Context, channels ...string) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PubSubShardNumSub(ctx context.Context, channels ...string) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) PubSubNumPat(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }
